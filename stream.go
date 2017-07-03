package gocep

type Stream struct {
	capacity int
	queue    chan Event
	close    chan bool
	window   []Window
}

func NewStream(capacity int) *Stream {
	s := &Stream{
		capacity,
		make(chan Event, capacity),
		make(chan bool, 1),
		[]Window{},
	}

	go s.dispatch()
	return s
}

func (s *Stream) Close() {
	s.close <- true
	for _, w := range s.window {
		w.Close()
	}
}

func (s *Stream) Add(w Window) {
	s.window = append(s.window, w)
}

func (s *Stream) dispatch() {
	for {
		select {
		case c := <-s.close:
			if c {
				return
			}
		case e := <-s.queue:
			for _, w := range s.window {
				w.Input() <- e
			}
		}
	}
}

func (s *Stream) Push(event interface{}) {
	s.queue <- NewEvent(event)
}
