package gocep

import (
	"log"
	"sync"
	"time"
)

type Window interface {
	Selector() []Selector
	Function() []Function
	View() []View

	SetSelector(s Selector)
	SetFunction(f Function)
	SetView(v View)

	Input() chan interface{}
	Output() chan []Event
	Event() []Event
	Capacity() int

	Work()
	Listen(input interface{})
	Update(input interface{}) []Event

	Close()
}

type IdentityWindow struct {
	capacity int
	in       chan interface{}
	out      chan []Event
	event    []Event
	selector []Selector
	function []Function
	view     []View
	closed   bool
	mutex    sync.RWMutex
	Canceller
}

func Capacity(capacity ...int) int {
	if len(capacity) > 0 {
		return capacity[0]
	}
	return 1024
}

func NewIdentityWindow(capacity ...int) Window {
	cap := Capacity(capacity...)
	w := &IdentityWindow{
		cap,
		make(chan interface{}, cap),
		make(chan []Event, cap),
		[]Event{},
		[]Selector{},
		[]Function{},
		[]View{},
		false,
		sync.RWMutex{},
		NewCanceller(),
	}

	return w
}

func (w *IdentityWindow) Selector() []Selector {
	return w.selector
}

func (w *IdentityWindow) Function() []Function {
	return w.function
}

func (w *IdentityWindow) View() []View {
	return w.view
}

func (w *IdentityWindow) SetSelector(s Selector) {
	w.selector = append(w.selector, s)
}

func (w *IdentityWindow) SetFunction(f Function) {
	w.function = append(w.function, f)
}

func (w *IdentityWindow) SetView(v View) {
	w.view = append(w.view, v)
}

func (w *IdentityWindow) Input() chan interface{} {
	return w.in
}

func (w *IdentityWindow) Output() chan []Event {
	return w.out
}

func (w *IdentityWindow) Event() []Event {
	return w.event
}

func (w *IdentityWindow) Capacity() int {
	return w.capacity
}

func (w *IdentityWindow) Work() {
	for {
		select {
		case <-w.ctx.Done():
			return
		case input := <-w.in:
			// sequencial call
			w.Listen(input)
		}
	}
}

func (w *IdentityWindow) Listen(input interface{}) {
	if w.IsClosed() {
		return
	}

	event := w.Update(input)
	if len(event) == 0 {
		return
	}

	w.Output() <- event
}

func (w *IdentityWindow) Update(input interface{}) []Event {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("[WARNING] recover() %v %v", err, input)
		}
	}()

	e := NewEvent(input)
	for _, s := range w.selector {
		if !s.Select(e) {
			return []Event{}
		}
	}

	w.event = append(w.event, e)
	for _, f := range w.function {
		w.event = f.Apply(w.event)
	}

	event := append([]Event{}, w.event...)
	for _, f := range w.view {
		event = f.Apply(event)
	}

	return event
}

func (w *IdentityWindow) Close() {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	if w.IsClosed() {
		return
	}

	w.closed = true
	w.cancel()
	close(w.Input())
	close(w.Output())
}

func (w *IdentityWindow) IsClosed() bool {
	return w.closed
}

type LengthWindow struct {
	Window
}

func NewLengthWindow(length int, capacity ...int) Window {
	w := &LengthWindow{NewIdentityWindow(capacity...)}

	w.SetFunction(Length{length})
	go w.Work()

	return w
}

type LengthBatchWindow struct {
	Window
}

func NewLengthBatchWindow(length int, capacity ...int) Window {
	w := &LengthWindow{NewIdentityWindow(capacity...)}

	w.SetFunction(&LengthBatch{length, []Event{}})
	go w.Work()

	return w
}

type TimeWindow struct {
	Window
}

func NewTimeWindow(expire time.Duration, capacity ...int) Window {
	w := &TimeWindow{NewIdentityWindow(capacity...)}

	w.SetFunction(TimeDuration{expire})
	go w.Work()

	return w
}

type TimeBatchWindow struct {
	Window
}

func NewTimeBatchWindow(expire time.Duration, capacity ...int) Window {
	w := &TimeBatchWindow{NewIdentityWindow(capacity...)}

	start := time.Now()
	end := start.Add(expire)
	w.SetFunction(&TimeDurationBatch{start, end, expire})
	go w.Work()

	return w
}
