package gocep

type Stream struct {
	capacity int
	in       chan interface{}
	out      chan []Event
	close    chan bool
	window   []Window
}

func NewStream(capacity int) *Stream {
	s := &Stream{
		capacity,
		make(chan interface{}, capacity),
		make(chan []Event, capacity),
		make(chan bool, 1),
		[]Window{},
	}

	go s.dispatch()
	return s
}

func (s *Stream) Close() {
	s.close <- true // dispacth()
	for _, w := range s.window {
		s.close <- true // collect()
		w.Close()
	}
}

func (s *Stream) Add(w Window) {
	s.window = append(s.window, w)
	go s.collect(w)
}

func (s *Stream) dispatch() {
	for {
		select {
		case c := <-s.close:
			if c {
				return
			}
		case e := <-s.in:
			for _, w := range s.window {
				w.Input() <- NewEvent(e)
			}
		}
	}
}

func (s *Stream) collect(w Window) {
	for {
		select {
		case c := <-s.close:
			if c {
				return
			}
		case e := <-w.Output():
			s.out <- e
		}
	}
}

func (s *Stream) Input() chan interface{} {
	return s.in
}

func (s *Stream) Output() chan []Event {
	return s.out
}
