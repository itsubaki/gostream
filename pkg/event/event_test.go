package event_test

import (
	"testing"

	"github.com/itsubaki/gostream/pkg/event"
)

func TestFloat(t *testing.T) {
	type FloatEvent struct {
		Name  string
		Value float64
	}

	event := event.New(FloatEvent{"foobar", 12.3})
	if event.Float("Value") != 12.3 {
		t.Errorf("failed.")
	}
}

func TestBool(t *testing.T) {
	type BoolEvent struct {
		Value bool
	}

	event := event.New(BoolEvent{true})
	if !event.Bool("Value") {
		t.Errorf("failed.")
	}
}

func TestString(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	event := event.New(IntEvent{"foobar", 123})
	if event.String("Name") != "foobar" {
		t.Errorf("failed.")
	}
}

func TestRecordString(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	event := event.New(IntEvent{"foobar", 123})
	event.Record["Name"] = "foobar"

	if event.RecordString("Name") != "foobar" {
		t.Errorf("failed.")
	}
}
