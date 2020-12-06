package parser_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/itsubaki/gostream/pkg/parser"
)

func TestParserFloat(t *testing.T) {
	type LogEvent struct {
		ID      string
		Time    time.Time
		Level   float64
		Message string
	}

	p := parser.New()
	p.Register("LogEvent", LogEvent{})

	q := "select avg(Level), sum(Level), max(Level), min(Level), med(Level), count(*) from LogEvent.length(10) where Level > 2.5 and Level < 10.5"
	s, err := p.Parse(q)
	if err != nil {
		t.Error(err)
	}

	if s.Length != 10 {
		t.Fail()
	}

	if reflect.TypeOf(s.EventType) != reflect.TypeOf(LogEvent{}) {
		t.Fail()
	}

	if reflect.TypeOf(s.Function[0]).Name() != "AverageFloat" {
		t.Error(reflect.TypeOf(s.Function[0]).Name())
	}

	if reflect.TypeOf(s.Function[1]).Name() != "SumFloat" {
		t.Error(reflect.TypeOf(s.Function[1]).Name())
	}

	if reflect.TypeOf(s.Function[2]).Name() != "MaxFloat" {
		t.Error(reflect.TypeOf(s.Function[2]).Name())
	}

	if reflect.TypeOf(s.Function[3]).Name() != "MinFloat" {
		t.Error(reflect.TypeOf(s.Function[3]).Name())
	}

	if reflect.TypeOf(s.Function[4]).Name() != "MedianFloat" {
		t.Error(reflect.TypeOf(s.Function[4]).Name())
	}

	if reflect.TypeOf(s.Function[5]).Name() != "Count" {
		t.Error(reflect.TypeOf(s.Function[5]).Name())
	}

	if reflect.TypeOf(s.Where[0]).Name() != "LargerThanFloat" {
		t.Error(reflect.TypeOf(s.Where[0]).Name())
	}

	if reflect.TypeOf(s.Where[1]).Name() != "LessThanFloat" {
		t.Error(reflect.TypeOf(s.Where[1]).Name())
	}
}

func TestParserInt(t *testing.T) {
	type LogEvent struct {
		ID      string
		Time    time.Time
		Level   int
		Message string
	}

	p := parser.New()
	p.Register("LogEvent", LogEvent{})

	q := "select avg(Level), sum(Level), max(Level), min(Level), med(Level), count(*) from LogEvent.time(10 sec) where Level > 2 and Level < 10"
	s, err := p.Parse(q)
	if err != nil {
		t.Error(err)
	}

	if s.Time != 10*time.Second {
		t.Fail()
	}

	if reflect.TypeOf(s.EventType) != reflect.TypeOf(LogEvent{}) {
		t.Fail()
	}

	if reflect.TypeOf(s.Function[0]).Name() != "AverageInt" {
		t.Error(reflect.TypeOf(s.Function[0]).Name())
	}

	if reflect.TypeOf(s.Function[1]).Name() != "SumInt" {
		t.Error(reflect.TypeOf(s.Function[1]).Name())
	}

	if reflect.TypeOf(s.Function[2]).Name() != "MaxInt" {
		t.Error(reflect.TypeOf(s.Function[2]).Name())
	}

	if reflect.TypeOf(s.Function[3]).Name() != "MinInt" {
		t.Error(reflect.TypeOf(s.Function[3]).Name())
	}

	if reflect.TypeOf(s.Function[4]).Name() != "MedianInt" {
		t.Error(reflect.TypeOf(s.Function[4]).Name())
	}

	if reflect.TypeOf(s.Function[5]).Name() != "Count" {
		t.Error(reflect.TypeOf(s.Function[5]).Name())
	}

	if reflect.TypeOf(s.Where[0]).Name() != "LargerThanInt" {
		t.Error(reflect.TypeOf(s.Where[0]).Name())
	}

	if reflect.TypeOf(s.Where[1]).Name() != "LessThanInt" {
		t.Error(reflect.TypeOf(s.Where[1]).Name())
	}
}

func TestParserError(t *testing.T) {
	p := parser.New()

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

	p := parser.New()
	p.Register("MapEvent", MapEvent{})

	q := "select * from MapEvent.length(10)"
	s, err := p.Parse(q)
	if err != nil {
		t.Error(err)
		return
	}
	w := s.New(1024)
	defer w.Close()

	m := make(map[string]interface{})
	m["Value"] = "foobar"

	w.Input() <- MapEvent{m}
	event := <-w.Output()
	if event[0].MapString("Record", "Value") != "foobar" {
		t.Error(event)
	}
}

func TestNewStatementTime(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	p := parser.New()
	p.Register("MapEvent", MapEvent{})

	q := "select * from MapEvent.time(10 min)"
	s, err := p.Parse(q)
	if err != nil {
		t.Error(err)
	}

	w := s.New()
	defer w.Close()

	m := make(map[string]interface{})
	m["Value"] = "foobar"

	w.Input() <- MapEvent{m}
	event := <-w.Output()
	if event[0].MapString("Record", "Value") != "foobar" {
		t.Error(event)
	}
}
