package midlmock

// BodyProcessor is a configurable mock implementation of
// the midl.BodyProcessor interface.
type BodyProcessor struct {
	ProcessFunc func([]byte) error
}

// Process is a passthrough for the function stored at the
// BodyProcessor.ProcessFunc property.
func (b BodyProcessor) Process(in []byte) error {
	return b.ProcessFunc(in)
}
