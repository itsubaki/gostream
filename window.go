package gocep

import (
	"log"
	"sync"
	"time"
)

type Window interface {
	SetSelector(s Selector)
	SetFunction(f Function)
	SetView(v View)
	Input() chan interface{}
	Output() chan []Event
	Event() []Event
	Close()
	Listen(input interface{})
	Update(input interface{}) []Event
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

	go w.work()
	return w
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

func (w *IdentityWindow) work() {
	for {
		select {
		case <-w.ctx.Done():
			return
		case input := <-w.in:
			w.Listen(input)
		}
	}
}

func (w *IdentityWindow) Listen(input interface{}) {
	w.mutex.RLock()
	defer w.mutex.RUnlock()

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

type LengthWindow struct {
	IdentityWindow
}

func NewLengthWindow(length int, capacity ...int) Window {
	cap := Capacity(capacity...)
	w := &LengthWindow{
		IdentityWindow{
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
		},
	}

	w.SetFunction(Length{length})
	go w.work()
	return w
}

type LengthBatchWindow struct {
	IdentityWindow
}

func NewLengthBatchWindow(length int, capacity ...int) Window {
	cap := Capacity(capacity...)
	w := &LengthBatchWindow{
		IdentityWindow{
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
		},
	}

	w.SetFunction(&LengthBatch{length, []Event{}})
	go w.work()
	return w
}

type TimeWindow struct {
	IdentityWindow
}

func NewTimeWindow(expire time.Duration, capacity ...int) Window {
	cap := Capacity(capacity...)
	w := &TimeWindow{
		IdentityWindow{
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
		},
	}
	w.SetFunction(TimeDuration{expire})

	go w.work()
	return w
}

type TimeBatchWindow struct {
	IdentityWindow
}

func NewTimeBatchWindow(expire time.Duration, capacity ...int) Window {
	cap := Capacity(capacity...)
	w := &TimeBatchWindow{
		IdentityWindow{
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
		},
	}

	start := time.Now()
	end := start.Add(expire)
	w.SetFunction(&TimeDurationBatch{start, end, expire})
	go w.work()
	return w
}
