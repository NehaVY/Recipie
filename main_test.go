package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestSearchHandler tests the search functionality for various queries
func TestSearchHandler(t *testing.T) {
	tests := []struct {
		query          string
		expectedStatus int
		expectedBody   string
	}{
		{
			query:          "espresso",
			expectedStatus: http.StatusOK,
			expectedBody:   "Espresso",
		},
		{
			query:          "americano",
			expectedStatus: http.StatusOK,
			expectedBody:   "Americano",
		},
		{
			query:          "latte",
			expectedStatus: http.StatusOK,
			expectedBody:   "Latte",
		},
		{
			query:          "unknown",
			expectedStatus: http.StatusOK,
			expectedBody:   "Recipe not found",
		},
		{
			query:          "matcha",
			expectedStatus: http.StatusOK,
			expectedBody:   "Recipe not found", // Testing failure case
		},
	}

	for _, test := range tests {
		req, err := http.NewRequest("GET", "/search?query="+test.query, nil)
		if err != nil {
			t.Fatalf("Could not create request: %v", err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(searchHandler)
		handler.ServeHTTP(rr, req)

		if rr.Code != test.expectedStatus {
			t.Errorf("Expected status %v; got %v", test.expectedStatus, rr.Code)
		}

		if !strings.Contains(rr.Body.String(), test.expectedBody) {
			t.Errorf("For query %q, expected body to contain %q; got %q", test.query, test.expectedBody, rr.Body.String())
		}
	}
}

// TestServeExistingFile tests serving existing static files
func TestServeExistingFile(t *testing.T) {
	req, err := http.NewRequest("GET", "/static/example.html", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.FileServer(http.Dir("./static")) // Change "./static" to your actual directory
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status %v; got %v", http.StatusOK, rr.Code)
	}

	if !strings.Contains(rr.Body.String(), "<html>") {
		t.Errorf("Expected HTML content in response body; got %q", rr.Body.String())
	}
}

// TestHandleNonExistentFile tests handling non-existent files
func TestHandleNonExistentFile(t *testing.T) {
	req, err := http.NewRequest("GET", "/static/non-existent-file.txt", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.FileServer(http.Dir("./static")) // Change "./static" to your actual directory
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("Expected status %v for non-existent file; got %v", http.StatusNotFound, rr.Code)
	}
}

// TestServeFileWithCorrectMIMEType tests serving files with correct MIME types
func TestServeFileWithCorrectMIMEType(t *testing.T) {
	tests := []struct {
		filePath     string
		expectedMIME string
	}{
		{"/static/example.html", "text/html; charset=utf-8"},
		{"/static/example.json", "application/json"},
		{"/static/example.png", "image/png"},
	}

	for _, test := range tests {
		req, err := http.NewRequest("GET", test.filePath, nil)
		if err != nil {
			t.Fatalf("Could not create request: %v", err)
		}

		rr := httptest.NewRecorder()
		handler := http.FileServer(http.Dir("./static")) // Change "./static" to your actual directory
		handler.ServeHTTP(rr, req)

		if contentType := rr.Header().Get("Content-Type"); contentType != test.expectedMIME {
			t.Errorf("For file %q, expected MIME type %q; got %q", test.filePath, test.expectedMIME, contentType)
		}
	}
}
