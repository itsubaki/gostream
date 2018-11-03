package gocep

import (
	"testing"
)

func TestStream(t *testing.T) {
	s := NewStream()
	s.SetWindow(NewIdentityWindow())
	defer s.Close()

	s.Input() <- "test"
	if Oldest(<-s.Output()).Underlying != "test" {
		t.Error("failed")
	}
}
