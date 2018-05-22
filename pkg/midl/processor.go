package midl

type BodyProcessor interface {
	Process([]byte) error
}

type BodyProcessorFunc func([]byte) error

func (b BodyProcessorFunc) Process(in []byte) error {
	return b(in)
}
