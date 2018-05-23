package midl

// Empty response body serializer.
//
// Provides a chance to set default output in the event of
// an empty response return from middleware.
type EmptyHandler interface {
	Handle(Request, Response) []byte
}

type EmptyHandlerFunc func(Request, Response) []byte

func (e EmptyHandlerFunc) Handle(q Request, s Response) []byte {
	return e(q, s)
}

// Default no op empty handler
func DefaultEmptyHandler() EmptyHandler {
	return new(defaultEmptyHandler)
}

type defaultEmptyHandler struct{}

func (d defaultEmptyHandler) Handle(Request, Response) []byte {
	return nil
}
