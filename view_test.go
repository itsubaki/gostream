package gocep

import (
	"testing"
)

func TestFirst(t *testing.T) {
	v := First{}

	event := []Event{}
	empty := v.Apply(event)
	if len(empty) != 0 {
		t.Error(empty)
	}

	event = append(event, NewEvent(IntEvent{"foo", 10}))
	event = append(event, NewEvent(IntEvent{"foo", 20}))
	result := v.Apply(event)

	if len(result) != 1 {
		t.Error(result)
	}

	if result[0].Int("Value") != 10 {
		t.Error(result)
	}
}

func TestLast(t *testing.T) {
	v := Last{}

	event := []Event{}
	empty := v.Apply(event)
	if len(empty) != 0 {
		t.Error(empty)
	}

	event = append(event, NewEvent(IntEvent{"foo", 10}))
	event = append(event, NewEvent(IntEvent{"foo", 20}))
	result := v.Apply(event)

	if len(result) != 1 {
		t.Error(result)
	}

	if result[0].Int("Value") != 20 {
		t.Error(result)
	}
}

func TestLimit(t *testing.T) {
	v := Limit{4, 2}

	event := []Event{}
	empty := v.Apply(event)
	if len(empty) != 0 {
		t.Error(empty)
	}

	event = append(event, NewEvent(IntEvent{"foo", 10}))
	event = append(event, NewEvent(IntEvent{"foo", 20}))
	event = append(event, NewEvent(IntEvent{"foo", 30}))
	event = append(event, NewEvent(IntEvent{"foo", 40}))
	event = append(event, NewEvent(IntEvent{"foo", 50}))
	event = append(event, NewEvent(IntEvent{"foo", 60}))
	result := v.Apply(event)

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
	v := OrderByInt{"Value", false}

	event := []Event{}

	event = append(event, NewEvent(IntEvent{"foo", 10}))
	event = append(event, NewEvent(IntEvent{"foo", 30}))
	event = append(event, NewEvent(IntEvent{"foo", 20}))
	event = append(event, NewEvent(IntEvent{"foo", 40}))
	event = append(event, NewEvent(IntEvent{"foo", 60}))
	event = append(event, NewEvent(IntEvent{"foo", 50}))
	result := v.Apply(event)

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
	v := OrderByInt{"Value", true}

	event := []Event{}

	event = append(event, NewEvent(IntEvent{"foo", 10}))
	event = append(event, NewEvent(IntEvent{"foo", 30}))
	event = append(event, NewEvent(IntEvent{"foo", 20}))
	event = append(event, NewEvent(IntEvent{"foo", 40}))
	event = append(event, NewEvent(IntEvent{"foo", 60}))
	event = append(event, NewEvent(IntEvent{"foo", 50}))
	result := v.Apply(event)

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
	v := OrderByFloat{"Value", false}

	event := []Event{}

	event = append(event, NewEvent(FloatEvent{"foo", 10.0}))
	event = append(event, NewEvent(FloatEvent{"foo", 30.0}))
	event = append(event, NewEvent(FloatEvent{"foo", 20.0}))
	event = append(event, NewEvent(FloatEvent{"foo", 40.0}))
	event = append(event, NewEvent(FloatEvent{"foo", 60.0}))
	event = append(event, NewEvent(FloatEvent{"foo", 50.0}))
	result := v.Apply(event)

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
	v := OrderByFloat{"Value", true}

	event := []Event{}
	event = append(event, NewEvent(FloatEvent{"foo", 10.0}))
	event = append(event, NewEvent(FloatEvent{"foo", 30.0}))
	event = append(event, NewEvent(FloatEvent{"foo", 20.0}))
	event = append(event, NewEvent(FloatEvent{"foo", 40.0}))
	event = append(event, NewEvent(FloatEvent{"foo", 60.0}))
	event = append(event, NewEvent(FloatEvent{"foo", 50.0}))
	result := v.Apply(event)

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
