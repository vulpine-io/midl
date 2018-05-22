package midlmock

import "github.com/foxcapades/go-midl/pkg/midl"

type Middleware struct {
	HandleFunc func(midl.Request) midl.Response
}

func (m Middleware) Handle(q midl.Request) midl.Response {
	return m.HandleFunc(q)
}
