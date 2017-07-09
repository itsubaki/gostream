package gocep

import "context"

type Stream struct {
	capacity int
	in       chan interface{}
	out      chan []Event
	ctx      context.Context
	cancel   func()
	window   []Window
	insert   *Stream
}

func NewStream(capacity int) *Stream {
	ctx, cancel := context.WithCancel(context.Background())
	s := &Stream{
		capacity,
		make(chan interface{}, capacity),
		make(chan []Event, capacity),
		ctx,
		cancel,
		[]Window{},
		nil,
	}

	go s.dispatch(s.ctx)
	return s
}

func (s *Stream) Close() {
	s.cancel()
	for _, w := range s.window {
		w.Close()
	}
}

func (s *Stream) Window(w Window) {
	s.window = append(s.window, w)
	go s.collect(s.ctx, w)
}

func (s *Stream) InsertInto(stream *Stream) {
	s.insert = stream
	go s.transfer(s.ctx)
}

func (s *Stream) Input() chan interface{} {
	return s.in
}

func (s *Stream) Output() chan []Event {
	return s.out
}

func (s *Stream) dispatch(ctx context.Context) {
	for {
		select {
		case <-s.ctx.Done():
			return
		case input := <-s.in:
			for _, w := range s.window {
				w.Input() <- input
			}
		}
	}
}

func (s *Stream) collect(ctx context.Context, w Window) {
	for {
		select {
		case <-s.ctx.Done():
			return
		case event := <-w.Output():
			s.out <- event
		}
	}
}

func (s *Stream) transfer(ctx context.Context) {
	for {
		select {
		case <-s.ctx.Done():
			return
		case input := <-s.Output():
			for _, e := range input {
				s.insert.Input() <- MapEvent{e.Record}
			}
		}
	}
}
