package gocep

import (
	"testing"
)

func TestParser(t *testing.T) {
	q := "select * from MapEvent.length(10)"

	p := NewParser(q, 1024)
	stmt, _ := p.Parse()

	w := stmt.window
	for _, s := range stmt.selector {
		w.Selector(s)
	}

	for _, f := range stmt.function {
		w.Function(f)
	}

	for _, v := range stmt.view {
		w.View(v)
	}

	m := make(map[string]interface{})
	m["Value"] = "foobar"

	w.Input() <- MapEvent{m}
	event := <-w.Output()
	if event[0].RecordString("Value") != "foobar" {
		t.Error(event)
	}
}
