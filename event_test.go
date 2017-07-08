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

type MapEvent struct {
	Name string
	Map  map[string]interface{}
}

func TestFloatValue(t *testing.T) {
	event := NewEvent(FloatEvent{"foobar", 12.3})

	if event.FloatValue("Value") != 12.3 {
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

func TestMapValue(t *testing.T) {
	m := make(map[string]interface{})
	m["foo"] = "bar"
	m["piyo"] = 123
	m["hoge"] = 12.3
	m["fuga"] = false
	e := NewEvent(MapEvent{"foobar", m})

	if e.MapStringValue("Map", "foo") != "bar" {
		t.Error(e)
	}
	if e.MapIntValue("Map", "piyo") != 123 {
		t.Error(e)
	}
	if e.MapFloatValue("Map", "hoge") != 12.3 {
		t.Error(e)
	}
	if e.MapBoolValue("Map", "fuga") {
		t.Error(e)
	}
}

func TestRecordStringValue(t *testing.T) {
	event := NewEvent(IntEvent{"foobar", 123})
	event.Record["Name"] = "foobar"

	if event.RecordStringValue("Name") != "foobar" {
		t.Errorf("failed.")
	}
}
