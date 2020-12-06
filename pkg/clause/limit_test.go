package clause_test

import (
	"testing"

	"github.com/itsubaki/gostream/pkg/clause"
	"github.com/itsubaki/gostream/pkg/event"
)

func TestFirst(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	v := clause.First{}

	events := event.List()
	empty := v.Apply(events)
	if len(empty) != 0 {
		t.Error(empty)
	}

	events = append(events, event.New(IntEvent{"foo", 10}))
	events = append(events, event.New(IntEvent{"foo", 20}))
	result := v.Apply(events)

	if len(result) != 1 {
		t.Error(result)
	}

	if result[0].Int("Value") != 10 {
		t.Error(result)
	}
}

func TestLast(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	v := clause.Last{}

	events := event.List()
	empty := v.Apply(events)
	if len(empty) != 0 {
		t.Error(empty)
	}

	events = append(events, event.New(IntEvent{"foo", 10}))
	events = append(events, event.New(IntEvent{"foo", 20}))
	result := v.Apply(events)

	if len(result) != 1 {
		t.Error(result)
	}

	if result[0].Int("Value") != 20 {
		t.Error(result)
	}
}

func TestLimit(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	v := clause.Limit{4, 2}

	events := event.List()
	empty := v.Apply(events)
	if len(empty) != 0 {
		t.Error(empty)
	}

	events = append(events, event.New(IntEvent{"foo", 10}))
	events = append(events, event.New(IntEvent{"foo", 20}))
	events = append(events, event.New(IntEvent{"foo", 30}))
	events = append(events, event.New(IntEvent{"foo", 40}))
	events = append(events, event.New(IntEvent{"foo", 50}))
	events = append(events, event.New(IntEvent{"foo", 60}))
	result := v.Apply(events)

	if len(result) != 4 {
		t.Error(result)
	}

	if result[0].Int("Value") != 30 {
		t.Error(result)
	}
	if result[3].Int("Value") != 60 {
		t.Error(result)
	}
}
