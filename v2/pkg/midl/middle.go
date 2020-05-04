package midl

// Middleware defines a service that can be called by
// a midl.Adapter with a midl.Response to get a
// midl.Response.
type Middleware interface {

	// Handle handles an HTTP request.
	//
	// If Handle returns a non-nil value, this signals to the
	// midl.Adapter that this request has been processed.
	Handle(Request) Response
}

// MiddlewareFunc is a convenience wrapper to allow the use
// of an arbitrary function as an instance of Middleware.
//
//   handler := MiddlewareFunc(func(Request) Response {
//       return NewResponse().SetCode(http.StatusNoContent)
//   })
type MiddlewareFunc func(Request) Response

// Handle will call the wrapped function.
func (e MiddlewareFunc) Handle(r Request) Response {
	return e(r)
}
