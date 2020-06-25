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

	// AddWrappers appends request wrappers to the current list of wrappers.
	//
	// Wrappers are executed in the following order:
	//
	// * For requests: Wrappers will be called in the order they were appended,
	// but before any Middleware instances are called.
	// * For responses: Wrappers will be called in reverse from the order they
	// were appended.
	AddWrappers(...RequestWrapper) Adapter

	// AddWrappers sets the list of request wrappers to the given input,
	// overwriting the previous wrapper list with the given list.
	//
	// Wrappers are executed in the following order:
	//
	// * For requests: Wrappers will be called in the order they appeared in the
	// input to this function, but before any Middleware instances are called.
	// * For responses: Wrappers will be called in reverse from the order they
	// appeared in the input to this function.
	SetWrappers(...RequestWrapper) Adapter
}
