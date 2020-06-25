package midl

// RequestWrapper defines a middleware layer that should be
// called both before and after a request is handled.
//
// Useful for logging or metric gathering.
type RequestWrapper interface {

	// Request takes an incoming request instance and returns
	// a request.
	//
	// The returned request will be passed on to additional
	// middleware/handlers.
	Request(Request) Request

	// Response takes a request and an outgoing response and
	// returns a response.
	Response(Request, Response) Response
}
