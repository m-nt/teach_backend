package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
)

func post(h http.Handler, body string) *httptest.ResponseRecorder {
	r := httptest.NewRequest("POST", "/bookings", strings.NewReader(body))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	return w
}

func TestInvalidRequestsDoNotConsumeSeats(t *testing.T) {
	h := newHandler()
	for _, tc := range []struct {
		body   string
		status int
	}{
		{`{`, 400},
		{`{"session_id":1,"student_name":" "}`, 400},
		{`{"session_id":99,"student_name":"Ada"}`, 404},
		{`{"session_id":1,"student_name":"Ada","extra":true}`, 400},
		{`{"session_id":1,"student_name":"Ada"} {}`, 400},
		{`{"session_id":1,"student_name":"` + strings.Repeat("a", 5000) + `"}`, 400},
	} {
		if w := post(h, tc.body); w.Code != tc.status {
			t.Fatalf("got %d: %s", w.Code, w.Body)
		}
	}
	for i := 0; i < 2; i++ {
		if w := post(h, `{"session_id":1,"student_name":"Ada"}`); w.Code != 201 {
			t.Fatalf("seat lost: %d", w.Code)
		}
	}
	if w := post(h, `{"session_id":1,"student_name":"Ada"}`); w.Code != 409 {
		t.Fatalf("full class status = %d", w.Code)
	}
}

func TestConcurrentRequestsRespectCapacity(t *testing.T) {
	h := newHandler()
	results := make(chan int, 10)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			results <- post(h, fmt.Sprintf(`{"session_id":1,"student_name":"student-%d"}`, i)).Code
		}(i)
	}
	wg.Wait()
	close(results)
	accepted, rejected := 0, 0
	for status := range results {
		switch status {
		case 201:
			accepted++
		case 409:
			rejected++
		default:
			t.Fatalf("unexpected %d", status)
		}
	}
	if accepted != 2 || rejected != 8 {
		t.Fatalf("accepted=%d rejected=%d", accepted, rejected)
	}
}
