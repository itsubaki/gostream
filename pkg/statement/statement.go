package statement

import (
	"time"

	"github.com/itsubaki/gostream-core/pkg/function"
	"github.com/itsubaki/gostream-core/pkg/lexer"
	"github.com/itsubaki/gostream-core/pkg/selector"
	"github.com/itsubaki/gostream-core/pkg/view"
	"github.com/itsubaki/gostream-core/pkg/window"
)

type Statement struct {
	Window    lexer.Token
	EventType interface{}
	Length    int
	Time      time.Duration
	Selector  []selector.Selector
	Function  []function.Function
	View      []view.View
}

func New() *Statement {
	return &Statement{
		Selector: []selector.Selector{},
		Function: []function.Function{},
		View:     []view.View{},
	}
}

func (st *Statement) SetEventType(_type interface{}) {
	st.EventType = _type
}

func (st *Statement) SetWindow(token lexer.Token) {
	st.Window = token
}

func (st *Statement) SetLength(length int) {
	st.Length = length
}

func (st *Statement) SetTime(t time.Duration) {
	st.Time = t
}

func (st *Statement) SetSelector(s ...selector.Selector) {
	st.Selector = append(st.Selector, s...)
}

func (st *Statement) SetFunction(f ...function.Function) {
	st.Function = append(st.Function, f...)
}

func (st *Statement) SetView(v ...view.View) {
	st.View = append(st.View, v...)
}

func (st *Statement) New(capacity ...int) (w window.Window) {
	if st.Window == lexer.LENGTH {
		w = window.NewLength(st.Length, capacity...)
	}

	if st.Window == lexer.LENGTH_BATCH {
		w = window.NewLengthBatch(st.Length, capacity...)
	}

	if st.Window == lexer.TIME {
		w = window.NewTime(st.Time, capacity...)
	}

	if st.Window == lexer.TIME_BATCH {
		w = window.NewTimeBatch(st.Time, capacity...)
	}

	w.SetSelector(st.Selector...)
	w.SetFunction(st.Function...)
	w.SetView(st.View...)

	return w
}
