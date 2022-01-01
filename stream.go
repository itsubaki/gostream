package gostream

import (
	"log"
	"sync"
	"time"
)

type Stream interface {
	Input() chan interface{}
	Output() chan []Event
	Listen(input interface{})
	Update(input interface{}) []Event
	Run()
	Close() error
}

type IdentityStream struct {
	in     chan interface{}
	out    chan []Event
	fn     []Function
	where  []Where
	events []Event
	closed bool
	mutex  sync.RWMutex
}

func (s *IdentityStream) Input() chan interface{} {
	return s.in
}

func (s *IdentityStream) Output() chan []Event {
	return s.out
}

func (s *IdentityStream) Listen(input interface{}) {
	if s.IsClosed() {
		return
	}

	events := s.Update(input)
	if len(events) == 0 {
		return
	}

	s.Output() <- events
}

func (s *IdentityStream) Update(input interface{}) []Event {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("[WARNING] recover() %v %v", err, input)
		}
	}()

	for _, w := range s.where {
		if w.Apply(input) {
			continue
		}

		return make([]Event, 0)
	}

	buf := append(s.events, NewEvent(input))
	for _, f := range s.fn {
		buf = f.Apply(buf)
	}
	s.events = buf

	return s.events
}

func (s *IdentityStream) IsClosed() bool {
	return s.closed
}

func (s *IdentityStream) Run() {
	for input := range s.in {
		s.Listen(input)
	}
}

func (s *IdentityStream) Close() error {
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

func NewLength(accept interface{}, length int) Stream {
	s := &IdentityStream{}
	s.where = append(s.where, EqualsType{Accept: accept})
	s.fn = append(s.fn, &Length{Length: length})
	return s
}

func NewLengthBatch(accept interface{}, length int) Stream {
	s := &IdentityStream{}
	s.where = append(s.where, EqualsType{Accept: accept})
	s.fn = append(s.fn, &LengthBatch{Length: length, Batch: make([]Event, 0)})
	return s
}

func NewTime(accept interface{}, expire time.Duration, capacity ...int) Stream {
	s := &IdentityStream{}
	s.where = append(s.where, EqualsType{Accept: accept})
	s.fn = append(s.fn, &Time{Expire: expire})
	return s
}

func NewTimeBatch(accept interface{}, expire time.Duration) Stream {
	s := &IdentityStream{}
	s.where = append(s.where, EqualsType{Accept: accept})

	start := time.Now()
	end := start.Add(expire)
	s.fn = append(s.fn, &TimeBatch{
		Start:  start,
		End:    end,
		Expire: expire,
	})

	return s
}
