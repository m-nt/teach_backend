package main

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDecodeName(t *testing.T) {
	for _, tc := range []struct {
		name, body, want string
		invalid          bool
	}{
		{"trim", `{"name":" Ada "}`, "Ada", false},
		{"unicode", `{"name":"سارا"}`, "سارا", false},
		{"blank", `{"name":" "}`, "", true},
		{"malformed", `{`, "", true},
		{"unknown", `{"name":"Ada","role":"admin"}`, "", true},
		{"trailing", `{"name":"Ada"} {}`, "", true},
		{"too long", `{"name":"` + strings.Repeat("a", 101) + `"}`, "", true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			r := httptest.NewRequest("POST", "/students", strings.NewReader(tc.body))
			got, err := decodeName(httptest.NewRecorder(), r)
			if got != tc.want || (err != nil) != tc.invalid {
				t.Fatalf("got %q, %v", got, err)
			}
		})
	}
}
