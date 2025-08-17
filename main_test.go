package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSayHello(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	sayHello(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	expected := "Hello, CI/CD World!"
	if w.Body.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, w.Body.String())
	}
}
