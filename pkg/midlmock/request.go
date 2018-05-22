package midlmock

import (
	"github.com/foxcapades/go-midl/pkg/midl"
	"net/http"
)

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

func (r Request) Header(key string) (string, bool) {
	return r.HeaderFunc(key)
}

func (r Request) Headers(key string) ([]string, bool) {
	return r.HeadersFunc(key)
}

func (r Request) Body() []byte {
	return r.BodyFunc()
}

func (r Request) Host() string {
	return r.HostFunc()
}

func (r Request) Parameter(key string) (string, bool) {
	return r.ParameterFunc(key)
}

func (r Request) Parameters(key string) ([]string, bool) {
	return r.ParametersFunc(key)
}

func (r Request) RawRequest() *http.Request {
	return r.RawRequestFunc()
}

func (r Request) Error() error {
	return r.ErrorFunc()
}

func (r *Request) ProcessBody(in midl.BodyProcessor) midl.Request {
	r.ProcessBodyFunc(in)
	return r
}
