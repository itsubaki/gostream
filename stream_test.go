package gocep

import (
	"testing"
)

func TestStream(t *testing.T) {
	s := NewStream()
	defer s.Close()

	wnum := 2
	for i := 0; i < wnum; i++ {
		w := NewIdentityWindow(16)
		go w.Work()
		s.SetWindow(w)
	}

	if len(s.Window()) != wnum {
		t.Error("failed.")
	}

	s.Input() <- "test"

	for i := 0; i < wnum; i++ {
		if Oldest(<-s.Output()).Underlying != "test" {
			t.Error("failed")
		}
	}

}
