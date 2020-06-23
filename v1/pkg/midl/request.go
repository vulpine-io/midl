package midl

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// Request defines a wrapper/accessor for the default Golang
// http.Request struct.
type Request interface {
	// Header gets the first header stored at the given key.
	Header(key string) (value string, ok bool)

	// Headers gets all headers stored at the given key.
	Headers(key string) (values []string, ok bool)

	// Body retrieves the request body bytes.
	//
	// In the event of a read error, returns nil; the error
	// will be available by calling the Error method.
	//
	// Multiple calls will only read the http.Request.Body
	// reader once.
	Body() []byte

	// Host retrieves the host string from the request.
	Host() string

	// Parameter get the query parameter stored at a given
	// key.
	Parameter(key string) (value string, ok bool)

	// Parameters gets all query parameters stored at a given
	// key.
	Parameters(key string) (values []string, ok bool)

	// RawRequest retrieves the raw Go standard library
	// request.
	//
	// Warning: Modifying the raw http.Request may impact the
	// output of other Request methods.
	RawRequest() *http.Request

	// Error retrieves the last error (if any) encountered by
	// method calls to Request.
	// If no error is present, returns nil.
	Error() error

	// ProcessBody runs the given BodyProcessor against the
	// body bytes of this request.
	//
	// Errors emitted by the body processor will be
	// retrievable from the Error method.
	//
	// Calls to this method will do nothing if an error has
	// been previously encountered.
	ProcessBody(BodyProcessor) Request

	// AdditionalContext returns a map for use in assigning
	// additional arbitrary context data to a request.
	AdditionalContext() map[interface{}]interface{}
}

// NewRequest constructs a new instance of midl.Request
// wrapping a provided http.Request.  If the provided
// request pointer is nil, an error is returned instead of
// a Request instance.
func NewRequest(r *http.Request) (Request, error) {
	if r == nil {
		return nil, ErrWrappedNil
	}

	return &request{raw: r, ctx: map[interface{}]interface{}{}}, nil
}

type request struct {
	raw     *http.Request
	error   error
	body    []byte
	hasBody bool
	ctx     map[interface{}]interface{}
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

func (r *request) AdditionalContext() map[interface{}]interface{} {
	return r.ctx
}
