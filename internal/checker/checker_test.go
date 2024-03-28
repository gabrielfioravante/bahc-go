package checker

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockChecker struct {
	result string
	error  string
}

func (mc *MockChecker) Check() (Result, error) {
	return Result{Message: "test message", Success: true}, nil
}

func TestExecuteChecker(t *testing.T) {
	mockChecker := &MockChecker{}

	resultHandler := func(r Result) error {
		mockChecker.result = r.Message
		return errors.New("test error")
	}

	errorHandler := func(err error) error {
		mockChecker.error = err.Error()
		return nil
	}

	executeChecker(mockChecker, resultHandler, errorHandler)

	assert.Equal(t, mockChecker.result, "test message", "result message mismatch")
	assert.Equal(t, mockChecker.error, "test error", "error message mismatch")
}
