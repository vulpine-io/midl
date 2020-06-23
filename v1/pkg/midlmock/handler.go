package midlmock

import (
	"github.com/vulpine-io/go-midl/v1/pkg/midl"
)

// EmptyHandler is a configurable mock implementation of the
// midl.EmptyHandler interface.
type EmptyHandler struct {
	HandleFunc func(midl.Request, midl.Response) []byte
}

// Handle is a passthrough for the function stored in the
// EmptyHandler.HandleFunc property.
func (e EmptyHandler) Handle(q midl.Request, s midl.Response) []byte {
	return e.HandleFunc(q, s)
}
