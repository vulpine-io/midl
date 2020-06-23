package midl

import (
	"errors"
	"net/http"
	"testing"

	c "github.com/smartystreets/goconvey/convey"
)

func TestErrorSerializerFunc_Serialize(t *testing.T) {
	c.Convey("", t, func() {
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

		c.So(count, c.ShouldEqual, 1)
		c.So(err, c.ShouldResemble, errors.New("some error"))
		c.So(val, c.ShouldResemble, []byte("test value"))
		c.So(req, c.ShouldResemble, &request{body: []byte("body")})
		c.So(res, c.ShouldResemble, &response{body: map[string]string{"test": "value"}})
	})
}

func TestSerializerFunc_Serialize(t *testing.T) {
	c.Convey("", t, func() {
		var count int
		var out interface{}

		run := SerializerFunc(func(in interface{}) ([]byte, error) {
			count++
			out = in
			return []byte("test value"), errors.New("test error")
		})
		val, err := run.Serialize(map[string]string{"test": "value"})

		c.So(count, c.ShouldEqual, 1)
		c.So(out, c.ShouldResemble, map[string]string{"test": "value"})
		c.So(val, c.ShouldResemble, []byte("test value"))
		c.So(err, c.ShouldResemble, errors.New("test error"))
	})
}

func TestJSONErrorSerializer(t *testing.T) {
	c.Convey("", t, func() {
		res := response{code: http.StatusOK}
		data := DefaultJSONErrorSerializer().
			Serialize(errors.New("something"), nil, &res)
		c.So(res.code, c.ShouldEqual, http.StatusInternalServerError)
		c.So(data, c.ShouldResemble, []byte(`{"error":"something"}`))
	})
}

func TestXMLErrorSerializer(t *testing.T) {
	c.Convey("", t, func() {
		res := response{code: http.StatusOK}
		data := DefaultXMLErrorSerializer().
			Serialize(errors.New("something"), nil, &res)

		c.So(res.code, c.ShouldEqual, http.StatusInternalServerError)
		c.So(string(data), c.ShouldResemble,
			"<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<error>something</error>")
	})
}
