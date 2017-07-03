package gocep

type Builder struct {
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) Build() Window {
	return NewLengthWindow(16, 64)
}
