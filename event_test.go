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
	Value float64
}

type BoolEvent struct {
	Value bool
}

func TestFloat(t *testing.T) {
	event := NewEvent(FloatEvent{"foobar", 12.3})

	if event.Float("Value") != 12.3 {
		t.Errorf("failed.")
	}
}

func TestBool(t *testing.T) {
	event := NewEvent(BoolEvent{true})

	if !event.Bool("Value") {
		t.Errorf("failed.")
	}
}

func TestString(t *testing.T) {
	event := NewEvent(IntEvent{"foobar", 123})

	if event.String("Name") != "foobar" {
		t.Errorf("failed.")
	}
}

func TestMap(t *testing.T) {
	m := make(map[string]interface{})
	m["foo"] = "bar"
	m["piyo"] = 123
	m["hoge"] = 12.3
	m["fuga"] = false
	e := NewEvent(MapEvent{m})

	if e.MapString("Record", "foo") != "bar" {
		t.Error(e)
	}
	if e.MapInt("Record", "piyo") != 123 {
		t.Error(e)
	}
	if e.MapFloat("Record", "hoge") != 12.3 {
		t.Error(e)
	}
	if e.MapBool("Record", "fuga") {
		t.Error(e)
	}
}

func TestRecordString(t *testing.T) {
	event := NewEvent(IntEvent{"foobar", 123})
	event.Record["Name"] = "foobar"

	if event.RecordString("Name") != "foobar" {
		t.Errorf("failed.")
	}
}
