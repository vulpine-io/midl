package midlmock

import (
	"net/http"

	"gopkg.in/foxcapades/go-midl.v1/pkg/midl"
)

// Request is a configurable mock implementation of the
// midl.Request interface.
type Request struct {
	HeaderFunc     func(string) (string, bool)
	HeadersFunc    func(string) ([]string, bool)
	BodyFunc       func() []byte
	HostFunc       func() string
	ParameterFunc  func(string) (string, bool)
	ParametersFunc func(string) ([]string, bool)
	RawRequestFunc func() *http.Request
	ErrorFunc      func() error

	ProcessBodyFunc func(midl.BodyProcessor)
}

// Header is a passthrough for the function stored at the
// Request.HeaderFunc property.
func (r Request) Header(key string) (string, bool) {
	return r.HeaderFunc(key)
}

// Headers is a passthrough for the function stored at the
// Request.HeadersFunc property.
func (r Request) Headers(key string) ([]string, bool) {
	return r.HeadersFunc(key)
}

// Body is a passthrough for the function stored at the
// Request.BodyFunc property.
func (r Request) Body() []byte {
	return r.BodyFunc()
}

// Host is a passthrough for the function stored at the
// Request.HostFunc property.
func (r Request) Host() string {
	return r.HostFunc()
}

// Parameter is a passthrough for the function stored at the
// Request.ParameterFunc property.
func (r Request) Parameter(key string) (string, bool) {
	return r.ParameterFunc(key)
}

// Parameters is a passthrough for the function stored at
// the Request.ParametersFunc property.
func (r Request) Parameters(key string) ([]string, bool) {
	return r.ParametersFunc(key)
}

// RawRequest is a passthrough for the function stored at
// the Request.RawRequestFunc property.
func (r Request) RawRequest() *http.Request {
	return r.RawRequestFunc()
}

// Error is a passthrough for the function stored at the
// Request.ErrorFunc property.
func (r Request) Error() error {
	return r.ErrorFunc()
}

// ProcessBody is a passthrough for the function stored at
// the Request.ProcessBodyFunc property.
// Returns the current Request instance.
func (r *Request) ProcessBody(in midl.BodyProcessor) midl.Request {
	r.ProcessBodyFunc(in)
	return r
}
