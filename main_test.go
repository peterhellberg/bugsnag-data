package main

import (
	"net/http"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	key := "foo"

	c := newClient(key)

	if got, want := c.key, key; got != want {
		t.Fatalf("c.key = %q, want %q", got, want)
	}

	if got, want := c.Timeout, 10*time.Second; got != want {
		t.Fatalf("c.Timeout = %v, want %v", got, want)
	}
}

func TestClientRequest(t *testing.T) {
	for _, tt := range []struct {
		key    string
		rawurl string
	}{
		{"foo", "https://example.com/foo"},
		{"bar", "http://example.org/bar"},
	} {
		t.Run(tt.key, func(t *testing.T) {
			c := newClient(tt.key)

			req, err := c.request(tt.rawurl)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got, want := req.Method, http.MethodGet; got != want {
				t.Fatalf("req.Method = %q, want %q", got, want)
			}

			if got, want := req.Header.Get("Authorization"), "token "+tt.key; got != want {
				t.Fatalf("req.Header.Get(\"Authorization\") = %q, want %q", got, want)
			}
		})
	}
}
