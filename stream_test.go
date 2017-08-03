package gocep

import (
	"testing"
)

func TestStream(t *testing.T) {
	s := NewStream(32)
	defer s.Close()

	wnum := 2
	for i := 0; i < wnum; i++ {
		s.SetWindow(NewIdentityWindow(16))
	}

	if len(s.Window()) != wnum {
		t.Error("failed.")
	}

	s.Input() <- "test"

	for i := 0; i < wnum; i++ {
		e := <-s.Output()
		if e[0].Underlying != "test" {
			t.Error("failed")
		}
	}

}

func TestStreamInsert(t *testing.T) {
	stream := NewStream(32)
	defer stream.Close()
	stream.SetWindow(NewIdentityWindow(16))

	insert := NewStream(32)
	defer insert.Close()
	insert.SetWindow(NewIdentityWindow(16))

	stream.InsertInto(insert)

	stream.Input() <- "test"
	e := <-insert.Output()
	if len(e) != 1 {
		t.Error(e)
	}
}
