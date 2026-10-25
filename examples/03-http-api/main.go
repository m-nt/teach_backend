package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type Session struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Capacity int    `json:"capacity"`
	Booked   int    `json:"booked"`
}

type BookingRequest struct {
	SessionID   int    `json:"session_id"`
	StudentName string `json:"student_name"`
}

type app struct {
	mu            sync.Mutex // Protect both the availability check and the update.
	session       Session
	nextBookingID int
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(value); err != nil {
		log.Print("response write failed")
	}
}

func fail(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

func newHandler() http.Handler {
	a := &app{session: Session{ID: 1, Title: "Go fundamentals", Capacity: 2}, nextBookingID: 1}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /sessions", func(w http.ResponseWriter, r *http.Request) {
		a.mu.Lock()
		snapshot := a.session
		a.mu.Unlock()
		writeJSON(w, http.StatusOK, []Session{snapshot})
	})
	mux.HandleFunc("POST /bookings", a.book)
	return mux
}

func (a *app) book(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 4096)
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	var input BookingRequest
	if err := decoder.Decode(&input); err != nil {
		fail(w, 400, "invalid JSON body")
		return
	}
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		fail(w, 400, "provide one JSON object")
		return
	}
	input.StudentName = strings.TrimSpace(input.StudentName)
	if input.SessionID <= 0 || input.StudentName == "" || len(input.StudentName) > 100 {
		fail(w, 400, "positive session_id and student_name (1-100 bytes) required")
		return
	}
	a.mu.Lock()
	if input.SessionID != a.session.ID {
		a.mu.Unlock()
		fail(w, 404, "session not found")
		return
	}
	if a.session.Booked >= a.session.Capacity {
		a.mu.Unlock()
		fail(w, 409, "class is full")
		return
	}
	a.session.Booked++
	id := a.nextBookingID
	a.nextBookingID++
	a.mu.Unlock()
	writeJSON(w, 201, map[string]int{"id": id, "session_id": input.SessionID})
}

func main() {
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = "127.0.0.1:8080"
	}
	server := &http.Server{Addr: addr, Handler: newHandler(), ReadHeaderTimeout: 5 * time.Second, ReadTimeout: 10 * time.Second, WriteTimeout: 10 * time.Second, IdleTimeout: 60 * time.Second}
	log.Printf("local teaching API listening on %s; in-memory data only", addr)
	log.Fatal(server.ListenAndServe())
}
