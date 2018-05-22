package midl

import (
	"net/http"
)

// HTTP Response builder
type Response interface {

	// Retrieve the body (if any) stored on the HTTP response.
	Body() interface{}

	// Retrieve HTTP status code from the HTTP response.
	// Defaults to 200
	Code() int

	// Retrieve the error (if any) stored on the HTTP response.
	Error() error

	// Retrieve a header stored header value by key.
	Header(key string) string

	// Retrieve stored headers values by key.
	Headers(key string) []string

	SetBody(any interface{}) Response
	SetCode(code int) Response
	SetError(error) Response
	AddHeader(key, value string) Response
	AddHeaders(key string, value []string) Response
	SetHeader(key, value string) Response
	SetHeaders(key string, values []string) Response
	RawHeaders() http.Header
}

func MakeResponse(code int, body interface{}) Response {
	return &response{
		code: code,
		body: body,
		head: make(http.Header),
	}
}

func MakeErrorResponse(code int, err error) Response {
	return &response{
		code: code,
		error: err,
		head: make(http.Header),
	}
}

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
