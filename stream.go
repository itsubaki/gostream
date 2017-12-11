package gocep

type Stream struct {
	capacity int
	in       chan interface{}
	out      chan []Event
	window   []Window
	Canceller
}

func NewStream(capacity ...int) *Stream {
	cap := Capacity(capacity...)
	s := &Stream{
		cap,
		make(chan interface{}, cap),
		make(chan []Event, cap),
		[]Window{},
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

func (s *Stream) SetWindow(w Window) {
	s.window = append(s.window, w)
	go s.collect(w)
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
