package stream

import (
	"log"
	"sync"
	"time"

	"github.com/itsubaki/gostream/pkg/event"
	"github.com/itsubaki/gostream/pkg/expr"
)

type Window interface {
	Where() []expr.Where
	Function() []expr.Function
	OrderBy() []expr.OrderBy
	Limit() []expr.LimitIF

	SetWhere(w ...expr.Where)
	SetFunction(f ...expr.Function)
	SetOrderBy(o ...expr.OrderBy)
	SetLimit(l ...expr.LimitIF)

	Input() chan interface{}
	Output() chan []event.Event
	Event() []event.Event

	Capacity() int
	Work()
	Listen(input interface{})
	Update(input interface{}) []event.Event
	Close()
}

type IdentityWindow struct {
	capacity int
	in       chan interface{}
	out      chan []event.Event
	event    []event.Event
	where    []expr.Where
	function []expr.Function
	orderBy  []expr.OrderBy
	limit    []expr.LimitIF
	closed   bool
	mutex    sync.RWMutex
}

func Capacity(capacity ...int) int {
	if len(capacity) > 0 {
		return capacity[0]
	}

	return 1024
}

func NewIdentity(capacity ...int) Window {
	cap := Capacity(capacity...)
	w := &IdentityWindow{
		capacity: cap,
		in:       make(chan interface{}, cap),
		out:      make(chan []event.Event, cap),
		event:    []event.Event{},
		where:    []expr.Where{},
		function: []expr.Function{},
		orderBy:  []expr.OrderBy{},
		limit:    []expr.LimitIF{},
		closed:   false,
		mutex:    sync.RWMutex{},
	}

	go w.Work()
	return w
}

func (w *IdentityWindow) Where() []expr.Where {
	return w.where
}

func (w *IdentityWindow) Function() []expr.Function {
	return w.function
}

func (w *IdentityWindow) OrderBy() []expr.OrderBy {
	return w.orderBy
}

func (w *IdentityWindow) Limit() []expr.LimitIF {
	return w.limit
}

func (w *IdentityWindow) SetWhere(wh ...expr.Where) {
	w.where = append(w.where, wh...)
}

func (w *IdentityWindow) SetFunction(f ...expr.Function) {
	w.function = append(w.function, f...)
}

func (w *IdentityWindow) SetOrderBy(o ...expr.OrderBy) {
	w.orderBy = append(w.orderBy, o...)
}

func (w *IdentityWindow) SetLimit(l ...expr.LimitIF) {
	w.limit = append(w.limit, l...)
}

func (w *IdentityWindow) Input() chan interface{} {
	return w.in
}

func (w *IdentityWindow) Output() chan []event.Event {
	return w.out
}

func (w *IdentityWindow) Event() []event.Event {
	return w.event
}

func (w *IdentityWindow) Capacity() int {
	return w.capacity
}

func (w *IdentityWindow) Work() {
	for input := range w.in {
		w.Listen(input)
	}
}

func (w *IdentityWindow) Listen(input interface{}) {
	if w.IsClosed() {
		return
	}

	events := w.Update(input)
	if len(events) == 0 {
		return
	}

	w.Output() <- events
}

func (w *IdentityWindow) Update(input interface{}) []event.Event {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("[WARNING] recover() %v %v", err, input)
		}
	}()

	e := event.New(input)

	// where
	for _, s := range w.where {
		if !s.Apply(e) {
			return event.List()
		}
	}

	// function
	w.event = append(w.event, e)
	for _, f := range w.function {
		w.event = f.Apply(w.event)
	}

	// order by
	events := append(event.List(), w.event...)
	for _, f := range w.orderBy {
		events = f.Apply(events)
	}

	// limit
	for _, f := range w.limit {
		events = f.Apply(events)
	}

	return events
}

func (w *IdentityWindow) Close() {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	if w.IsClosed() {
		return
	}

	w.closed = true

	close(w.Input())
	close(w.Output())
}

func (w *IdentityWindow) IsClosed() bool {
	return w.closed
}

func NewLength(_type interface{}, length int, capacity ...int) Window {
	w := NewIdentity(capacity...)

	w.SetWhere(
		expr.EqualsType{
			Accept: _type,
		},
	)

	w.SetFunction(
		&expr.Length{
			Length: length,
		},
	)

	return w
}

func NewLengthBatch(_type interface{}, length int, capacity ...int) Window {
	w := NewIdentity(capacity...)

	w.SetWhere(
		expr.EqualsType{
			Accept: _type,
		},
	)

	w.SetFunction(
		&expr.LengthBatch{
			Length: length,
			Batch:  event.List(),
		},
	)

	return w
}

func NewTime(_type interface{}, expire time.Duration, capacity ...int) Window {
	w := NewIdentity(capacity...)

	w.SetWhere(
		expr.EqualsType{
			Accept: _type,
		},
	)

	w.SetFunction(
		&expr.TimeDuration{
			Expire: expire,
		},
	)

	return w
}

func NewTimeBatch(_type interface{}, expire time.Duration, capacity ...int) Window {
	w := NewIdentity(capacity...)

	w.SetWhere(
		expr.EqualsType{
			Accept: _type,
		},
	)

	start := time.Now()
	end := start.Add(expire)
	w.SetFunction(
		&expr.TimeDurationBatch{
			Start:  start,
			End:    end,
			Expire: expire,
		},
	)

	return w
}
