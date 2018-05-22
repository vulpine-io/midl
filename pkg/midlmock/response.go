package midlmock

import (
	"net/http"
	"github.com/foxcapades/go-midl/pkg/midl"
)

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
}

func (r Response) Body() interface{} {
	return r.BodyFunc()
}

func (r Response) Code() int {
	return r.CodeFunc()
}

func (r Response) Error() error {
	return r.ErrorFunc()
}

func (r Response) Header(key string) string {
	return r.HeaderFunc(key)
}

func (r Response) Headers(key string) []string {
	return r.HeadersFunc(key)
}

func (r *Response) SetBody(any interface{}) midl.Response {
	r.SetBodyFunc(any)
	return r
}

func (r *Response) SetCode(code int) midl.Response {
	r.SetCodeFunc(code)
	return r
}

func (r *Response) SetError(e error) midl.Response {
	r.SetErrorFunc(e)
	return r
}

func (r *Response) AddHeader(key, value string) midl.Response {
	r.AddHeaderFunc(key, value)
	return r
}

func (r *Response) AddHeaders(key string, value []string) midl.Response {
	r.AddHeadersFunc(key, value)
	return r
}

func (r *Response) SetHeader(key, value string) midl.Response {
	r.SetHeaderFunc(key, value)
	return r
}

func (r *Response) SetHeaders(key string, values []string) midl.Response {
	r.SetHeadersFunc(key, values)
	return r
}

func (r Response) RawHeaders() http.Header {
	return r.RawHeadersFunc()
}



