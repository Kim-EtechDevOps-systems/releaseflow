package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthEndpoint(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet, "/health", nil,
	)
	recorder := httptest.NewRecorder()

	newHandler().ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", recorder.Code)
	}

	var body map[string]string
	if err := json.Unmarshal(recorder.Body.Bytes(), &body); err != nil {
		t.Fatal(err)
	}

	if body["status"] != "unhealthy" {
		t.Errorf("unexpected status: %q", body["status"])
	}
}

func TestUnknownEndpoint(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet, "/unknown", nil,
	)
	recorder := httptest.NewRecorder()

	newHandler().ServeHTTP(recorder, req)

	if recorder.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", recorder.Code)
	}
}
