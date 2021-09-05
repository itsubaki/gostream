package window

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/itsubaki/gostream/pkg/event"
	"github.com/itsubaki/gostream/pkg/function"
)

type Window interface {
	SetWhere(w ...function.Where)
	SetFunction(f ...function.Function)
	SetOrderBy(o ...function.OrderBy)
	SetLimit(l ...function.LimitIF)

	Input() chan interface{}
	Output() chan []event.Event
	Event() []event.Event

	Capacity() int
	Work()
	Listen(input interface{})
	Update(input interface{}) []event.Event
	Close()

	// shortcut
	Function() *Function
	Select() *Select
	Sum() *Sum
	Average() *Average
	Count()
	Where() *Where
	OrderBy() *OrderBy
	Limit(limit int) *Limit
}

type Where struct {
	w *IdentityWindow
}

type LargerThan struct {
	w *Where
}

func (w *Where) LargerThan() *LargerThan {
	return &LargerThan{w}
}

func (l *LargerThan) Int(name string, value int) {
	l.w.w.SetWhere(function.LargerThanInt{Name: name, Value: value})
}

type Function struct {
	w *IdentityWindow
}

func (f *Function) Count() {
	f.w.SetFunction(function.Count{As: "count(*)"})
}

type Select struct {
	f *Function
}

func (f *Function) Select() *Select {
	return &Select{f}
}

func (s *Select) String(name string) {
	s.f.w.SetFunction(function.SelectString{Name: name, As: name})
}

func (s *Select) Int(name string) {
	s.f.w.SetFunction(function.SelectInt{Name: name, As: name})
}

type Average struct {
	f *Function
}

func (f *Function) Average() *Average {
	return &Average{f}
}

func (a *Average) Int(name string) {
	as := fmt.Sprintf("avg(%s)", name)
	a.f.w.SetFunction(function.AverageInt{Name: name, As: as})
}

type Sum struct {
	f *Function
}

func (f *Function) Sum() *Sum {
	return &Sum{f}
}

func (a *Sum) Int(name string) {
	as := fmt.Sprintf("sum(%s)", name)
	a.f.w.SetFunction(function.SumInt{Name: name, As: as})
}

type OrderBy struct {
	w    *IdentityWindow
	desc bool
}

func (o *OrderBy) Desc() *OrderBy {
	o.desc = true
	return o
}

func (o *OrderBy) Int(name string) {
	o.w.SetOrderBy(function.OrderByInt{Name: name, Desc: o.desc})
}

type Limit struct {
	w     *IdentityWindow
	limit int
}

func (l *Limit) Offset(offset int) {
	l.w.SetLimit(function.Limit{Limit: l.limit, Offset: offset})
}

type IdentityWindow struct {
	capacity int
	in       chan interface{}
	out      chan []event.Event
	event    []event.Event
	where    []function.Where
	function []function.Function
	orderBy  []function.OrderBy
	limit    []function.LimitIF
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
		where:    []function.Where{},
		function: []function.Function{},
		orderBy:  []function.OrderBy{},
		limit:    []function.LimitIF{},
		closed:   false,
		mutex:    sync.RWMutex{},
	}

	go w.Work()
	return w
}

func (w *IdentityWindow) SetWhere(wh ...function.Where) {
	w.where = append(w.where, wh...)
}

func (w *IdentityWindow) SetFunction(f ...function.Function) {
	w.function = append(w.function, f...)
}

func (w *IdentityWindow) SetOrderBy(o ...function.OrderBy) {
	w.orderBy = append(w.orderBy, o...)
}

func (w *IdentityWindow) SetLimit(l ...function.LimitIF) {
	if len(l) < 1 {
		return
	}

	if len(w.limit) < 1 {
		w.limit = append(w.limit, l[0])
	}

	w.limit[0] = l[0]
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

func (w *IdentityWindow) Select() *Select {
	return w.Function().Select()
}

func (w *IdentityWindow) Sum() *Sum {
	return w.Function().Sum()
}

func (w *IdentityWindow) Average() *Average {
	return w.Function().Average()
}

func (w *IdentityWindow) Count() {
	w.Function().Count()
}

func (w *IdentityWindow) Where() *Where {
	return &Where{w}
}

func (w *IdentityWindow) Function() *Function {
	return &Function{w}
}

func (w *IdentityWindow) OrderBy() *OrderBy {
	return &OrderBy{w, false}
}

func (w *IdentityWindow) Limit(limit int) *Limit {
	w.SetLimit(function.Limit{Limit: limit, Offset: 0})
	return &Limit{w, limit}
}

func (w *IdentityWindow) First() {
	w.SetLimit(function.First{})
}

func (w *IdentityWindow) Last() {
	w.SetLimit(function.Last{})
}

func NewLength(accept interface{}, length int, capacity ...int) Window {
	w := NewIdentity(capacity...)

	w.SetWhere(function.EqualsType{Accept: accept})
	w.SetFunction(&function.Length{Length: length})

	return w
}

func NewLengthBatch(accept interface{}, length int, capacity ...int) Window {
	w := NewIdentity(capacity...)

	w.SetWhere(function.EqualsType{Accept: accept})
	w.SetFunction(&function.LengthBatch{Length: length, Batch: event.List()})

	return w
}

func NewTime(accept interface{}, expire time.Duration, capacity ...int) Window {
	w := NewIdentity(capacity...)

	w.SetWhere(function.EqualsType{Accept: accept})
	w.SetFunction(&function.TimeDuration{Expire: expire})

	return w
}

func NewTimeBatch(accept interface{}, expire time.Duration, capacity ...int) Window {
	w := NewIdentity(capacity...)

	w.SetWhere(function.EqualsType{Accept: accept})

	start := time.Now()
	end := start.Add(expire)
	w.SetFunction(&function.TimeDurationBatch{
		Start:  start,
		End:    end,
		Expire: expire,
	},
	)

	return w
}
