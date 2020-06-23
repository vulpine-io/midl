package midl

import (
	"net/http"
)

type writer = http.ResponseWriter
type header = http.Header

// Adapter for conversion between Middleware and Golang
// http.Handlers.
//
//   handler := JSONAdapter(NewInputValidator(), ..., NewResponder())
//   http.Handle("/", handler)
//   log.Fatal(http.ListenAndServe(":8080", nil))
type Adapter interface {
	http.Handler

	// EmptyHandler registers a handler for empty response
	// bodies.
	//
	// Defaults to an empty byte array
	EmptyHandler(EmptyHandler) Adapter

	// Content-Type sets the default content type header.
	ContentType(string) Adapter

	// ErrorSerializer registers a handler for serializing
	// errors.
	ErrorSerializer(ErrorSerializer) Adapter

	// Serializer registers the default body serializer.
	Serializer(Serializer) Adapter

	// AddHandlers appends handlers to the list of Middleware
	// handlers.
	AddHandlers(...Middleware) Adapter

	// SetHandlers set and/or overwrite the current list of
	// Middleware handlers.
	SetHandlers(...Middleware) Adapter
}
