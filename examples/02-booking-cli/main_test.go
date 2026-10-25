package main

import "testing"

func TestReserve(t *testing.T) {
	cases := []struct {
		name                   string
		capacity, booked, want int
		wantError              bool
	}{
		{"available", 2, 0, 1, false},
		{"last seat", 2, 1, 2, false},
		{"full", 2, 2, 2, true},
		{"zero capacity", 0, 0, 0, true},
		{"negative booked", 2, -1, -1, true},
		{"overbooked input", 2, 3, 3, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := reserve(tc.capacity, tc.booked)
			if got != tc.want || (err != nil) != tc.wantError {
				t.Fatalf("reserve(%d, %d) = %d, %v; want %d, error=%v", tc.capacity, tc.booked, got, err, tc.want, tc.wantError)
			}
		})
	}
}

func TestFindSessionUsesID(t *testing.T) {
	sessions := []ClassSession{{ID: 1, Title: "Go"}, {ID: 2, Title: "Go"}}
	got, err := findSession(sessions, 2)
	if err != nil || got.ID != 2 {
		t.Fatalf("got %+v, %v", got, err)
	}
	if _, err := findSession(sessions, 99); err == nil {
		t.Fatal("missing ID should fail")
	}
}
