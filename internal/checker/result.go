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
