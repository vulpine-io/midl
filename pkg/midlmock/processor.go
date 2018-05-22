package midlmock

type BodyProcessor struct {
	ProcessFunc func([]byte) error
}

func (b BodyProcessor) Process(in []byte) error {
	return b.ProcessFunc(in)
}

