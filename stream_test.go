package gocep

import (
	"testing"
)

func TestStream(t *testing.T) {
	s := NewStream(10)
	defer s.Close()

	w := NewLengthWindow(10, 10)
	s.Add(w)
	s.Push("test")

	event := <-w.Output()
	if event[0].Underlying != "test" {
		t.Error(event)
	}
}
