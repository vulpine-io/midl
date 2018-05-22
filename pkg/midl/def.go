package midl

import "errors"

var (
	ErrorWrappedNil = errors.New("cannot wrap a nil request")
	ErrorNoHandlers = errors.New("no handlers")
)
