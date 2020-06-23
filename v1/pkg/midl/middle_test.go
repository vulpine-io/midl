package midl

import (
	"testing"

	c "github.com/smartystreets/goconvey/convey"
)

func TestMiddlewareFunc_Handle(t *testing.T) {
	c.Convey("", t, func() {
		var req Request
		var out Request

		req = &request{body: []byte("testing")}
		res := MiddlewareFunc(func(r Request) Response {
			out = r
			return &response{code: 204}
		}).Handle(req)

		c.So(req, c.ShouldResemble, out)
		c.So(res, c.ShouldResemble, &response{code: 204})
	})
}
