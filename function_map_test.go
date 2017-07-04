package gocep

import "testing"

func TestSumMapInt(t *testing.T) {
	m := make(map[string]interface{})
	m["piyo"] = 123

	event := []Event{}
	event = append(event, NewEvent(MapEvent{"foobar", m}).New())
	event = append(event, NewEvent(MapEvent{"foobar", m}).New())

	f := SumMapInt{"Map", "piyo"}
	result := f.Apply(event)

	var test = []struct {
		index int
		sum   int
	}{
		{0, 246},
		{1, 246},
	}

	for _, tt := range test {
		if result[tt.index].Record["sum(Map:piyo)"] != tt.sum {
			t.Error(result)
		}
	}
}

func TestSumMapFloat(t *testing.T) {
	m := make(map[string]interface{})
	m["piyo"] = 12.3

	event := []Event{}
	event = append(event, NewEvent(MapEvent{"foobar", m}).New())
	event = append(event, NewEvent(MapEvent{"foobar", m}).New())

	f := SumMapFloat{"Map", "piyo"}
	result := f.Apply(event)

	var test = []struct {
		index int
		sum   float64
	}{
		{0, 24.6},
		{1, 24.6},
	}

	for _, tt := range test {
		if result[tt.index].Record["sum(Map:piyo)"] != tt.sum {
			t.Error(result)
		}
	}
}

func TestAverageMapInt(t *testing.T) {
	m := make(map[string]interface{})
	m["piyo"] = 15

	event := []Event{}
	event = append(event, NewEvent(MapEvent{"foobar", m}).New())
	event = append(event, NewEvent(MapEvent{"foobar", m}).New())

	f := AverageMapInt{"Map", "piyo"}
	result := f.Apply(event)

	var test = []struct {
		index int
		avg   float64
	}{
		{0, 15.0},
		{1, 15.0},
	}

	for _, tt := range test {
		if result[tt.index].Record["avg(Map:piyo)"] != tt.avg {
			t.Error(result)
		}
	}
}

func TestAverageMapFloat(t *testing.T) {
	m := make(map[string]interface{})
	m["piyo"] = 15.0

	event := []Event{}
	event = append(event, NewEvent(MapEvent{"foobar", m}).New())
	event = append(event, NewEvent(MapEvent{"foobar", m}).New())

	f := AverageMapFloat{"Map", "piyo"}
	result := f.Apply(event)

	var test = []struct {
		index int
		avg   float64
	}{
		{0, 15.0},
		{1, 15.0},
	}

	for _, tt := range test {
		if result[tt.index].Record["avg(Map:piyo)"] != tt.avg {
			t.Error(result)
		}
	}
}
