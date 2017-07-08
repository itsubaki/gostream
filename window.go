package gocep

import (
	"context"
	"log"
	"reflect"
	"time"
)

type Window interface {
	Input() chan interface{}
	Output() chan []Event
	Close()
}

type SimpleWindow struct {
	capacity int
	in       chan interface{}
	out      chan []Event
	ctx      context.Context
	cancel   func()
	event    []Event
	selector []Selector
	function []Function
	view     []View
}

func NewSimpleWindow(capacity int) *SimpleWindow {
	ctx, cancel := context.WithCancel(context.Background())
	w := &SimpleWindow{
		capacity,
		make(chan interface{}, capacity),
		make(chan []Event, capacity),
		ctx,
		cancel,
		[]Event{},
		[]Selector{},
		[]Function{},
		[]View{},
	}

	go w.work()
	return w
}

func (w *SimpleWindow) Close() {
	w.cancel()
}

func (w *SimpleWindow) Selector(s Selector) {
	w.selector = append(w.selector, s)
}

func (w *SimpleWindow) Function(f Function) {
	w.function = append(w.function, f)
}

func (w *SimpleWindow) View(v View) {
	w.view = append(w.view, v)
}

func (w *SimpleWindow) Input() chan interface{} {
	return w.in
}

func (w *SimpleWindow) Output() chan []Event {
	return w.out
}

func (w *SimpleWindow) work() {
	for {
		select {
		case <-w.ctx.Done():
			return
		case input := <-w.in:
			w.Listen(input)
		}
	}
}

func (w *SimpleWindow) Listen(input interface{}) {
	event := w.Update(input)
	if len(event) == 0 {
		return
	}
	w.Output() <- event
}

func (w *SimpleWindow) Update(input interface{}) (event []Event) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err, reflect.TypeOf(input))
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
	SimpleWindow
}

func NewLengthWindow(length, capacity int) *LengthWindow {
	ctx, cancel := context.WithCancel(context.Background())
	w := &LengthWindow{
		SimpleWindow{
			capacity,
			make(chan interface{}, capacity),
			make(chan []Event, capacity),
			ctx,
			cancel,
			[]Event{},
			[]Selector{},
			[]Function{},
			[]View{},
		},
	}

	w.Function(Length{length})
	go w.work()
	return w
}

type LengthBatchWindow struct {
	SimpleWindow
}

func NewLengthBatchWindow(length, capacity int) *LengthBatchWindow {
	ctx, cancel := context.WithCancel(context.Background())
	w := &LengthBatchWindow{
		SimpleWindow{
			capacity,
			make(chan interface{}, capacity),
			make(chan []Event, capacity),
			ctx,
			cancel,
			[]Event{},
			[]Selector{},
			[]Function{},
			[]View{},
		},
	}

	w.Function(&LengthBatch{length, []Event{}})
	go w.work()
	return w
}

type TimeWindow struct {
	SimpleWindow
}

func NewTimeWindow(expire time.Duration, capacity int) *TimeWindow {
	ctx, cancel := context.WithCancel(context.Background())
	w := &TimeWindow{
		SimpleWindow{
			capacity,
			make(chan interface{}, capacity),
			make(chan []Event, capacity),
			ctx,
			cancel,
			[]Event{},
			[]Selector{},
			[]Function{},
			[]View{},
		},
	}
	w.Function(TimeDuration{expire})

	go w.work()
	return w
}

type TimeBatchWindow struct {
	SimpleWindow
}

func NewTimeBatchWindow(expire time.Duration, capacity int) *TimeBatchWindow {
	ctx, cancel := context.WithCancel(context.Background())
	w := &TimeBatchWindow{
		SimpleWindow{
			capacity,
			make(chan interface{}, capacity),
			make(chan []Event, capacity),
			ctx,
			cancel,
			[]Event{},
			[]Selector{},
			[]Function{},
			[]View{},
		},
	}

	start := time.Now()
	end := start.Add(expire)
	w.Function(&TimeDurationBatch{start, end, expire})
	go w.work()
	return w
}
