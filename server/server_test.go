package server

import (
	"net/http"
	"testing"
)

func TestNewStaticServerMux(t *testing.T) {
	mux := NewStaticServerMux("../static")
	if mux == nil {
		t.Error("NewStaticServerMux() returned nil")
	}
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		t.Error("ListenAndServe() returned error:", err)
	}
}
