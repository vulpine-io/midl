package midlmock

import (
	"net/http"

	"github.com/foxcapades/go-midl/pkg/midl"
)

type Adapter struct {
	ServeHTTPFunc       func(http.ResponseWriter, *http.Request)
	EmptyHandlerFunc    func(midl.EmptyHandler)
	ContentTypeFunc     func(string)
	ErrorSerializerFunc func(midl.ErrorSerializer)
	SerializerFunc      func(midl.Serializer)
	AddHandlerFunc      func(...midl.Middleware)
	SetHandlerFunc      func(...midl.Middleware)
}

func (a *Adapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.ServeHTTPFunc(w, r)
}

func (a *Adapter) EmptyHandler(in midl.EmptyHandler) midl.Adapter {
	a.EmptyHandlerFunc(in)
	return a
}

func (a *Adapter) ContentType(in string) midl.Adapter {
	a.ContentTypeFunc(in)
	return a
}

func (a *Adapter) ErrorSerializer(in midl.ErrorSerializer) midl.Adapter {
	a.ErrorSerializerFunc(in)
	return a
}

func (a *Adapter) Serializer(in midl.Serializer) midl.Adapter {
	a.SerializerFunc(in)
	return a
}

func (a *Adapter) AddHandlers(in ...midl.Middleware) midl.Adapter {
	a.AddHandlerFunc(in...)
	return a
}

func (a *Adapter) SetHandlers(in ...midl.Middleware) midl.Adapter {
	a.SetHandlerFunc(in...)
	return a
}
