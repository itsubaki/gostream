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
	e := NewEvent(MapEvent{"foobar", m})

	if e.MapString("Map", "foo") != "bar" {
		t.Error(e)
	}
	if e.MapInt("Map", "piyo") != 123 {
		t.Error(e)
	}
	if e.MapFloat("Map", "hoge") != 12.3 {
		t.Error(e)
	}
	if e.MapBool("Map", "fuga") {
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
