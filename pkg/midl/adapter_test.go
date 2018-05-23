package midl

import (
	"errors"
	"net/http/httptest"
	"testing"

	c "github.com/smartystreets/goconvey/convey"
)

func TestDefaultAdapter_AddHandlers(t *testing.T) {
	c.Convey("", t, func() {
		test := adapter{}
		mid1 := MiddlewareFunc(func(Request) Response { return nil })
		mid2 := MiddlewareFunc(func(Request) Response { return nil })
		mid3 := MiddlewareFunc(func(Request) Response { return nil })

		test.AddHandlers(mid1, mid2)

		c.So(mid1, c.ShouldEqual, test.handlers[0])
		c.So(mid2, c.ShouldEqual, test.handlers[1])
		c.So(len(test.handlers), c.ShouldEqual, 2)

		test.AddHandlers(mid3)

		c.So(mid3, c.ShouldEqual, test.handlers[2])
		c.So(len(test.handlers), c.ShouldEqual, 3)
	})
}

func TestDefaultAdapter_ContentType(t *testing.T) {
	c.Convey("", t, func() {
		test := adapter{}

		test.ContentType("application/json")
		c.So(test.contentType, c.ShouldEqual, "application/json")
	})
}

func TestDefaultAdapter_EmptyHandler(t *testing.T) {
	c.Convey("", t, func() {
		test := adapter{}
		empt := EmptyHandlerFunc(func(Request, Response) []byte { return nil })

		test.EmptyHandler(empt)
		c.So(test.emptyHandler, c.ShouldEqual, empt)
	})
}

func TestDefaultAdapter_ErrorSerializer(t *testing.T) {
	c.Convey("", t, func() {
		test := adapter{}
		seri := ErrorSerializerFunc(func(error, Request, Response) []byte {
			return nil
		})

		test.ErrorSerializer(seri)
		c.So(test.errSerializer, c.ShouldEqual, seri)
	})
}

func TestDefaultAdapter_Serializer(t *testing.T) {
	c.Convey("", t, func() {
		test := adapter{}
		seri := SerializerFunc(func(interface{}) ([]byte, error) {
			return nil, nil
		})

		test.Serializer(seri)
		c.So(test.serializer, c.ShouldEqual, seri)
	})
}

func TestDefaultAdapter_ServeHTTP(t *testing.T) {
	c.Convey("stops at the first non nil response", t, func() {
		mid1Count := 0
		mid2Count := 0
		mid3Count := 0

		mid1 := MiddlewareFunc(func(r Request) Response {
			mid1Count++
			return nil
		})
		mid2 := MiddlewareFunc(func(r Request) Response {
			mid2Count++
			return &response{code: 404, body: "test2"}
		})
		mid3 := MiddlewareFunc(func(r Request) Response {
			mid3Count++
			return &response{code: 500, body: "test3"}
		})

		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "http://foo.bar", nil)

		adapter{
			handlers: []Middleware{mid1, mid2, mid3},
			serializer: SerializerFunc(func(interface{}) ([]byte, error) {
				return nil, nil
			}),
		}.ServeHTTP(w, r)

		c.So(w.Code, c.ShouldEqual, 404)
		c.So(mid1Count, c.ShouldEqual, 1)
		c.So(mid2Count, c.ShouldEqual, 1)
		c.So(mid3Count, c.ShouldEqual, 0)
	})

	c.Convey("sends empty responses to empty handler", t, func() {
		midCount := 0

		mid := MiddlewareFunc(func(r Request) Response {
			midCount++
			return &response{code: 300}
		})

		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "http://foo.bar", nil)

		adapter{
			handlers: []Middleware{mid},
			emptyHandler: EmptyHandlerFunc(func(_ Request, s Response) []byte {
				s.SetCode(302)
				return []byte("test")
			}),
		}.ServeHTTP(w, r)

		c.So(w.Code, c.ShouldEqual, 302)
		c.So(w.Body.Bytes(), c.ShouldResemble, []byte("test"))
		c.So(midCount, c.ShouldEqual, 1)
	})

	c.Convey("skips nil handlers", t, func() {
		mid := MiddlewareFunc(func(r Request) Response {
			return &response{code: 300}
		})

		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "http://foo.bar", nil)

		adapter{
			handlers: []Middleware{nil, nil, mid},
			emptyHandler: EmptyHandlerFunc(func(_ Request, s Response) []byte {
				return nil
			}),
		}.ServeHTTP(w, r)

		c.So(w.Code, c.ShouldEqual, 300)
	})

	c.Convey("writes nil request error if input request is nil", t, func() {
		w := httptest.NewRecorder()
		adapter{
			handlers: []Middleware{},
			errSerializer: ErrorSerializerFunc(
				func(err error, _ Request, s Response) []byte {
					return []byte(err.Error())
				},
			),
		}.ServeHTTP(w, nil)

		c.So(w.Body.Bytes(), c.ShouldResemble, []byte(ErrorWrappedNil.Error()))
	})

	c.Convey("writes no handler error if no handlers are present", t, func() {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "http://foo.bar", nil)

		adapter{
			handlers: []Middleware{},
			errSerializer: ErrorSerializerFunc(
				func(err error, _ Request, s Response) []byte {
					return []byte(err.Error())
				},
			),
		}.ServeHTTP(w, r)

		c.So(w.Body.Bytes(), c.ShouldResemble, []byte(ErrorNoHandlers.Error()))
	})

	c.Convey("writes serializer error if present", t, func() {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "http://foo.bar", nil)

		adapter{
			handlers: []Middleware{MiddlewareFunc(func(r Request) Response {
				return &response{code: 300, body: "something"}
			})},
			serializer: SerializerFunc(func(interface{}) ([]byte, error) {
				return nil, errors.New("serializer error")
			}),
			errSerializer: ErrorSerializerFunc(
				func(err error, _ Request, s Response) []byte {
					return []byte(err.Error())
				},
			),
		}.ServeHTTP(w, r)

		c.So(w.Body.Bytes(), c.ShouldResemble, []byte("serializer error"))
	})

	c.Convey("writes response headers", t, func() {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "http://foo.bar", nil)

		adapter{
			handlers: []Middleware{MiddlewareFunc(func(r Request) Response {
				return &response{
					code: 300,
					body: "something",
					head: Header{"Something": {"foo", "bar"}},
				}
			})},
			serializer: SerializerFunc(func(interface{}) ([]byte, error) {
				return []byte("something"), nil
			}),
		}.ServeHTTP(w, r)

		c.So(w.HeaderMap, c.ShouldResemble, Header{"Something": {"foo", "bar"}})
	})

	c.Convey("appends content type header when present", t, func() {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "http://foo.bar", nil)

		adapter{
			contentType: "text/plain",
			handlers: []Middleware{MiddlewareFunc(func(r Request) Response {
				return &response{
					code: 300,
					body: "something",
				}
			})},
			serializer: SerializerFunc(func(interface{}) ([]byte, error) {
				return []byte("something"), nil
			}),
		}.ServeHTTP(w, r)

		c.So(w.HeaderMap, c.ShouldResemble, Header{"Content-Type": {"text/plain"}})
	})

	c.Convey("writes response errors", t, func() {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "http://foo.bar", nil)

		adapter{
			handlers: []Middleware{MiddlewareFunc(func(r Request) Response {
				return &response{
					code:  300,
					error: errors.New("test handler error"),
				}
			})},
			errSerializer: ErrorSerializerFunc(
				func(err error, _ Request, _ Response) []byte {
					return []byte(err.Error())
				},
			),
		}.ServeHTTP(w, r)

		c.So(w.Body.Bytes(), c.ShouldResemble, []byte("test handler error"))
	})
}

func TestDefaultAdapter_SetHandlers(t *testing.T) {
	c.Convey("", t, func() {
		test := adapter{}
		mid1 := MiddlewareFunc(func(Request) Response { return nil })
		mid2 := MiddlewareFunc(func(Request) Response { return nil })
		mid3 := MiddlewareFunc(func(Request) Response { return nil })

		test.SetHandlers(mid1, mid2)

		c.So(mid1, c.ShouldEqual, test.handlers[0])
		c.So(mid2, c.ShouldEqual, test.handlers[1])
		c.So(len(test.handlers), c.ShouldEqual, 2)

		test.SetHandlers(mid3)

		c.So(mid3, c.ShouldEqual, test.handlers[0])
		c.So(len(test.handlers), c.ShouldEqual, 1)
	})
}
