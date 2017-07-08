package gocep

import "testing"

func TestStream(t *testing.T) {
	s := NewStream(10)
	defer s.Close()

	wnum := 2
	for i := 0; i < wnum; i++ {
		s.Add(NewSimpleWindow(16))
	}

	s.Input() <- "test"

	for i := 0; i < wnum; i++ {
		e := <-s.Output()
		if e[0].Underlying != "test" {
			t.Error("failed")
		}
	}

}
