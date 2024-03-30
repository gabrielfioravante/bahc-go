// Package checker provides utilities for executing a series of checks concurrently and handling their results.
package checker

import "fmt"

type Result struct {
	Message string
	Success bool
}

func (r *Result) SetAvailable(id string) {
	r.Message = fmt.Sprintf("%s - Service available", id)
	r.Success = true
}

func (r *Result) SetUnavailable(id string, info string) {
	r.Message = fmt.Sprintf("%s - Service unavailable: %s", id, info)
	r.Success = false
}

type Checker interface {
	GetID() string
	Check() (Result, error)
}

type ResultHandler func(Result) error
type ErrorHandler func(error) error

func ExecuteCheckers(checkers []Checker, resultHandler ResultHandler, errorHandler ErrorHandler) {
	for _, checker := range checkers {
		go executeChecker(checker, resultHandler, errorHandler)
	}
}

func executeChecker(checker Checker, resultHandler ResultHandler, errorHandler ErrorHandler) {
	r, err := checker.Check()

	if err == nil {
		err = resultHandler(r)
	}

	if err != nil {
		if err = errorHandler(err); err != nil {
			panic(err)
		}
	}
}
