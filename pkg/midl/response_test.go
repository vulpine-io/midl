package midl

import (
	"errors"
	"net/http"
	"testing"

	c "github.com/smartystreets/goconvey/convey"
)

func TestDefaultResponse_AddHeader(t *testing.T) {
	c.Convey("stores the passed header", t, func() {
		test := NewResponse()
		test.AddHeader("Foo", "Bar")
		c.So(
			test.RawHeaders(),
			c.ShouldResemble,
			http.Header{"Foo": {"Bar"}},
		)
	})
}

func TestDefaultResponse_AddHeaders(t *testing.T) {
	c.Convey("stores the passed header", t, func() {
		test := NewResponse()
		test.AddHeaders("Foo", []string{"Bar", "Fizz"})
		c.So(
			test.RawHeaders(),
			c.ShouldResemble,
			http.Header{"Foo": {"Bar", "Fizz"}},
		)
	})
}

func TestDefaultResponse_Error(t *testing.T) {
	c.Convey("", t, func() {})
}

func TestDefaultResponse_Header(t *testing.T) {
	c.Convey("returns stored headers", t, func() {
		test := NewResponse()
		test.RawHeaders().Add("Foo", "Bar")
		test.RawHeaders().Add("Foo", "Fizz")
		test.RawHeaders().Add("Foo", "Buzz")

		c.So(test.Header("Foo"), c.ShouldResemble, "Bar")
	})
}

func TestDefaultResponse_Headers(t *testing.T) {
	c.Convey("returns stored headers", t, func() {
		test := NewResponse()
		test.RawHeaders().Add("Foo", "Bar")
		test.RawHeaders().Add("Foo", "Fizz")
		test.RawHeaders().Add("Foo", "Buzz")

		c.So(
			test.Headers("Foo"),
			c.ShouldResemble,
			[]string{"Bar", "Fizz", "Buzz"},
		)
	})
}

func TestDefaultResponse_SetBody(t *testing.T) {
	c.Convey("sets the passed in body on the response", t, func() {
		value := map[string][]int{}
		test := NewResponse()
		test.SetBody(value)
		c.So(test.Body(), c.ShouldResemble, value)
	})
}

func TestDefaultResponse_SetCode(t *testing.T) {
	c.Convey("stores the passed in code on the response", t, func() {
		value := 403
		test := NewResponse()
		test.SetCode(value)
		c.So(test.Code(), c.ShouldResemble, value)
	})
}

func TestDefaultResponse_SetError(t *testing.T) {
	c.Convey("stores error", t, func() {
		test := NewResponse()
		test.SetError(errors.New("testing"))
		c.So(
			test.Error(),
			c.ShouldResemble,
			errors.New("testing"),
		)
	})
}

func TestDefaultResponse_SetHeader(t *testing.T) {
	c.Convey("stores the passed header", t, func() {
		test := NewResponse()
		test.SetHeader("Foo", "Bar")
		c.So(
			test.RawHeaders(),
			c.ShouldResemble,
			http.Header{"Foo": {"Bar"}},
		)
	})
}

func TestDefaultResponse_SetHeaders(t *testing.T) {
	c.Convey("stores the passed header", t, func() {
		test := NewResponse()
		test.SetHeaders("Foo", []string{"Bar", "Fizz"})
		c.So(
			test.RawHeaders(),
			c.ShouldResemble,
			http.Header{"Foo": {"Bar", "Fizz"}},
		)
	})
}

func TestMakeErrorResponse(t *testing.T) {
	c.Convey("", t, func() {
		res := MakeErrorResponse(http.StatusBadRequest, errors.New("some error"))

		c.So(res.Code(), c.ShouldEqual, http.StatusBadRequest)
		c.So(res.Error(), c.ShouldResemble, errors.New("some error"))
		c.So(res.Body(), c.ShouldResemble, nil)
	})
}

func TestMakeResponse(t *testing.T) {
	c.Convey("", t, func() {
		res := MakeResponse(http.StatusOK, []string{"ok", "no"})

		c.So(res.Code(), c.ShouldEqual, http.StatusOK)
		c.So(res.Error(), c.ShouldResemble, nil)
		c.So(res.Body(), c.ShouldResemble, []string{"ok", "no"})
	})
}
