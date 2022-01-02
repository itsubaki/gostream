package stream

import (
	"log"
	"strings"
	"sync"
	"time"

	"github.com/itsubaki/gostream/lexer"
)

type Stream struct {
	in     chan interface{}
	out    chan []Event
	events []Event
	window Window
	where  []Where
	order  OrderByIF
	limit  LimitIF
	closed bool
	mutex  sync.RWMutex
}

func New() *Stream {
	return &Stream{
		in:     make(chan interface{}, 0),
		out:    make(chan []Event, 0),
		events: make([]Event, 0),
		where:  make([]Where, 0),
		order:  &NoOrder{},
		limit:  &NoLimit{},
		mutex:  sync.RWMutex{},
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

	events := s.Update(input)
	if len(events) == 0 {
		return
	}

	s.Output() <- events
}

func (s *Stream) Update(input interface{}) []Event {
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

		return make([]Event, 0)
	}

	// window
	buf := append(s.events, NewEvent(input))
	s.events = s.window.Apply(buf)

	// order by
	ordered := s.order.Apply(s.events)

	// limit
	return s.limit.Apply(ordered)
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

func (s *Stream) OrderBy(name string, desc bool) {
	s.order = &OrderBy{
		Name: name,
		Desc: desc,
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

	buf.WriteString("SELECT * FROM ")
	buf.WriteString(s.where[0].String())
	buf.WriteString(".")
	buf.WriteString(s.window.String())
	buf.WriteString(" ")
	buf.WriteString(s.order.String())
	buf.WriteString(" ")
	buf.WriteString(s.limit.String())

	return buf.String()
}
