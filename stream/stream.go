package stream

import (
	"log"
	"strings"
	"sync"
	"time"

	"github.com/itsubaki/gostream/lexer"
)

type Stream struct {
	in      chan interface{}
	out     chan []Event
	events  []Event
	sel     []SelectIF
	window  Window
	where   []Where
	orderby OrderByIF
	limit   LimitIF
	closed  bool
	mutex   sync.RWMutex
}

func New() *Stream {
	return &Stream{
		in:      make(chan interface{}, 1024),
		out:     make(chan []Event, 1024),
		events:  make([]Event, 0),
		sel:     make([]SelectIF, 0),
		where:   make([]Where, 0),
		orderby: &NoOrder{},
		limit:   &NoLimit{},
		mutex:   sync.RWMutex{},
	}
}

func (s *Stream) Input() chan interface{} {
	return s.in
}

func (s *Stream) Output() chan []Event {
	return s.out
}

func (s *Stream) Listen(input interface{}) {
	if s.IsClosed() {
		return
	}

	s.Update(input)

	// order by limit offset
	// no effect to s.events
	out := s.limit.Apply(s.orderby.Apply(s.events))

	if len(out) == 0 {
		return
	}

	s.Output() <- out
}

func (s *Stream) Update(input interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("[WARNING] recover() %v %v", err, input)
		}
	}()

	// where
	for _, w := range s.where {
		if w.Apply(input) {
			continue
		}
		return
	}

	// window
	buf := append(s.events, NewEvent(input))
	s.events = s.window.Apply(buf)

	// select
	for _, sl := range s.sel {
		s.events = sl.Apply(s.events)
	}
}

func (s *Stream) IsClosed() bool {
	return s.closed
}

func (s *Stream) Run() {
	for input := range s.in {
		s.Listen(input)
	}
}

func (s *Stream) Close() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.IsClosed() {
		return nil
	}
	s.closed = true

	close(s.Input())
	close(s.Output())

	return nil
}

func (s *Stream) Accept(t interface{}) {
	s.where = append(s.where, Accept{Type: t})
}

func (s *Stream) Length(length int) {
	s.window = &Length{Length: length}
}

func (s *Stream) LengthBatch(length int) {
	s.window = &LengthBatch{Length: length, Batch: make([]Event, 0)}
}

func (s *Stream) Time(expire time.Duration, unit lexer.Token) {
	s.window = &Time{Expire: expire, Unit: unit}
}

func (s *Stream) TimeBatch(expire time.Duration, unit lexer.Token) {
	start := time.Now()
	end := start.Add(expire)
	s.window = &TimeBatch{
		Start:  start,
		End:    end,
		Expire: expire,
		Unit:   unit,
	}
}

func (s *Stream) SelectAll() {
	s.sel = append(s.sel, SelectAll{})
}

func (s *Stream) Select(name string) {
	s.sel = append(s.sel, Select{Name: name})
}

func (s *Stream) Distinct(name string) {
	s.sel = append(s.sel, Distinct{Name: name})
}

func (s *Stream) OrderBy(name string, index int, desc bool) {
	s.orderby = &OrderBy{
		Name:  name,
		Index: index,
		Desc:  desc,
	}
}

func (s *Stream) Limit(limit, offset int) {
	s.limit = &Limit{
		Limit:  limit,
		Offset: offset,
	}
}

func (s *Stream) String() string {
	var buf strings.Builder

	buf.WriteString("SELECT ")
	var sel strings.Builder
	for _, e := range s.sel {
		sel.WriteString(e.String())
		sel.WriteString(", ")
	}
	buf.WriteString(strings.TrimRight(sel.String(), ", "))
	buf.WriteString(" ")
	buf.WriteString("FROM ")
	buf.WriteString(s.where[0].String())
	buf.WriteString(".")
	buf.WriteString(s.window.String())
	buf.WriteString(" ")
	buf.WriteString(s.orderby.String())
	buf.WriteString(" ")
	buf.WriteString(s.limit.String())

	rep := strings.ReplaceAll(buf.String(), "  ", " ")
	return strings.TrimRight(rep, " ")
}
