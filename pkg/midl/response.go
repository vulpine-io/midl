package midl

import (
	"net/http"
)

// Response defines a builder that can be used to build
// an HTTP response.
type Response interface {

	// Body returns the body (if any) stored on the HTTP
	// response.
	Body() interface{}

	// Code returns the HTTP status code from the HTTP
	// response.
	Code() int

	// Error returns the error (if any) stored on the HTTP
	// response.
	Error() error

	// Header retrieves a header stored header value by key.
	Header(key string) string

	// Headers retrieves a slice of stored headers values by
	// key.
	Headers(key string) []string

	// SetBody stores the given value as the body value for
	// this response.
	SetBody(any interface{}) Response

	// SetCode stores the given value as the HTTP status code
	// for this response.
	SetCode(code int) Response

	// SetError stores the given value as the error value for
	// this response.
	SetError(error) Response

	// AddHeader creates or appends to the header values for
	// this response.
	AddHeader(key, value string) Response

	// AddedHeaders creates or appends a list of headers to
	// this response.
	AddHeaders(key string, value []string) Response

	// SetHeader creates or overwrites a header on this
	// response.
	SetHeader(key, value string) Response

	// SetHeaders creates or overwrites a list of headers on
	// this response.
	SetHeaders(key string, values []string) Response

	// RawHeaders grants access to the internal http.Header
	// map.
	RawHeaders() http.Header
}

// MakeResponse creates a Response instance with the given
// body and status code values.
func MakeResponse(code int, body interface{}) Response {
	return &response{
		code: code,
		body: body,
		head: make(http.Header),
	}
}

// MakeErrorResponse creates a Response instance with the
// given error and status code.
func MakeErrorResponse(code int, err error) Response {
	return &response{
		code:  code,
		error: err,
		head:  make(http.Header),
	}
}

// NewResponse creates a new Response instance.
// Default status code is 200 (OK).
func NewResponse() Response {
	return &response{
		code: http.StatusOK,
		head: make(http.Header),
	}
}

type response struct {
	body  interface{}
	code  int
	error error
	head  http.Header
}

func (d response) Body() interface{} {
	return d.body
}

func (d response) Code() int {
	return d.code
}

func (d response) Error() error {
	return d.error
}

func (d response) Header(key string) string {
	return d.head.Get(key)
}

func (d response) Headers(key string) []string {
	return d.head[key]
}

func (d *response) SetBody(any interface{}) Response {
	d.body = any
	return d
}

func (d *response) SetCode(code int) Response {
	d.code = code
	return d
}

func (d *response) SetError(err error) Response {
	d.error = err
	return d
}

func (d *response) AddHeader(key, value string) Response {
	d.head.Add(key, value)
	return d
}

func (d *response) AddHeaders(key string, value []string) Response {
	for _, v := range value {
		d.head.Add(key, v)
	}
	return d
}

func (d *response) SetHeader(key, value string) Response {
	d.head.Set(key, value)
	return d
}

func (d *response) SetHeaders(key string, values []string) Response {
	d.head.Del(key)
	return d.AddHeaders(key, values)
}

func (d response) RawHeaders() http.Header {
	return d.head
}
