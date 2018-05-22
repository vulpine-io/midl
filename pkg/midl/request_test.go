package midl

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	c "github.com/smartystreets/goconvey/convey"
)

type readCloser struct{ read func([]byte) (int, error) }

func (r readCloser) Read(p []byte) (int, error) { return r.read(p) }
func (r readCloser) Close() error               { return nil }

func TestRequest_Body(t *testing.T) {
	c.Convey("only read the raw body once", t, func() {
		count := 0
		reader := readCloser{func([]byte) (int, error) {
			count++
			return 0, io.EOF
		}}
		req := httptest.NewRequest(http.MethodGet, "http://foo.bar", reader)

		test := request{raw: req}
		test.Body()
		test.Body()

		c.So(count, c.ShouldEqual, 1)
	})

	c.Convey("returns the same value every time", t, func() {
		body := []byte("some test value")
		reader := ioutil.NopCloser(bytes.NewBuffer(body))
		req := httptest.NewRequest(http.MethodGet, "http://foo.bar", reader)

		test := request{raw: req}

		c.So(test.Body(), c.ShouldResemble, body)
		c.So(test.Body(), c.ShouldResemble, body)
		c.So(test.Body(), c.ShouldResemble, body)
	})

	c.Convey("does not attempt to read if an error is present", t, func() {
		count := 0
		reader := readCloser{func([]byte) (int, error) {
			count++
			return 0, io.EOF
		}}
		req := httptest.NewRequest(http.MethodGet, "http://foo.bar", reader)

		test := request{raw: req, error: errors.New("")}
		test.Body()

		c.So(count, c.ShouldEqual, 0)
	})

	c.Convey("stores any errors returned from the reader", t, func() {
		reader := readCloser{func([]byte) (int, error) {
			return 0, errors.New("some err")
		}}
		req := httptest.NewRequest(http.MethodGet, "http://foo.bar", reader)

		test := request{raw: req}
		test.Body()

		c.So(test.error, c.ShouldResemble, errors.New("some err"))
	})

	c.Convey("writes the body back to the raw request", t, func() {
		body := []byte("some test value")
		reader := ioutil.NopCloser(bytes.NewBuffer(body))
		req := httptest.NewRequest(http.MethodGet, "http://foo.bar", reader)

		test := request{raw: req}
		test.Body()
		val, _ := ioutil.ReadAll(req.Body)

		c.So(val, c.ShouldResemble, body)
	})
}

func TestRequest_Error(t *testing.T) {
	c.Convey("returns the stored error", t, func() {
		test := request{error: errors.New("some test error")}
		c.So(test.Error(), c.ShouldResemble, errors.New("some test error"))
	})
}

func TestRequest_Header(t *testing.T) {
	c.Convey("returns the first stored header value", t, func() {
		req := httptest.NewRequest(http.MethodGet, "http://foo.bar", nil)
		req.Header.Add("Foo", "Bar")
		req.Header.Add("Foo", "Fizz")
		req.Header.Add("Foo", "Buzz")

		test := request{raw: req}
		val, found := test.Header("Foo")
		c.So(val, c.ShouldEqual, "Bar")
		c.So(found, c.ShouldBeTrue)
	})

	c.Convey("returns the empty string if header not found", t, func() {
		req := httptest.NewRequest(http.MethodGet, "http://foo.bar", nil)
		req.Header.Add("Foo", "Bar")
		req.Header.Add("Foo", "Fizz")
		req.Header.Add("Foo", "Buzz")

		test := request{raw: req}
		val, found := test.Header("Bar")
		c.So(val, c.ShouldEqual, "")
		c.So(found, c.ShouldBeFalse)
	})
}

func TestRequest_Headers(t *testing.T) {
	c.Convey("returns the first stored header value", t, func() {
		req := httptest.NewRequest(http.MethodGet, "http://foo.bar", nil)
		req.Header.Add("Foo", "Bar")
		req.Header.Add("Foo", "Fizz")
		req.Header.Add("Foo", "Buzz")

		test := request{raw: req}
		val, ok := test.Headers("Foo")
		c.So(val, c.ShouldResemble, []string{"Bar", "Fizz", "Buzz"})
		c.So(ok, c.ShouldBeTrue)
	})

	c.Convey("returns the empty slice if header not found", t, func() {
		req := httptest.NewRequest(http.MethodGet, "http://foo.bar", nil)
		req.Header.Add("Foo", "Bar")
		req.Header.Add("Foo", "Fizz")
		req.Header.Add("Foo", "Buzz")

		test := request{raw: req}
		val, ok := test.Headers("Bar")
		c.So(val, c.ShouldResemble, []string(nil))
		c.So(ok, c.ShouldBeFalse)
	})
}

func TestRequest_Host(t *testing.T) {
	c.Convey("returns the stored host", t, func() {
		req := httptest.NewRequest(http.MethodGet, "http://foo.bar", nil)
		test := request{raw: req}
		c.So(test.Host(), c.ShouldEqual, "foo.bar")
	})
}

func TestRequest_Parameter(t *testing.T) {
	c.Convey("returns the parameter if present", t, func() {
		req := httptest.NewRequest(http.MethodGet, "http://foo.bar?test=foo", nil)
		test := request{raw: req}
		val, ok := test.Parameter("test")
		c.So(val, c.ShouldEqual, "foo")
		c.So(ok, c.ShouldBeTrue)
	})

	c.Convey("returns the empty string if not present", t, func() {
		req := httptest.NewRequest(http.MethodGet, "http://foo.bar?foo=test", nil)
		test := request{raw: req}
		val, ok := test.Parameter("test")
		c.So(val, c.ShouldEqual, "")
		c.So(ok, c.ShouldBeFalse)
	})
}

func TestRequest_Parameters(t *testing.T) {
	c.Convey("returns the parameters if present", t, func() {
		req := httptest.NewRequest(
			http.MethodGet,
			"http://foo.bar?test=foo&test=bar",
			nil,
		)
		test := request{raw: req}
		val, ok := test.Parameters("test")
		c.So(val, c.ShouldResemble, []string{"foo", "bar"})
		c.So(ok, c.ShouldBeTrue)
	})

	c.Convey("returns the empty set if not present", t, func() {
		req := httptest.NewRequest(http.MethodGet, "http://foo.bar?foo=test&foo=test2", nil)
		test := request{raw: req}
		val, ok := test.Parameters("test")
		c.So(val, c.ShouldResemble, []string(nil))
		c.So(ok, c.ShouldBeFalse)
	})
}

func TestRequest_ProcessBody(t *testing.T) {
	c.Convey("passes body if present into processor", t, func() {
		var input []byte
		body := []byte("test1")
		req := request{body: body, hasBody: true}

		req.ProcessBody(BodyProcessorFunc(func(b []byte) error {
			input = b
			return nil
		}))

		c.So(input, c.ShouldResemble, body)
	})

	c.Convey("does not call processor if error is present", t, func() {
		var input []byte
		body := []byte("test2")
		req := request{body: body, error: errors.New("test2 error"), hasBody: true}

		req.ProcessBody(BodyProcessorFunc(func(b []byte) error {
			input = b
			return nil
		}))
	})

	c.Convey("does not call processor if processor is nil", t, func() {
		body := []byte("test3")
		req := request{body: body, hasBody: true}
		c.So(func() { req.ProcessBody(nil) }, c.ShouldNotPanic)
	})

	c.Convey("stores error returned by processor", t, func() {
		body := []byte("test4")
		req := request{body: body, hasBody: true}
		req.ProcessBody(BodyProcessorFunc(func([]byte) error {
			return errors.New("test4 error")
		}))

		c.So(req.error, c.ShouldResemble, errors.New("test4 error"))
	})
}

func TestRequest_RawRequest(t *testing.T) {
	c.Convey("returns the stored request", t, func() {
		req := httptest.NewRequest(http.MethodGet, "http://foo.bar", nil)
		test := request{raw: req}
		c.So(test.RawRequest(), c.ShouldEqual, req)
	})
}

func TestNewRequest(t *testing.T) {
	c.Convey("returns error if request is nil", t, func() {
		req, err := NewRequest(nil)
		c.So(req, c.ShouldBeNil)
		c.So(err, c.ShouldResemble, ErrorWrappedNil)
	})

	c.Convey("returns wrapped request if request is not nil", t, func() {
		in := httptest.NewRequest(http.MethodGet, "http://foo.bar", nil)
		req, err := NewRequest(in)
		c.So(req.RawRequest(), c.ShouldEqual, in)
		c.So(err, c.ShouldBeNil)
	})
}
