package midl

import (
	"net/http"
	"fmt"
	"strconv"
)

type Serializer interface {
	Serialize(interface{}) ([]byte, error)
}

type SerializerFunc func(interface{}) ([]byte, error)

func (s SerializerFunc) Serialize(in interface{}) ([]byte, error) {
	return s(in)
}

type ErrorSerializer interface {
	Serialize(error, Request, Response) []byte
}

type ErrorSerializerFunc func(error, Request, Response) []byte

func (f ErrorSerializerFunc) Serialize(e error, q Request, s Response) []byte {
	return f(e, q, s)
}

func DefaultJSONErrorSerializer() ErrorSerializer {
	return ErrorSerializerFunc(defJSONErrSerializer)
}

func defJSONErrSerializer(e error, _ Request, s Response) []byte {
	s.SetCode(http.StatusInternalServerError)
	return []byte(fmt.Sprintf(`{"error":%s}`, strconv.Quote(e.Error())))
}

func DefaultXMLErrorSerializer() ErrorSerializer {
	return ErrorSerializerFunc(defXMLErrSerializer)
}

func defXMLErrSerializer(e error, _ Request, s Response) []byte {
	s.SetCode(http.StatusInternalServerError)
	return []byte(fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<error>%s</error>`, e.Error()))
}
