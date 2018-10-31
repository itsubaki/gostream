package gocep

import (
	"testing"
)

func TestStream(t *testing.T) {
	s := NewStream()
	defer s.Close()

	n := 2
	for i := 0; i < n; i++ {
		s.SetWindow(NewIdentityWindow())
	}

	if len(s.Window()) != n {
		t.Error("failed.")
	}

	s.Input() <- "test"

	for i := 0; i < n; i++ {
		if Oldest(<-s.Output()).Underlying != "test" {
			t.Error("failed")
		}
	}
}
