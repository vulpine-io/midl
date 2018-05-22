package midlmock

import "github.com/foxcapades/go-midl/pkg/midl"

type Serializer struct {
	SerializeFunc func(interface{}) ([]byte, error)
}

func (s Serializer) Serialize(in interface{}) ([]byte, error) {
	return s.SerializeFunc(in)
}

type ErrorSerializer struct {
	SerializeFunc func(error, midl.Request, midl.Response) []byte
}

func (m ErrorSerializer) Serialize(
	e error,
	q midl.Request,
	s midl.Response,
) []byte {
	return m.SerializeFunc(e, q, s)
}
