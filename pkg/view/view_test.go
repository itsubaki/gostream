package view

import (
	"testing"

	"github.com/itsubaki/gocep/pkg/event"
)

func TestFirst(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	v := First{}

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

	v := Last{}

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

	v := Limit{4, 2}

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

func TestOrderByInt(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	events := event.List()
	events = append(events, event.New(IntEvent{"foo", 10}))
	events = append(events, event.New(IntEvent{"foo", 30}))
	events = append(events, event.New(IntEvent{"foo", 20}))
	events = append(events, event.New(IntEvent{"foo", 40}))
	events = append(events, event.New(IntEvent{"foo", 60}))
	events = append(events, event.New(IntEvent{"foo", 50}))

	v := OrderByInt{"Value", false}
	result := v.Apply(events)

	var test = []struct {
		index int
		value int
	}{
		{0, 10},
		{1, 20},
		{2, 30},
		{3, 40},
		{4, 50},
		{5, 60},
	}

	for _, tt := range test {
		if result[tt.index].Int("Value") != tt.value {
			t.Error(result)
		}
	}
}

func TestOrderByIntReverse(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	events := event.List()
	events = append(events, event.New(IntEvent{"foo", 10}))
	events = append(events, event.New(IntEvent{"foo", 30}))
	events = append(events, event.New(IntEvent{"foo", 20}))
	events = append(events, event.New(IntEvent{"foo", 40}))
	events = append(events, event.New(IntEvent{"foo", 60}))
	events = append(events, event.New(IntEvent{"foo", 50}))

	v := OrderByInt{"Value", true}
	result := v.Apply(events)

	var test = []struct {
		index int
		value int
	}{
		{0, 60},
		{1, 50},
		{2, 40},
		{3, 30},
		{4, 20},
		{5, 10},
	}

	for _, tt := range test {
		if result[tt.index].Int("Value") != tt.value {
			t.Error(result)
		}
	}
}

func TestOrderByFloat(t *testing.T) {
	type FloatEvent struct {
		Name  string
		Value float64
	}

	events := event.List()
	events = append(events, event.New(FloatEvent{"foo", 10.0}))
	events = append(events, event.New(FloatEvent{"foo", 30.0}))
	events = append(events, event.New(FloatEvent{"foo", 20.0}))
	events = append(events, event.New(FloatEvent{"foo", 40.0}))
	events = append(events, event.New(FloatEvent{"foo", 60.0}))
	events = append(events, event.New(FloatEvent{"foo", 50.0}))

	v := OrderByFloat{"Value", false}
	result := v.Apply(events)

	var test = []struct {
		index int
		value float64
	}{
		{0, 10.0},
		{1, 20.0},
		{2, 30.0},
		{3, 40.0},
		{4, 50.0},
		{5, 60.0},
	}

	for _, tt := range test {
		if result[tt.index].Float("Value") != tt.value {
			t.Error(result)
		}
	}
}

func TestOrderByFloatReverse(t *testing.T) {
	type FloatEvent struct {
		Name  string
		Value float64
	}

	events := event.List()
	events = append(events, event.New(FloatEvent{"foo", 10.0}))
	events = append(events, event.New(FloatEvent{"foo", 30.0}))
	events = append(events, event.New(FloatEvent{"foo", 20.0}))
	events = append(events, event.New(FloatEvent{"foo", 40.0}))
	events = append(events, event.New(FloatEvent{"foo", 60.0}))
	events = append(events, event.New(FloatEvent{"foo", 50.0}))

	v := OrderByFloat{"Value", true}
	result := v.Apply(events)

	var test = []struct {
		index int
		value float64
	}{
		{0, 60.0},
		{1, 50.0},
		{2, 40.0},
		{3, 30.0},
		{4, 20.0},
		{5, 10.0},
	}

	for _, tt := range test {
		if result[tt.index].Float("Value") != tt.value {
			t.Error(result)
		}
	}
}
