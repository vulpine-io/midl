package midl

// An EmptyHandler provides a chance to set default response
// output in the event of an empty body from middleware.
type EmptyHandler interface {
	Handle(Request, Response) []byte
}

// EmptyHandlerFunc provides a function wrapper for simple
// EmptyHandlers.
type EmptyHandlerFunc func(Request, Response) []byte

func (e EmptyHandlerFunc) Handle(q Request, s Response) []byte {
	return e(q, s)
}

// DefaultEmptyHandler returns a no op empty handler
func DefaultEmptyHandler() EmptyHandler {
	return new(defaultEmptyHandler)
}

type defaultEmptyHandler struct{}

func (d defaultEmptyHandler) Handle(Request, Response) []byte {
	return nil
}
