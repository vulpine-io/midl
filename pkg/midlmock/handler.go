package midlmock

import (
	"github.com/foxcapades/go-midl/pkg/midl"
)

type EmptyHandler struct {
	HandleFunc func(midl.Request, midl.Response) []byte
}

func (e EmptyHandler) Handle(q midl.Request, s midl.Response) []byte {
	return e.HandleFunc(q, s)
}
