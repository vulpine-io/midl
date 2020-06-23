package midl

// BodyProcessor defines a service which can be run against
// a Request body.
type BodyProcessor interface {

	// Process is used to handle some work which requires a
	// request body.  This function is meant to be used with
	// Request.ProcessBody() which allows chaining various
	// operations, such as validation, deserialization, etc...
	Process([]byte) error
}

// BodyProcessorFunc is a convenience wrapper which allows
// the use of a function as a BodyProcessor implementation.
type BodyProcessorFunc func([]byte) error

// Process is a simple passthrough for the wrapped function.
func (b BodyProcessorFunc) Process(in []byte) error {
	return b(in)
}
