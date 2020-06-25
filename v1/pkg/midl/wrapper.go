package midl

// RequestWrapper defines a middleware layer that should be
// called both before and after a request is handled.
//
// Useful for logging or metric gathering.
type RequestWrapper interface {

	// Request takes an incoming request instance and returns
	// a request.
	Request(Request)

	// Response takes a request and an outgoing response and
	// returns a response.
	//
	// The response returned by this method will be used when
	// calling the next wrapper (or if no more wrappers exist
	// it will be serialized and returned to the client).
	Response(Request, Response) Response
}
