package checker

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type MockChecker struct {
	result string
	error  string
}

func (mc *MockChecker) Run() (Result, error) {
	return Result{Message: "test message", Success: true}, nil
}

func (mc *MockChecker) GetID() string {
	return "Mock ID"
}

func (mc *MockChecker) GetInterval() time.Duration {
	return 0
}

func TestCheck(t *testing.T) {
	mockChecker := &MockChecker{}

	resultHandler := func(r Result) error {
		mockChecker.result = r.Message
		return errors.New("test error")
	}

	errorHandler := func(err error) error {
		mockChecker.error = err.Error()
		return nil
	}

	handlers := &Handlers{
		Result: resultHandler,
		Error:  errorHandler,
	}

	check(mockChecker, handlers)

	assert.Equal(t, mockChecker.result, "test message", "result message mismatch")
	assert.Equal(t, mockChecker.error, "test error", "error message mismatch")
}
