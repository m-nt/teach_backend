package main

import (
	"errors"
	"fmt"
)

type ClassSession struct {
	ID       int
	Title    string
	Capacity int
	Booked   int
}

// reserve returns the new count. It does not mutate state on failure.
func reserve(capacity, booked int) (int, error) {
	if capacity <= 0 || booked < 0 || booked > capacity {
		return booked, errors.New("invalid capacity or booked count")
	}
	if booked == capacity {
		return booked, errors.New("class is full")
	}
	return booked + 1, nil
}

func findSession(sessions []ClassSession, id int) (ClassSession, error) {
	for _, session := range sessions {
		if session.ID == id {
			return session, nil
		}
	}
	return ClassSession{}, fmt.Errorf("session %d not found", id)
}

func main() {
	sessions := []ClassSession{
		{ID: 1, Title: "Go fundamentals", Capacity: 2},
		{ID: 2, Title: "Go fundamentals", Capacity: 4},
	}
	session, err := findSession(sessions, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	for attempt := 1; attempt <= 3; attempt++ {
		next, err := reserve(session.Capacity, session.Booked)
		if err != nil {
			fmt.Printf("attempt %d: rejected (%s)\n", attempt, err)
			continue
		}
		session.Booked = next
		fmt.Printf("attempt %d: accepted; remaining=%d\n", attempt, session.Capacity-session.Booked)
	}
}
