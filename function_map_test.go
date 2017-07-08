package gocep

import "testing"

func TestSumMapInt(t *testing.T) {
	m := make(map[string]interface{})
	m["piyo"] = 123

	event := []Event{}
	event = append(event, NewEvent(MapEvent{"foobar", m}))
	event = append(event, NewEvent(MapEvent{"foobar", m}))

	f := SumMapInt{"Map", "piyo", "sum(Map:piyo)"}
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
	event = append(event, NewEvent(MapEvent{"foobar", m}))
	event = append(event, NewEvent(MapEvent{"foobar", m}))

	f := SumMapFloat{"Map", "piyo", "sum(Map:piyo)"}
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
	event = append(event, NewEvent(MapEvent{"foobar", m}))
	event = append(event, NewEvent(MapEvent{"foobar", m}))

	f := AverageMapInt{"Map", "piyo", "avg(Map:piyo)"}
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
	event = append(event, NewEvent(MapEvent{"foobar", m}))
	event = append(event, NewEvent(MapEvent{"foobar", m}))

	f := AverageMapFloat{"Map", "piyo", "avg(Map:piyo)"}
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

func TestCastMapStringToInt(t *testing.T) {
	m := make(map[string]interface{})
	m["piyo"] = "123"

	event := []Event{}
	event = append(event, NewEvent(MapEvent{"foobar", m}))

	cast := CastMapStringToInt{"Map", "piyo", "cast(Map:piyo)"}
	casted := cast.Apply(event)
	if casted[0].RecordInt("cast(Map:piyo)") != 123 {
		t.Error(casted)
	}

	event = []Event{}
	event = append(event, NewEvent(MapEvent{"new", casted[0].Record}))
	event = append(event, NewEvent(MapEvent{"new", casted[0].Record}))

	sum := SumMapInt{"Map", "cast(Map:piyo)", "sum(Map:cast(Map:piyo))"}
	result := sum.Apply(event)
	if result[0].RecordInt("sum(Map:cast(Map:piyo))") != 246 {
		t.Error(result)
	}
}

func TestCastMapStringToFloat(t *testing.T) {
	m := make(map[string]interface{})
	m["piyo"] = "12.3"

	event := []Event{}
	event = append(event, NewEvent(MapEvent{"foobar", m}))

	cast := CastMapStringToFloat{"Map", "piyo", "cast(Map:piyo)"}
	casted := cast.Apply(event)
	if casted[0].RecordFloat("cast(Map:piyo)") != 12.3 {
		t.Error(event)
	}
}

func TestCastMapStringToBool(t *testing.T) {
	m := make(map[string]interface{})
	m["piyo"] = "false"

	event := []Event{}
	event = append(event, NewEvent(MapEvent{"foobar", m}))

	cast := CastMapStringToBool{"Map", "piyo", "cast(Map:piyo)"}
	casted := cast.Apply(event)
	if casted[0].RecordBool("cast(Map:piyo)") {
		t.Error(event)
	}
}
