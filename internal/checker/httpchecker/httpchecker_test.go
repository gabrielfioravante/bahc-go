package httpchecker

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHttpChecker_Run(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	defer server.Close()

	config := &HttpCheckerConfig{
		ID:     "test-check",
		Url:    server.URL,
		Method: GET,
	}

	httpChecker := NewHttpChecker(config)
	result, err := httpChecker.Run()

	assert.NoError(t, err)
	assert.Equal(t, "test-check - Service available", result.Message)
	assert.Equal(t, result.Success, true)
}

func TestHttpChecker_Run_Failed(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))

	defer server.Close()

	config := &HttpCheckerConfig{
		ID:     "test-check",
		Url:    server.URL,
		Method: GET,
	}

	httpChecker := NewHttpChecker(config)
	result, err := httpChecker.Run()

	assert.NoError(t, err)
	assert.Equal(t, "test-check - Service unavailable: 500 Internal Server Error", result.Message)
	assert.Equal(t, result.Success, false)
}

func TestHttpChecker_Run_Timeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(250 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))

	defer server.Close()

	config := &HttpCheckerConfig{
		ID:      "test-check",
		Url:     server.URL,
		Method:  GET,
		Timeout: 200 * time.Millisecond,
	}

	httpChecker := NewHttpChecker(config)
	result, err := httpChecker.Run()

	assert.NoError(t, err)
	assert.NotEqual(t, result.Message, "")
	assert.Equal(t, result.Success, false)
}
