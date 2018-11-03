package gocep

import (
	"sync"
)

type Stream struct {
	capacity int
	in       chan interface{}
	out      chan []Event
	window   []Window
	closed   bool
	mutex    sync.RWMutex
	wg       sync.WaitGroup
	Canceller
}

func NewStream(capacity ...int) *Stream {
	cap := Capacity(capacity...)
	s := &Stream{
		cap,
		make(chan interface{}, cap),
		make(chan []Event, cap),
		[]Window{},
		false,
		sync.RWMutex{},
		sync.WaitGroup{},
		NewCanceller(),
	}

	go s.fanout()
	return s
}

func (s *Stream) Close() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.IsClosed() {
		return
	}

	s.closed = true
	s.cancel()

	s.wg.Wait()
	close(s.Input())
	close(s.Output())
	for _, w := range s.window {
		w.Close()
	}
}

func (s *Stream) IsClosed() bool {
	return s.closed
}

func (s *Stream) SetWindow(w Window) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.window = append(s.window, w)
	go w.Work()
	go s.fanin(w)
}

func (s *Stream) Window() []Window {
	return s.window
}

func (s *Stream) Input() chan interface{} {
	return s.in
}

func (s *Stream) Output() chan []Event {
	return s.out
}

func (s *Stream) fanout() {
	s.wg.Add(1)
	for {
		select {
		case <-s.ctx.Done():
			s.wg.Done()
			return
		case input := <-s.in:
			s.mutex.RLock()
			for _, w := range s.window {
				// TODO: need deep copy
				w.Input() <- input
			}
			s.mutex.RUnlock()
		}
	}
}

func (s *Stream) fanin(w Window) {
	s.wg.Add(1)
	for {
		select {
		case <-s.ctx.Done():
			s.wg.Done()
			return
		case event := <-w.Output():
			s.out <- event
		}
	}
}
