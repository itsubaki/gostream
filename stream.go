package gocep

type Stream struct {
	capacity int
	in       chan interface{}
	out      chan []Event
	window   []Window
	insert   *Stream
	Canceller
}

func NewStream(capacity int) *Stream {
	s := &Stream{
		capacity,
		make(chan interface{}, capacity),
		make(chan []Event, capacity),
		[]Window{},
		nil,
		NewCanceller(),
	}

	go s.dispatch()
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
	go s.collect(w)
}

func (s *Stream) InsertInto(stream *Stream) {
	s.insert = stream
	go s.transfer()
}

func (s *Stream) Input() chan interface{} {
	return s.in
}

func (s *Stream) Output() chan []Event {
	return s.out
}

func (s *Stream) dispatch() {
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

func (s *Stream) collect(w Window) {
	for {
		select {
		case <-s.ctx.Done():
			return
		case event := <-w.Output():
			s.out <- event
		}
	}
}

func (s *Stream) transfer() {
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
