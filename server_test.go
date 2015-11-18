package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Simple server test

func TestHandler(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(successHandler))

	defer ts.Close()

	// Make the HTTP requset to the test server
	req, err := http.NewRequest("GET", ts.URL, nil)
	w := httptest.NewRecorder()
	assert.Nil(t, err)

	// calling handler directly
	successHandler(w, req)

	// read the body of the HTTP response
	assert.Equal(t, "OK", w.Body.String())
}

func TestFailure(t *testing.T) {
	// Suppose service is unavailable, and module depends on
	// service. Then assert that module fails too
	ts := httptest.NewServer(http.HandlerFunc(failHandler))
	defer ts.Close()

	err := getter(ts.URL)
	assert.NotNil(t, err)
}
