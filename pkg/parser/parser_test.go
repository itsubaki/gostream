package parser

import (
	"reflect"
	"testing"
	"time"
)

func TestParserSum(t *testing.T) {
	type LogEvent struct {
		ID      string
		Time    time.Time
		Level   int
		Message string
	}

	p := New()
	p.Register("LogEvent", LogEvent{})

	q := "select sum(Level) from LogEvent.time(10 sec) where Level > 2 and Level < 10"
	st, err := p.Parse(q)
	if err != nil {
		t.Error(err)
	}

	if reflect.TypeOf(st.Function[0]).Name() != "SumInt" {
		t.Error(reflect.TypeOf(st.Function[0]).Name())
	}
}

func TestParserAvg(t *testing.T) {
	type LogEvent struct {
		ID      string
		Time    time.Time
		Level   float64
		Message string
	}

	p := New()
	p.Register("LogEvent", LogEvent{})

	q := "select avg(Level) from LogEvent.time(10 sec) where Level > 2.5 and Level < 10.5"
	st, err := p.Parse(q)
	if err != nil {
		t.Error(err)
	}

	if reflect.TypeOf(st.Function[0]).Name() != "AverageInt" {
		t.Error(reflect.TypeOf(st.Function[0]).Name())
	}

	if reflect.TypeOf(st.Selector[1]).Name() != "LargerThanFloat" {
		t.Error(reflect.TypeOf(st.Selector[1]).Name())
	}

	if reflect.TypeOf(st.Selector[2]).Name() != "LessThanFloat" {
		t.Error(reflect.TypeOf(st.Selector[2]).Name())
	}
}

func TestParserTimeWindow(t *testing.T) {
	type LogEvent struct {
		ID      string
		Time    time.Time
		Level   int
		Message string
	}

	p := New()
	p.Register("LogEvent", LogEvent{})

	q := "select count(*) from LogEvent.time(10 sec) where Level > 2 and Level < 10"
	st, err := p.Parse(q)
	if err != nil {
		t.Error(err)
	}

	if st.Time != 10*time.Second {
		t.Fail()
	}

	if reflect.TypeOf(st.Function[0]).Name() != "Count" {
		t.Fail()
	}

	if reflect.TypeOf(st.Selector[0]).Name() != "EqualsType" {
		t.Fail()
	}

	if reflect.TypeOf(st.Selector[1]).Name() != "LargerThanInt" {
		t.Fail()
	}

	if reflect.TypeOf(st.Selector[2]).Name() != "LessThanInt" {
		t.Fail()
	}
}

func TestParserError(t *testing.T) {
	p := New()

	q := "select * from MapEvent.length(10)"
	_, err := p.Parse(q)
	if err == nil {
		t.Error("failed.")
	}

	if err.Error() != "parse event type: EventType [MapEvent] is not registered" {
		t.Errorf("failed: %v", err)
	}
}

func TestNewStatementLength(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	p := New()
	p.Register("MapEvent", MapEvent{})

	q := "select * from MapEvent.length(10)"
	stmt, err := p.Parse(q)
	if err != nil {
		t.Error(err)
		return
	}
	window := stmt.New(1024)
	defer window.Close()

	m := make(map[string]interface{})
	m["Value"] = "foobar"

	window.Input() <- MapEvent{m}
	event := <-window.Output()
	if event[0].MapString("Record", "Value") != "foobar" {
		t.Error(event)
	}
}

func TestNewStatementTime(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	p := New()
	p.Register("MapEvent", MapEvent{})

	q := "select * from MapEvent.time(10 sec)"
	stmt, err := p.Parse(q)
	if err != nil {
		t.Error(err)
	}

	window := stmt.New(1024)
	defer window.Close()

	m := make(map[string]interface{})
	m["Value"] = "foobar"

	window.Input() <- MapEvent{m}
	event := <-window.Output()
	if event[0].MapString("Record", "Value") != "foobar" {
		t.Error(event)
	}
}
