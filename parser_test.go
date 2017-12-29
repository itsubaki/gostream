package gocep

import (
	"reflect"
	"testing"
	"time"
)

func TestParserTimeWindow(t *testing.T) {
	type LogEvent struct {
		ID      string
		Time    time.Time
		Level   int
		Message string
	}

	p := NewParser()
	p.Register("LogEvent", LogEvent{})

	q := "select count(*) from LogEvent.time(10 sec) where Level > 2"
	stmt, err := p.Parse(q)
	if err != nil {
		t.Error(err)
	}

	if stmt.time != 10*time.Second {
		t.Fail()
	}

	if reflect.TypeOf(stmt.function[0]).Name() != "Count" {
		t.Fail()
	}

	if reflect.TypeOf(stmt.selector[0]).Name() != "EqualsType" {
		t.Fail()
	}
}

func TestParserError(t *testing.T) {
	p := NewParser()

	q := "select * from MapEvent.length(10)"
	_, err := p.Parse(q)
	if err == nil {
		t.Error("failed.")
	}

	if err.Error() != "EventType [MapEvent] is not registered" {
		t.Error("failed.")
	}
}

func TestParser(t *testing.T) {
	p := NewParser()
	p.Register("MapEvent", MapEvent{})

	q := "select * from MapEvent.length(10)"
	stmt, err := p.Parse(q)
	if err != nil {
		t.Error(err)
		return
	}
	window := stmt.New(1024)

	m := make(map[string]interface{})
	m["Value"] = "foobar"

	window.Input() <- MapEvent{m}
	event := <-window.Output()
	if event[0].RecordString("Value") != "foobar" {
		t.Error(event)
	}
}

func TestNewStream(t *testing.T) {
	p := NewParser()
	p.Register("MapEvent", MapEvent{})

	q := "select * from MapEvent.length(10)"
	stmt, err := p.Parse(q)
	if err != nil {
		t.Error(err)
		return
	}

	st := stmt.NewStream(1024)

	m := make(map[string]interface{})
	m["Value"] = "foobar"

	st.Input() <- MapEvent{m}
	event := <-st.Output()
	if event[0].RecordString("Value") != "foobar" {
		t.Error(event)
	}
}
