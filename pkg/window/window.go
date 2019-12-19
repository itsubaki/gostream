package window

import (
	"log"
	"sync"
	"time"

	"github.com/itsubaki/gostream/pkg/event"
	"github.com/itsubaki/gostream/pkg/function"
	"github.com/itsubaki/gostream/pkg/selector"
	"github.com/itsubaki/gostream/pkg/view"
)

type Window interface {
	Selector() []selector.Selector
	Function() []function.Function
	View() []view.View

	SetSelector(s ...selector.Selector)
	SetFunction(f ...function.Function)
	SetView(v ...view.View)

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
	selector []selector.Selector
	function []function.Function
	view     []view.View
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
		cap,
		make(chan interface{}, cap),
		make(chan []event.Event, cap),
		[]event.Event{},
		[]selector.Selector{},
		[]function.Function{},
		[]view.View{},
		false,
		sync.RWMutex{},
	}

	go w.Work()
	return w
}

func (w *IdentityWindow) Selector() []selector.Selector {
	return w.selector
}

func (w *IdentityWindow) Function() []function.Function {
	return w.function
}

func (w *IdentityWindow) View() []view.View {
	return w.view
}

func (w *IdentityWindow) SetSelector(s ...selector.Selector) {
	w.selector = append(w.selector, s...)
}

func (w *IdentityWindow) SetFunction(f ...function.Function) {
	w.function = append(w.function, f...)
}

func (w *IdentityWindow) SetView(v ...view.View) {
	w.view = append(w.view, v...)
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
	for input := range w.in{
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
	for _, s := range w.selector {
		if !s.Select(e) {
			return event.List()
		}
	}

	w.event = append(w.event, e)
	for _, f := range w.function {
		w.event = f.Apply(w.event)
	}

	events := append(event.List(), w.event...)
	for _, f := range w.view {
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

	w.SetSelector(
		selector.EqualsType{
			Accept: _type,
		},
	)

	w.SetFunction(
		&function.Length{
			Length: length,
		},
	)

	return w
}

func NewLengthBatch(_type interface{}, length int, capacity ...int) Window {
	w := NewIdentity(capacity...)

	w.SetSelector(
		selector.EqualsType{
			Accept: _type,
		},
	)

	w.SetFunction(
		&function.LengthBatch{
			Length: length,
			Batch:  event.List(),
		},
	)

	return w
}

func NewTime(_type interface{}, expire time.Duration, capacity ...int) Window {
	w := NewIdentity(capacity...)

	w.SetSelector(
		selector.EqualsType{
			Accept: _type,
		},
	)

	w.SetFunction(
		&function.TimeDuration{
			Expire: expire,
		},
	)

	return w
}

func NewTimeBatch(_type interface{}, expire time.Duration, capacity ...int) Window {
	w := NewIdentity(capacity...)

	w.SetSelector(
		selector.EqualsType{
			Accept: _type,
		},
	)

	start := time.Now()
	end := start.Add(expire)
	w.SetFunction(
		&function.TimeDurationBatch{
			Start:  start,
			End:    end,
			Expire: expire,
		},
	)

	return w
}
