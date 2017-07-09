package gocep

import (
	"fmt"
	"testing"
)

func TestStream(t *testing.T) {
	s := NewStream(32)
	defer s.Close()

	wnum := 2
	for i := 0; i < wnum; i++ {
		s.Window(NewSimpleWindow(16))
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
	stream.Window(NewSimpleWindow(16))

	insert := NewStream(32)
	defer insert.Close()
	insert.Window(NewSimpleWindow(16))

	stream.InsertInto(insert)

	stream.Input() <- "test"
	e := <-insert.Output()
	fmt.Println(e)
}
