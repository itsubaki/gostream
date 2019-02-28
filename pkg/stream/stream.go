package gocep

import (
	"sync"

	"github.com/itsubaki/gocep/pkg/event"
	"github.com/itsubaki/gocep/pkg/window"
)

type Stream struct {
	capacity int
	in       chan interface{}
	out      chan []event.Event
	window   []window.Window
	closed   bool
	mutex    sync.RWMutex
	wg       sync.WaitGroup
	window.Canceller
}

func New(capacity ...int) *Stream {
	cap := window.Capacity(capacity...)
	s := &Stream{
		cap,
		make(chan interface{}, cap),
		make(chan []event.Event, cap),
		[]window.Window{},
		false,
		sync.RWMutex{},
		sync.WaitGroup{},
		window.NewCanceller(),
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
	s.Cancel()

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

func (s *Stream) AddWindow(w window.Window) {
	s.window = append(s.window, w)
	go s.fanin(w)
}

func (s *Stream) Window() []window.Window {
	return s.window
}

func (s *Stream) Input() chan interface{} {
	return s.in
}

func (s *Stream) Output() chan []event.Event {
	return s.out
}

func (s *Stream) fanout() {
	s.wg.Add(1)
	for {
		select {
		case <-s.Context.Done():
			s.wg.Done()
			return
		case input := <-s.in:
			for _, w := range s.window {
				w.Input() <- input
			}
		}
	}
}

func (s *Stream) fanin(w window.Window) {
	s.wg.Add(1)
	for {
		select {
		case <-s.Context.Done():
			s.wg.Done()
			return
		case event := <-w.Output():
			s.out <- event
		}
	}
}
