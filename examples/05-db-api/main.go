package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
	"unicode/utf8"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Student struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type ClassSession struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Capacity int    `json:"capacity"`
}

func reply(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(value); err != nil {
		log.Print("response write failed")
	}
}

func dbFailure(w http.ResponseWriter) {
	// Avoid logging a DSN or unreviewed driver error containing private data.
	log.Print("database operation failed; check database availability and applied schema")
	reply(w, http.StatusServiceUnavailable, map[string]string{"error": "database unavailable"})
}

func decodeName(w http.ResponseWriter, r *http.Request) (string, error) {
	r.Body = http.MaxBytesReader(w, r.Body, 4096)
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	var input struct {
		Name string `json:"name"`
	}
	if err := decoder.Decode(&input); err != nil {
		return "", errors.New("invalid JSON")
	}
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		return "", errors.New("provide one JSON object")
	}
	name := strings.TrimSpace(input.Name)
	if name == "" || utf8.RuneCountInString(name) > 100 {
		return "", errors.New("name must contain 1-100 characters")
	}
	return name, nil
}

func newHandler(db *sql.DB) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /students", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()
		rows, err := db.QueryContext(ctx, "SELECT id, name FROM students ORDER BY id LIMIT 100")
		if err != nil {
			dbFailure(w)
			return
		}
		defer rows.Close()
		students := []Student{}
		for rows.Next() {
			var s Student
			if err := rows.Scan(&s.ID, &s.Name); err != nil {
				dbFailure(w)
				return
			}
			students = append(students, s)
		}
		if err := rows.Err(); err != nil {
			dbFailure(w)
			return
		}
		reply(w, http.StatusOK, students)
	})
	mux.HandleFunc("POST /students", func(w http.ResponseWriter, r *http.Request) {
		name, err := decodeName(w, r)
		if err != nil {
			reply(w, 400, map[string]string{"error": err.Error()})
			return
		}
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()
		s := Student{Name: name}
		// Input is bound as data through $1; never concatenate input into SQL.
		err = db.QueryRowContext(ctx, "INSERT INTO students (name) VALUES ($1) RETURNING id", name).Scan(&s.ID)
		if err != nil {
			dbFailure(w)
			return
		}
		reply(w, http.StatusCreated, s)
	})
	mux.HandleFunc("GET /sessions", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()
		rows, err := db.QueryContext(ctx, "SELECT id, title, capacity FROM class_sessions ORDER BY id LIMIT 100")
		if err != nil {
			dbFailure(w)
			return
		}
		defer rows.Close()
		sessions := []ClassSession{}
		for rows.Next() {
			var s ClassSession
			if err := rows.Scan(&s.ID, &s.Title, &s.Capacity); err != nil {
				dbFailure(w)
				return
			}
			sessions = append(sessions, s)
		}
		if err := rows.Err(); err != nil {
			dbFailure(w)
			return
		}
		reply(w, http.StatusOK, sessions)
	})
	return mux
}

func run() error {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return errors.New("DATABASE_URL is required; see README.md")
	}
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return errors.New("invalid database configuration")
	}
	defer db.Close()
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * time.Minute)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	err = db.PingContext(ctx)
	cancel()
	if err != nil {
		return errors.New("database startup check failed; check service and configuration")
	}
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = "127.0.0.1:8080"
	}
	server := &http.Server{Addr: addr, Handler: newHandler(db), ReadHeaderTimeout: 5 * time.Second, ReadTimeout: 10 * time.Second, WriteTimeout: 10 * time.Second, IdleTimeout: 60 * time.Second}
	stop, stopCancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stopCancel()
	serverResult := make(chan error, 1)
	go func() { serverResult <- server.ListenAndServe() }()
	log.Printf("local database teaching API listening on %s", addr)
	select {
	case err := <-serverResult:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	case <-stop.Done():
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			_ = server.Close()
			return err
		}
		return nil
	}
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
