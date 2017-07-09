package gocep

import "testing"

func TestSortMapInt(t *testing.T) {
	v := OrderByMapInt{"Record", "foo", false}

	event := []Event{}

	for i := 10; i < 70; i = i + 10 {
		m := make(map[string]interface{})
		m["foo"] = i
		event = append(event, NewEvent(MapEvent{m}))
	}
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
		if result[tt.index].MapInt("Record", "foo") != tt.value {
			t.Error(result)
		}
	}
}

func TestSortMapIntReverse(t *testing.T) {
	v := OrderByMapInt{"Record", "foo", true}

	event := []Event{}

	for i := 10; i < 70; i = i + 10 {
		m := make(map[string]interface{})
		m["foo"] = i
		event = append(event, NewEvent(MapEvent{m}))
	}
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
		if result[tt.index].MapInt("Record", "foo") != tt.value {
			t.Error(result)
		}
	}
}

func TestSortMapFloat(t *testing.T) {
	v := OrderByMapFloat{"Record", "foo", false}

	event := []Event{}

	for i := 10; i < 70; i = i + 10 {
		m := make(map[string]interface{})
		m["foo"] = float64(i)
		event = append(event, NewEvent(MapEvent{m}))
	}
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
		if result[tt.index].MapFloat("Record", "foo") != tt.value {
			t.Error(result)
		}
	}
}

func TestSortMapFloatReverse(t *testing.T) {
	v := OrderByMapFloat{"Record", "foo", true}

	event := []Event{}

	for i := 10; i < 70; i = i + 10 {
		m := make(map[string]interface{})
		m["foo"] = float64(i)
		event = append(event, NewEvent(MapEvent{m}))
	}
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
		if result[tt.index].MapFloat("Record", "foo") != tt.value {
			t.Error(result)
		}
	}
}
