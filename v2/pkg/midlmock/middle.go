package midlmock

import (
	"github.com/Foxcapades/go-midl/v2/pkg/midl"
)

// Middleware is a configurable mock implementation of the
// midl.Middleware interface.
type Middleware struct {
	HandleFunc func(midl.Request) midl.Response
}

// Handle is a passthrough to the function stored at the
// Middleware.HandleFunc property.
func (m Middleware) Handle(q midl.Request) midl.Response {
	return m.HandleFunc(q)
}
