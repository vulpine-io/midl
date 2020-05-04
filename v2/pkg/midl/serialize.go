package midl

import (
	"fmt"
	"net/http"
	"strconv"
)

// Serializer defines a service which can be used to
// serialize a response body into an array of bytes.
type Serializer interface {

	// Serialize is used to serialize the response body passed
	// as its input parameter into a byte array. If the
	// serialization fails the returned error will be passed
	// to the Adapter's ErrorSerializer.
	Serialize(interface{}) ([]byte, error)
}

// SerializerFunc defines a convenience wrapper for using a
// function as a Serializer implementation.
type SerializerFunc func(interface{}) ([]byte, error)

// Serialize is a simple passthrough to the wrapped
// function.
func (s SerializerFunc) Serialize(in interface{}) ([]byte, error) {
	return s(in)
}

// ErrorSerializer defines a service which can be used to
// serialize a given error into a byte array.
type ErrorSerializer interface {

	// Serialize is used to serialize the given error into an
	// array of bytes.  The HTTP request and response are
	// provided for context and to allow overriding/setting a
	// response status code or headers.
	Serialize(error, Request, Response) []byte
}

// ErrorSerializerFunc defines a convenience wrapper for
// using a function as an ErrorSerializer implementation.
type ErrorSerializerFunc func(error, Request, Response) []byte

// Serialize is a simple passthrough to the wrapped
// function.
func (f ErrorSerializerFunc) Serialize(e error, q Request, s Response) []byte {
	return f(e, q, s)
}

// DefaultJSONErrorSerializer returns an ErrorSerializer
// implementation which converts errors into a JSON byte
// array.
//
// The given error will be printed in a JSON string matching
// this pattern:
//   `{"error":"%s"}`
// The input error message is run through strconv.Quote for
// escaping special characters.
func DefaultJSONErrorSerializer() ErrorSerializer {
	return new(defJSONErrSerializer)
}

type defJSONErrSerializer struct{}

func (d defJSONErrSerializer) Serialize(e error, _ Request, s Response) []byte {
	s.SetCode(http.StatusInternalServerError)
	return []byte(fmt.Sprintf(`{"error":%s}`, strconv.Quote(e.Error())))
}

// DefaultXMLErrorSerializer returns an ErrorSerializer
// implementation which converts errors into an XML byte
// array.
//
// The given error will be printed in an XML string matching
// this pattern:
//   `<error>%s</error>`
func DefaultXMLErrorSerializer() ErrorSerializer {
	return new(defXMLErrSerializer)
}

type defXMLErrSerializer struct{}

func (d defXMLErrSerializer) Serialize(e error, _ Request, s Response) []byte {
	s.SetCode(http.StatusInternalServerError)
	return []byte(fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<error>%s</error>`, e.Error()))
}
