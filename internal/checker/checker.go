// Package checker provides utilities for executing a series of checks concurrently and handling their results.
package checker

import "time"

type Checker interface {
	GetID() string
	Run() (Result, error)
	GetInterval() time.Duration
}

func check(checker Checker, handlers *Handlers) {
	r, err := checker.Run()

	if err == nil {
		err = handlers.Result(r)
	}

	if err != nil {
		if err = handlers.Error(err); err != nil {
			panic(err)
		}
	}
}

func checkWithInterval(checker Checker, handlers *Handlers) {
	interval := checker.GetInterval()

	if interval <= 0 {
		check(checker, handlers)
		return
	}

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			check(checker, handlers)
		}
	}
}
