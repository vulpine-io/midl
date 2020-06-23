package midlmock

import (
	"net/http"

	"github.com/vulpine-io/midl/v1/pkg/midl"
)

// Response is a configurable mock implementation of the
// midl.Response interface.
type Response struct {
	BodyFunc    func() interface{}
	CodeFunc    func() int
	ErrorFunc   func() error
	HeaderFunc  func(key string) string
	HeadersFunc func(key string) []string

	SetBodyFunc    func(any interface{})
	SetCodeFunc    func(code int)
	SetErrorFunc   func(error)
	AddHeaderFunc  func(key, value string)
	AddHeadersFunc func(key string, value []string)
	SetHeaderFunc  func(key, value string)
	SetHeadersFunc func(key string, values []string)
	RawHeadersFunc func() http.Header
	CallbackFunc   func(f func())
	CallbacksFunc  func() []func()
}

func (r *Response) Callback(f func()) midl.Response {
	r.CallbackFunc(f)
	return r
}

func (r Response) Callbacks() []func() {
	return r.CallbacksFunc()
}

// Body is a passthrough for the function stored at the
// Response.BodyFunc property.
func (r Response) Body() interface{} {
	return r.BodyFunc()
}

// Code is a passthrough for the function stored at the
// Response.CodeFunc property.
func (r Response) Code() int {
	return r.CodeFunc()
}

// Error is a passthrough for the function stored at the
// Response.ErrorFunc property.
func (r Response) Error() error {
	return r.ErrorFunc()
}

// Header is a passthrough for the function stored at the
// Response.HeaderFunc property.
func (r Response) Header(key string) string {
	return r.HeaderFunc(key)
}

// Headers is a passthrough for the function stored at the
// Response.HeadersFunc property.
func (r Response) Headers(key string) []string {
	return r.HeadersFunc(key)
}

// SetBody is a passthrough for the function stored at the
// Response.SetBodyFunc property.
// Returns the current Response instance.
func (r *Response) SetBody(any interface{}) midl.Response {
	r.SetBodyFunc(any)
	return r
}

// SetCode is a passthrough for the function stored at the
// Response.SetCodeFunc property.
// Returns the current Response instance.
func (r *Response) SetCode(code int) midl.Response {
	r.SetCodeFunc(code)
	return r
}

// SetError is a passthrough for the function stored at the
// Response.SetErrorFunc property.
// Returns the current Response instance.
func (r *Response) SetError(e error) midl.Response {
	r.SetErrorFunc(e)
	return r
}

// AddHeader is a passthrough for the function stored at the
// Response.AddHeaderFunc property.
// Returns the current Response instance.
func (r *Response) AddHeader(key, value string) midl.Response {
	r.AddHeaderFunc(key, value)
	return r
}

// AddHeaders is a passthrough for the function stored at
// the Response.AddHeadersFunc property.
// Returns the current Response instance.
func (r *Response) AddHeaders(key string, value []string) midl.Response {
	r.AddHeadersFunc(key, value)
	return r
}

// SetHeader is a passthrough for the function stored at the
// Response.SetHeaderFunc property.
// Returns the current Response instance.
func (r *Response) SetHeader(key, value string) midl.Response {
	r.SetHeaderFunc(key, value)
	return r
}

// SetHeaders is a passthrough for the function stored at
// the Response.SetHeadersFunc property.
// Returns the current Response instance.
func (r *Response) SetHeaders(key string, values []string) midl.Response {
	r.SetHeadersFunc(key, values)
	return r
}

// RawHeaders is a passthrough for the function stored at
// the Response.RawHeadersFunc property.
func (r Response) RawHeaders() http.Header {
	return r.RawHeadersFunc()
}
