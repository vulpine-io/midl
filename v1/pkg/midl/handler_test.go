package midl

import (
	"testing"

	c "github.com/smartystreets/goconvey/convey"
)

func TestEmptyHandlerFunc_Handle(t *testing.T) {
	c.Convey("", t, func() {
		var out []byte
		var req Request
		var res Response

		out = EmptyHandlerFunc(func(a Request, b Response) []byte {
			req = a
			res = b
			return []byte("test output")
		}).Handle(&request{body: []byte("foo")}, &response{code: 200})

		c.So(out, c.ShouldResemble, []byte("test output"))
		c.So(res, c.ShouldResemble, &response{code: 200})
		c.So(req, c.ShouldResemble, &request{body: []byte("foo")})
	})
}

func TestDefaultEmptyHandler(t *testing.T) {
	c.Convey("", t, func() {
		c.So(DefaultEmptyHandler().Handle(nil, nil), c.ShouldBeNil)
	})
}
