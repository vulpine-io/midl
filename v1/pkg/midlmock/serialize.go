package midlmock

import (
	"github.com/vulpine-io/go-midl/v1/pkg/midl"
)

// Serializer is a configurable mock implementation of the
// midl.Serializer interface.
//
//   var ser midl.Serializer
//   ser = &midlmock.Serializer{
//       SerializerFunc: func(interface{}) ([]byte, error) {
//           return nil, errors.New("test error")
//       },
//   }
type Serializer struct {
	SerializeFunc func(interface{}) ([]byte, error)
}

// Serialize is a passthrough for the function stored at the
// Serializer.SerializeFunc property.
func (s Serializer) Serialize(in interface{}) ([]byte, error) {
	return s.SerializeFunc(in)
}

// ErrorSerializer is a configurable mock implementation of
// the midl.ErrorSerializer interface.
//
//   var ser midl.ErrorSerializer
//   ser = &midlmock.ErrorSerializer{
//       SerializerFunc: func(error, midl.Request, midl.Response) []byte {
//           return []byte("test body")
//       },
//   }
type ErrorSerializer struct {
	SerializeFunc func(error, midl.Request, midl.Response) []byte
}

// Serialize is a passthrough for the function stored at the
// ErrorSerializer.SerializeFunc property.
func (m ErrorSerializer) Serialize(
	e error,
	q midl.Request,
	s midl.Response,
) []byte {
	return m.SerializeFunc(e, q, s)
}
