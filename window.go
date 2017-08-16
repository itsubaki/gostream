package gocep

import (
	"log"
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
	Canceller
}

func NewIdentityWindow(capacity int) Window {
	w := &IdentityWindow{
		capacity,
		make(chan interface{}, capacity),
		make(chan []Event, capacity),
		[]Event{},
		[]Selector{},
		[]Function{},
		[]View{},
		NewCanceller(),
	}

	go w.work()
	return w
}

func (w *IdentityWindow) Close() {
	w.cancel()
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
	event := w.Update(input)
	if len(event) == 0 {
		return
	}
	w.Output() <- event
}

func (w *IdentityWindow) Update(input interface{}) (event []Event) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err, input)
		}
	}()

	e := NewEvent(input)
	for _, s := range w.selector {
		if !s.Select(e) {
			return event
		}
	}
	w.event = append(w.event, e)

	for _, f := range w.function {
		w.event = f.Apply(w.event)
	}

	for _, e := range w.event {
		event = append(event, e)
	}

	for _, f := range w.view {
		event = f.Apply(event)
	}

	return event
}

type LengthWindow struct {
	IdentityWindow
}

func NewLengthWindow(length, capacity int) Window {
	w := &LengthWindow{
		IdentityWindow{
			capacity,
			make(chan interface{}, capacity),
			make(chan []Event, capacity),
			[]Event{},
			[]Selector{},
			[]Function{},
			[]View{},
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

func NewLengthBatchWindow(length, capacity int) Window {
	w := &LengthBatchWindow{
		IdentityWindow{
			capacity,
			make(chan interface{}, capacity),
			make(chan []Event, capacity),
			[]Event{},
			[]Selector{},
			[]Function{},
			[]View{},
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

func NewTimeWindow(expire time.Duration, capacity int) Window {
	w := &TimeWindow{
		IdentityWindow{
			capacity,
			make(chan interface{}, capacity),
			make(chan []Event, capacity),
			[]Event{},
			[]Selector{},
			[]Function{},
			[]View{},
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

func NewTimeBatchWindow(expire time.Duration, capacity int) Window {
	w := &TimeBatchWindow{
		IdentityWindow{
			capacity,
			make(chan interface{}, capacity),
			make(chan []Event, capacity),
			[]Event{},
			[]Selector{},
			[]Function{},
			[]View{},
			NewCanceller(),
		},
	}

	start := time.Now()
	end := start.Add(expire)
	w.SetFunction(&TimeDurationBatch{start, end, expire})
	go w.work()
	return w
}
