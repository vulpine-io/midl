package midl

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDefaultAdapter_AddHandlers(t *testing.T) {
	convey.Convey("", t, func() {
		test := adapter{}
		mid1 := MiddlewareFunc(func(Request) Response { return nil })
		mid2 := MiddlewareFunc(func(Request) Response { return nil })
		mid3 := MiddlewareFunc(func(Request) Response { return nil })

		test.AddHandlers(mid1, mid2)

		convey.So(mid1, convey.ShouldEqual, test.handlers[0])
		convey.So(mid2, convey.ShouldEqual, test.handlers[1])
		convey.So(len(test.handlers), convey.ShouldEqual, 2)

		test.AddHandlers(mid3)

		convey.So(mid3, convey.ShouldEqual, test.handlers[2])
		convey.So(len(test.handlers), convey.ShouldEqual, 3)
	})
}

func TestDefaultAdapter_ContentType(t *testing.T) {
	convey.Convey("", t, func() {
		test := adapter{}

		test.ContentType("application/json")
		convey.So(test.contentType, convey.ShouldEqual, "application/json")
	})
}

func TestDefaultAdapter_EmptyHandler(t *testing.T) {
	convey.Convey("", t, func() {
		test := adapter{}
		empt := EmptyHandlerFunc(func(Request, Response) []byte { return nil })

		test.EmptyHandler(empt)
		convey.So(test.emptyHandler, convey.ShouldEqual, empt)
	})
}

func TestDefaultAdapter_ErrorSerializer(t *testing.T) {
	convey.Convey("", t, func() {
		test := adapter{}
		seri := ErrorSerializerFunc(func(error, Request, Response) []byte {
			return nil
		})

		test.ErrorSerializer(seri)
		convey.So(test.errSerializer, convey.ShouldEqual, seri)
	})
}

func TestDefaultAdapter_Serializer(t *testing.T) {
	convey.Convey("", t, func() {
		test := adapter{}
		seri := SerializerFunc(func(interface{}) ([]byte, error) { return nil, nil })

		test.Serializer(seri)
		convey.So(test.serializer, convey.ShouldEqual, seri)
	})
}

func TestDefaultAdapter_ServeHTTP(t *testing.T) {
}

func TestDefaultAdapter_SetHandlers(t *testing.T) {
	convey.Convey("", t, func() {
		test := adapter{}
		mid1 := MiddlewareFunc(func(Request) Response { return nil })
		mid2 := MiddlewareFunc(func(Request) Response { return nil })
		mid3 := MiddlewareFunc(func(Request) Response { return nil })

		test.SetHandlers(mid1, mid2)

		convey.So(mid1, convey.ShouldEqual, test.handlers[0])
		convey.So(mid2, convey.ShouldEqual, test.handlers[1])
		convey.So(len(test.handlers), convey.ShouldEqual, 2)

		test.SetHandlers(mid3)

		convey.So(mid3, convey.ShouldEqual, test.handlers[0])
		convey.So(len(test.handlers), convey.ShouldEqual, 1)
	})
}
