package midl

import (
	"net/http"
	"io/ioutil"
	"bytes"
)

// Inbound HTTP request
type Request interface {
	// Get the first header stored at the given key
	Header(key string) (value string, ok bool)

	// Get all headers stored at the given key
	Headers(key string) (values []string, ok bool)

	// Get the request body bytes.
	//
	// In the event of a read error, returns nil; the error
	// will be available by calling the Error method.
	//
	// Safe for multiple calls.
	Body() []byte
	Host() string

	// Get the query parameter stored at a given key
	Parameter(key string) (value string, ok bool)

	// Get all query parameters stored at a given key
	Parameters(key string) (values []string, ok bool)

	// Retrieve raw Go standard library request
	// Warning: Modifying this may affect the output of other
	// Request methods.
	RawRequest() *http.Request

	// Retrieve the last error (if any) encountered by method
	// calls to Request.
	// If no error is present, returns nil.
	Error() error

	// Run the given BodyProcessor against the body bytes of
	// this request.
	//
	// Errors emitted by the body processor will be
	// retrievable from the Error method.
	//
	// Calls to this method will do nothing if an error has
	// been previously encountered.
	ProcessBody(BodyProcessor) Request
}

func NewRequest(r *http.Request) (Request, error) {
	if r == nil {
		return nil, ErrorWrappedNil
	}

	return &request{raw: r}, nil
}

type request struct {
	raw     *http.Request
	error   error
	body    []byte
	hasBody bool
}

func (r *request) readBody() {
	if r.error != nil || r.hasBody {
		return
	}

	body, err := ioutil.ReadAll(r.raw.Body)
	if err != nil {
		r.error = err
	} else {
		r.body = body
	}
	r.hasBody = true
	r.raw.Body = ioutil.NopCloser(bytes.NewBuffer(r.body))
}

func (r *request) Header(key string) (string, bool) {
	values, ok := r.raw.Header[key]
	if !ok {
		return "", false
	}
	return values[0], ok
}

func (r *request) Headers(key string) ([]string, bool) {
	values, ok := r.raw.Header[key]
	return values, ok
}

func (r *request) Body() []byte {
	r.readBody()
	return r.body
}

func (r *request) Host() string {
	return r.raw.Host
}

func (r *request) Parameter(key string) (string, bool) {
	val, ok := r.raw.URL.Query()[key]
	if !ok {
		return "", false
	}
	return val[0], true
}

func (r *request) Parameters(key string) ([]string, bool) {
	val, ok := r.raw.URL.Query()[key]
	return val, ok
}

func (r *request) RawRequest() *http.Request {
	return r.raw
}

func (r *request) Error() error {
	return r.error
}

func (r *request) ProcessBody(processor BodyProcessor) Request {
	if r.error == nil && processor != nil {
		r.error = processor.Process(r.Body())
	}

	return r
}
