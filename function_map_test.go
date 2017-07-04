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
