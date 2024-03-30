package checker

type Runner struct {
	checkers []Checker
	handlers *Handlers
}

func NewRunner(handlers *Handlers) *Runner {
	return &Runner{
		checkers: make([]Checker, 0),
		handlers: handlers,
	}
}

func (r *Runner) AddChecker(checker Checker) {
	r.checkers = append(r.checkers, checker)
}

func (r *Runner) Run() {
	for _, checker := range r.checkers {
		go checkWithInterval(checker, r.handlers)
	}
}
