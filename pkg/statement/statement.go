package statement

import (
	"time"

	"github.com/itsubaki/gostream/pkg/clause"
	"github.com/itsubaki/gostream/pkg/lexer"
	"github.com/itsubaki/gostream/pkg/window"
)

type Statement struct {
	Window    lexer.Token
	EventType interface{}
	Length    int
	Time      time.Duration
	Where     []clause.Where
	Function  []clause.Function
	OrderBy   []clause.OrderBy
	Limit     []clause.LimitIF
}

func New() *Statement {
	return &Statement{
		Where:    []clause.Where{},
		Function: []clause.Function{},
		OrderBy:  []clause.OrderBy{},
		Limit:    []clause.LimitIF{},
	}
}

func (s *Statement) SetEventType(accept interface{}) {
	s.EventType = accept
}

func (s *Statement) SetWindow(token lexer.Token) {
	s.Window = token
}

func (s *Statement) SetLength(length int) {
	s.Length = length
}

func (s *Statement) SetTime(t time.Duration) {
	s.Time = t
}

func (s *Statement) SetWhere(w ...clause.Where) {
	s.Where = append(s.Where, w...)
}

func (s *Statement) SetFunction(f ...clause.Function) {
	s.Function = append(s.Function, f...)
}

func (s *Statement) SetOrderBy(o ...clause.OrderBy) {
	s.OrderBy = append(s.OrderBy, o...)
}

func (s *Statement) New(capacity ...int) (w window.Window) {
	if s.Window == lexer.LENGTH {
		w = window.NewLength(s.EventType, s.Length, capacity...)
	}

	if s.Window == lexer.LENGTH_BATCH {
		w = window.NewLengthBatch(s.EventType, s.Length, capacity...)
	}

	if s.Window == lexer.TIME {
		w = window.NewTime(s.EventType, s.Time, capacity...)
	}

	if s.Window == lexer.TIME_BATCH {
		w = window.NewTimeBatch(s.EventType, s.Time, capacity...)
	}

	w.SetWhere(s.Where...)
	w.SetFunction(s.Function...)
	w.SetOrderBy(s.OrderBy...)
	w.SetLimit(s.Limit...)

	return w
}
