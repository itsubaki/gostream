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

	event = append(event, NewEvent(IntEvent{"foo", 10}).New())
	event = append(event, NewEvent(IntEvent{"foo", 20}).New())
	result := v.Apply(event)

	if len(result) != 1 {
		t.Error(result)
	}

	if result[0].IntValue("Value") != 10 {
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

	event = append(event, NewEvent(IntEvent{"foo", 10}).New())
	event = append(event, NewEvent(IntEvent{"foo", 20}).New())
	result := v.Apply(event)

	if len(result) != 1 {
		t.Error(result)
	}

	if result[0].IntValue("Value") != 20 {
		t.Error(result)
	}
}

func TestLimit(t *testing.T) {
	v := Limit{2, 4}

	event := []Event{}
	empty := v.Apply(event)
	if len(empty) != 0 {
		t.Error(empty)
	}

	event = append(event, NewEvent(IntEvent{"foo", 10}).New())
	event = append(event, NewEvent(IntEvent{"foo", 20}).New())
	event = append(event, NewEvent(IntEvent{"foo", 30}).New())
	event = append(event, NewEvent(IntEvent{"foo", 40}).New())
	event = append(event, NewEvent(IntEvent{"foo", 50}).New())
	event = append(event, NewEvent(IntEvent{"foo", 60}).New())
	result := v.Apply(event)

	if len(result) != 4 {
		t.Error(result)
	}

	if result[0].IntValue("Value") != 30 {
		t.Error(result)
	}
	if result[3].IntValue("Value") != 60 {
		t.Error(result)
	}
}

func TestSortInt(t *testing.T) {
	v := SortInt{"Value", false}

	event := []Event{}

	event = append(event, NewEvent(IntEvent{"foo", 10}).New())
	event = append(event, NewEvent(IntEvent{"foo", 30}).New())
	event = append(event, NewEvent(IntEvent{"foo", 20}).New())
	event = append(event, NewEvent(IntEvent{"foo", 40}).New())
	event = append(event, NewEvent(IntEvent{"foo", 60}).New())
	event = append(event, NewEvent(IntEvent{"foo", 50}).New())
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
		if result[tt.index].IntValue("Value") != tt.value {
			t.Error(result)
		}
	}
}

func TestSortIntReverse(t *testing.T) {
	v := SortInt{"Value", true}

	event := []Event{}

	event = append(event, NewEvent(IntEvent{"foo", 10}).New())
	event = append(event, NewEvent(IntEvent{"foo", 30}).New())
	event = append(event, NewEvent(IntEvent{"foo", 20}).New())
	event = append(event, NewEvent(IntEvent{"foo", 40}).New())
	event = append(event, NewEvent(IntEvent{"foo", 60}).New())
	event = append(event, NewEvent(IntEvent{"foo", 50}).New())
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
		if result[tt.index].IntValue("Value") != tt.value {
			t.Error(result)
		}
	}
}

func TestSortFloat(t *testing.T) {
	v := SortFloat{"Value", false}

	event := []Event{}

	event = append(event, NewEvent(FloatEvent{"foo", 10.0}).New())
	event = append(event, NewEvent(FloatEvent{"foo", 30.0}).New())
	event = append(event, NewEvent(FloatEvent{"foo", 20.0}).New())
	event = append(event, NewEvent(FloatEvent{"foo", 40.0}).New())
	event = append(event, NewEvent(FloatEvent{"foo", 60.0}).New())
	event = append(event, NewEvent(FloatEvent{"foo", 50.0}).New())
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
		if result[tt.index].FloatValue("Value") != tt.value {
			t.Error(result)
		}
	}
}

func TestSortFloatReverse(t *testing.T) {
	v := SortFloat{"Value", true}

	event := []Event{}
	event = append(event, NewEvent(FloatEvent{"foo", 10.0}).New())
	event = append(event, NewEvent(FloatEvent{"foo", 30.0}).New())
	event = append(event, NewEvent(FloatEvent{"foo", 20.0}).New())
	event = append(event, NewEvent(FloatEvent{"foo", 40.0}).New())
	event = append(event, NewEvent(FloatEvent{"foo", 60.0}).New())
	event = append(event, NewEvent(FloatEvent{"foo", 50.0}).New())
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
		if result[tt.index].FloatValue("Value") != tt.value {
			t.Error(result)
		}
	}
}
