package midlmock

import (
	"net/http"

	"gopkg.in/foxcapades/go-midl.v1/pkg/midl"
)

// Adapter is a configurable mock implementation of the
// midl.Adapter interface.
type Adapter struct {
	ServeHTTPFunc       func(http.ResponseWriter, *http.Request)
	EmptyHandlerFunc    func(midl.EmptyHandler)
	ContentTypeFunc     func(string)
	ErrorSerializerFunc func(midl.ErrorSerializer)
	SerializerFunc      func(midl.Serializer)
	AddHandlerFunc      func(...midl.Middleware)
	SetHandlerFunc      func(...midl.Middleware)
}

// ServeHTTP is a passthrough for the function stored in the
// Adapter.ServeHTTPFunc property.
func (a *Adapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.ServeHTTPFunc(w, r)
}

// EmptyHandler is a passthrough for the function stored in
// the Adapter.EmptyHandlerFunc property.
// Returns the current Adapter instance.
func (a *Adapter) EmptyHandler(in midl.EmptyHandler) midl.Adapter {
	a.EmptyHandlerFunc(in)
	return a
}

// ContentType is a passthrough for the function stored in
// the Adapter.ContentTypeFunc property.
// Returns the current Adapter instance.
func (a *Adapter) ContentType(in string) midl.Adapter {
	a.ContentTypeFunc(in)
	return a
}

// ErrorSerializer is a passthrough for the function stored
// in the Adapter.ErrorSerializerFunc property.
// Returns the current Adapter instance.
func (a *Adapter) ErrorSerializer(in midl.ErrorSerializer) midl.Adapter {
	a.ErrorSerializerFunc(in)
	return a
}

// Serializer is a passthrough for the function stored in
// the Adapter.SerializerFunc property.
// Returns the current Adapter instance.
func (a *Adapter) Serializer(in midl.Serializer) midl.Adapter {
	a.SerializerFunc(in)
	return a
}

// AddHandlers is a passthrough for the function stored in
// the Adapter.AddHandlersFunc property.
// Returns the current Adapter instance.
func (a *Adapter) AddHandlers(in ...midl.Middleware) midl.Adapter {
	a.AddHandlerFunc(in...)
	return a
}

// SetHandlers is a passthrough for the function stored in
// the Adapter.SetHandlersFunc property.
// Returns the current Adapter instance.
func (a *Adapter) SetHandlers(in ...midl.Middleware) midl.Adapter {
	a.SetHandlerFunc(in...)
	return a
}
