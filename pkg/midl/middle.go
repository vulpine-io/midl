package midl

type Middleware interface {
	Handle(Request) Response
}

type MiddlewareFunc func(Request) Response

func (e MiddlewareFunc) Handle(r Request) Response {
	return e(r)
}
