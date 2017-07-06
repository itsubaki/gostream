package gocep

import (
	"log"
	"reflect"
	"time"
)

type Window interface {
	Input() chan Event
	Output() chan []Event
	Close()
}

type LengthWindow struct {
	capacity int
	in       chan Event
	out      chan []Event
	close    chan bool
	event    []Event
	selector []Selector
	function []Function
	view     []View
}

func NewLengthWindow(length, capacity int) *LengthWindow {
	w := &LengthWindow{
		capacity,
		make(chan Event, capacity),
		make(chan []Event, capacity),
		make(chan bool, 1),
		[]Event{},
		[]Selector{},
		[]Function{},
		[]View{},
	}
	w.Function(Length{length})

	go w.work()
	return w
}

func (w *LengthWindow) Selector(s Selector) {
	w.selector = append(w.selector, s)
}

func (w *LengthWindow) Function(f Function) {
	w.function = append(w.function, f)
}

func (w *LengthWindow) View(v View) {
	w.view = append(w.view, v)
}

func (w *LengthWindow) Input() chan Event {
	return w.in
}

func (w *LengthWindow) Output() chan []Event {
	return w.out
}

func (w *LengthWindow) Close() {
	w.close <- true
}

func (w *LengthWindow) work() {
	for {
		select {
		case close := <-w.close:
			if close {
				return
			}
		case e := <-w.in:
			w.Listen(e)
		}
	}
}

func (w *LengthWindow) Listen(e Event) {
	event := w.Update(e)
	if len(event) == 0 {
		return
	}
	w.Output() <- event
}

func (w *LengthWindow) Update(e Event) (event []Event) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err, reflect.TypeOf(e.Underlying))
		}
	}()

	for _, s := range w.selector {
		if !s.Select(e) {
			return event
		}
	}
	w.event = append(w.event, e.New())

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

type LengthBatchWindow struct {
	LengthWindow
}

func NewLengthBatchWindow(length, capacity int) *LengthBatchWindow {
	w := &LengthBatchWindow{
		LengthWindow{
			capacity,
			make(chan Event, capacity),
			make(chan []Event, capacity),
			make(chan bool, 1),
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
	LengthWindow
}

func NewTimeWindow(expire time.Duration, capacity int) *TimeWindow {
	w := &TimeWindow{
		LengthWindow{
			capacity,
			make(chan Event, capacity),
			make(chan []Event, capacity),
			make(chan bool, 1),
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
	LengthWindow
}

func NewTimeBatchWindow(expire time.Duration, capacity int) *TimeBatchWindow {
	w := &TimeBatchWindow{
		LengthWindow{
			capacity,
			make(chan Event, capacity),
			make(chan []Event, capacity),
			make(chan bool, 1),
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
