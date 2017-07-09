package gocep

type EevntBuilder struct {
}

func NewEventBuilder() *EevntBuilder {
	return &EevntBuilder{}
}

func (b *EevntBuilder) Build() Window {
	return &SimpleWindow{}
}
