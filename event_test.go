package gocep

import (
	"testing"
)

type IntEvent struct {
	Name  string
	Value int
}

type FloatEvent struct {
	Name  string
	Value float32
}

type BoolEvent struct {
	Value bool
}

func TestFloatValue(t *testing.T) {
	event := NewEvent(FloatEvent{"foobar", 12.3})

	if event.Float32Value("Value") != 12.3 {
		t.Errorf("failed.")
	}
}

func TestBoolValue(t *testing.T) {
	event := NewEvent(BoolEvent{true})

	if !event.BoolValue("Value") {
		t.Errorf("failed.")
	}
}

func TestStringValue(t *testing.T) {
	event := NewEvent(IntEvent{"foobar", 123})

	if event.StringValue("Name") != "foobar" {
		t.Errorf("failed.")
	}
}
