package gocep

type Builder struct {
	window   string
	selector []string
	function []string
	view     []string
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) Window(window string) *Builder {
	b.window = window
	return b
}

func (b *Builder) Selector(selector string) *Builder {
	b.selector = append(b.selector, selector)
	return b
}

func (b *Builder) Function(function string) *Builder {
	b.function = append(b.function, function)
	return b
}

func (b *Builder) View(view string) *Builder {
	b.view = append(b.view, view)
	return b
}

func (b *Builder) Build() Window {
	return NewLengthWindow(16, 64)
}

type EventBuilder struct {
}

func NewEventBuilder() *EventBuilder {
	return &EventBuilder{}
}

func (b *EventBuilder) Build() interface{} {
	return nil
}
