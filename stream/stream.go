package stream

import (
	"fmt"
	"log"
	"reflect"
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
	agr     []Aggeregate
	window  Window
	where   []Where
	orderby OrderByIF
	limit   LimitIF
	from    interface{}
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

	// aggregate function
	out := append(make([]Event, 0), s.events...)
	for _, a := range s.agr {
		out = a.Apply(out)
	}

	// order by limit offset
	out = s.limit.Apply(s.orderby.Apply(out))
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

func (s *Stream) From(t interface{}) *Stream {
	s.from = t
	s.where = append(s.where, From{Type: t})
	return s
}

func (s *Stream) Length(length int) *Stream {
	s.window = &Length{Length: length}
	return s
}

func (s *Stream) LengthBatch(length int) *Stream {
	s.window = &LengthBatch{Length: length, Batch: make([]Event, 0)}
	return s
}

func (s *Stream) Time(expire time.Duration, unit lexer.Token) *Stream {
	s.window = &Time{Expire: expire, Unit: unit}
	return s
}

func (s *Stream) TimeBatch(expire time.Duration, unit lexer.Token) *Stream {
	start := time.Now()
	end := start.Add(expire)
	s.window = &TimeBatch{
		Start:  start,
		End:    end,
		Expire: expire,
		Unit:   unit,
	}

	return s
}

func (s *Stream) SelectAll() *Stream {
	s.sel = append(s.sel, SelectAll{})
	return s
}

func (s *Stream) Select(name string) *Stream {
	s.sel = append(s.sel, Select{Name: name})
	return s
}

func (s *Stream) Average(name string) *Stream {
	s.agr = append(s.agr, Average{Name: name})
	return s
}

func (s *Stream) Sum(name string) *Stream {
	s.agr = append(s.agr, Sum{Name: name})
	return s
}

func (s *Stream) Count(name string) *Stream {
	s.agr = append(s.agr, Count{Name: name})
	return s
}

func (s *Stream) Max(name string) *Stream {
	s.agr = append(s.agr, Max{Name: name})
	return s
}

func (s *Stream) Min(name string) *Stream {
	s.agr = append(s.agr, Min{Name: name})
	return s
}

func (s *Stream) Distinct(name string) *Stream {
	s.agr = append(s.agr, Distinct{Name: name})
	return s
}

func (s *Stream) LargerThan(name string, value interface{}) *Stream {
	s.where = append(s.where, &LargerThan{
		Name:  name,
		Value: value,
	})

	return s
}

func (s *Stream) LessThan(name string, value interface{}) *Stream {
	s.where = append(s.where, &LessThan{
		Name:  name,
		Value: value,
	})

	return s
}

func (s *Stream) Equals(name string, value interface{}) *Stream {
	s.where = append(s.where, &Equals{
		Name:  name,
		Value: value,
	})

	return s
}

func (s *Stream) OrderBy(name string, desc bool) *Stream {
	if s.from == nil {
		panic(fmt.Errorf("from is nil"))
	}

	var index int
	v := reflect.ValueOf(s.from)
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Name == name {
			index = i
			break
		}
	}

	s.orderby = &OrderBy{
		Name:  name,
		Index: index,
		Desc:  desc,
	}

	return s
}

func (s *Stream) Limit(limit, offset int) *Stream {
	s.limit = &Limit{
		Limit:  limit,
		Offset: offset,
	}

	return s
}

func (s *Stream) String() string {
	var buf strings.Builder

	buf.WriteString("SELECT ")
	var sel strings.Builder
	for _, e := range s.sel {
		sel.WriteString(e.String())
		sel.WriteString(", ")
	}
	for _, e := range s.agr {
		sel.WriteString(e.String())
		sel.WriteString(", ")
	}
	buf.WriteString(strings.TrimRight(sel.String(), ", "))
	buf.WriteString(" ")
	buf.WriteString("FROM ")
	buf.WriteString(s.where[0].String())
	buf.WriteString(".")
	buf.WriteString(s.window.String())
	if len(s.where) > 1 {
		buf.WriteString(" ")
		buf.WriteString("WHERE ")
	}
	for i := 1; i < len(s.where); i++ {
		buf.WriteString(s.where[i].String())
	}
	buf.WriteString(" ")
	buf.WriteString(s.orderby.String())
	buf.WriteString(" ")
	buf.WriteString(s.limit.String())

	rep := strings.ReplaceAll(buf.String(), "  ", " ")
	return strings.TrimRight(rep, " ")
}
