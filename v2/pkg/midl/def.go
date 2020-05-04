package midl

import "errors"

// Listing of errors that can be returned by the midl
// library specifically.
var (
	ErrWrappedNil = errors.New("cannot wrap a nil request")
	ErrNoHandlers = errors.New("no handlers")
)
