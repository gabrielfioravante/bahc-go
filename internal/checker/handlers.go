package checker

type ResultHandler func(Result) error
type ErrorHandler func(error) error

type Handlers struct {
	Result ResultHandler
	Error  ErrorHandler
}
