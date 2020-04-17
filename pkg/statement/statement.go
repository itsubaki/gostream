package statement

import (
	"time"

	"github.com/itsubaki/gostream/pkg/expr"
	"github.com/itsubaki/gostream/pkg/lexer"
	"github.com/itsubaki/gostream/pkg/stream"
)

type Statement struct {
	Window    lexer.Token
	EventType interface{}
	Length    int
	Time      time.Duration
	Where     []expr.Where
	Function  []expr.Function
	OrderBy   []expr.OrderBy
	Limit     []expr.LimitIF
}

func New() *Statement {
	return &Statement{
		Where:    []expr.Where{},
		Function: []expr.Function{},
		OrderBy:  []expr.OrderBy{},
		Limit:    []expr.LimitIF{},
	}
}

func (s *Statement) SetEventType(_type interface{}) {
	s.EventType = _type
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

func (s *Statement) SetWhere(w ...expr.Where) {
	s.Where = append(s.Where, w...)
}

func (s *Statement) SetFunction(f ...expr.Function) {
	s.Function = append(s.Function, f...)
}

func (s *Statement) SetOrderBy(o ...expr.OrderBy) {
	s.OrderBy = append(s.OrderBy, o...)
}

func (s *Statement) New(capacity ...int) (w stream.Window) {
	if s.Window == lexer.LENGTH {
		w = stream.NewLength(s.EventType, s.Length, capacity...)
	}

	if s.Window == lexer.LENGTH_BATCH {
		w = stream.NewLengthBatch(s.EventType, s.Length, capacity...)
	}

	if s.Window == lexer.TIME {
		w = stream.NewTime(s.EventType, s.Time, capacity...)
	}

	if s.Window == lexer.TIME_BATCH {
		w = stream.NewTimeBatch(s.EventType, s.Time, capacity...)
	}

	w.SetWhere(s.Where...)
	w.SetFunction(s.Function...)
	w.SetOrderBy(s.OrderBy...)
	w.SetLimit(s.Limit...)

	return w
}
