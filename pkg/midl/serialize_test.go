package midl

import (
	"errors"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"net/http"
)

func TestErrorSerializerFunc_Serialize(t *testing.T) {
	convey.Convey("", t, func() {
		var count int
		var err error
		var req Request
		var res Response

		run := ErrorSerializerFunc(func(in error, q Request, s Response) []byte {
			count++
			err = in
			req = q
			res = s
			return []byte("test value")
		})
		val := run.Serialize(
			errors.New("some error"),
			&request{body: []byte("body")},
			&response{body: map[string]string{"test": "value"}},
		)

		convey.So(count, convey.ShouldEqual, 1)
		convey.So(err, convey.ShouldResemble, errors.New("some error"))
		convey.So(val, convey.ShouldResemble, []byte("test value"))
		convey.So(req, convey.ShouldResemble, &request{body: []byte("body")})
		convey.So(res, convey.ShouldResemble, &response{body: map[string]string{"test": "value"}})
	})
}

func TestSerializerFunc_Serialize(t *testing.T) {
	convey.Convey("", t, func() {
		var count int
		var out interface{}

		run := SerializerFunc(func(in interface{}) ([]byte, error) {
			count++
			out = in
			return []byte("test value"), errors.New("test error")
		})
		val, err := run.Serialize(map[string]string{"test": "value"})

		convey.So(count, convey.ShouldEqual, 1)
		convey.So(out, convey.ShouldResemble, map[string]string{"test": "value"})
		convey.So(val, convey.ShouldResemble, []byte("test value"))
		convey.So(err, convey.ShouldResemble, errors.New("test error"))
	})
}

func TestJSONErrorSerializer(t *testing.T) {
	convey.Convey("", t, func() {
		res := response{code: http.StatusOK}
		data := DefaultJSONErrorSerializer().
			Serialize(errors.New("something"), nil, &res)
		convey.So(res.code, convey.ShouldEqual, http.StatusInternalServerError)
		convey.So(data, convey.ShouldResemble, []byte(`{"error":"something"}`))
	})
}

func TestXMLErrorSerializer(t *testing.T) {
	convey.Convey("", t, func() {
		res := response{code: http.StatusOK}
		data := DefaultXMLErrorSerializer().
			Serialize(errors.New("something"), nil, &res)

		convey.So(res.code, convey.ShouldEqual, http.StatusInternalServerError)
		convey.So(string(data), convey.ShouldResemble,
			"<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<error>something</error>")
	})
}
