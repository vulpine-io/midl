package midl

type EmptyHandler interface {
	Handle(Request, Response) []byte
}

type EmptyHandlerFunc func(Request, Response) []byte

func (e EmptyHandlerFunc) Handle(q Request, s Response) []byte {
	return e(q, s)
}

func DefaultEmptyHandler() EmptyHandler {
	return new(defaultEmptyHandler)
}

type defaultEmptyHandler struct{}

func (d defaultEmptyHandler) Handle(Request, Response) []byte {
	return nil
}

