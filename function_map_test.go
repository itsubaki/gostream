package gocep

import "testing"

func TestSelectMapAll(t *testing.T) {
	m := make(map[string]interface{})
	m["Name"] = "foo"
	event := []Event{NewEvent(MapEvent{m})}
	f := SelectMapAll{"Record"}

	result := f.Apply(event)
	if result[0].RecordString("Name") != "foo" {
		t.Error(result)
	}

}

func TestSelectMapString(t *testing.T) {
	m := make(map[string]interface{})
	m["Name"] = "foo"
	event := []Event{NewEvent(MapEvent{m})}
	f := SelectMapString{"Record", "Name", "Name"}

	result := f.Apply(event)
	if result[0].RecordString("Name") != "foo" {
		t.Error(result)
	}
}

func TestSelectMapBool(t *testing.T) {
	m := make(map[string]interface{})
	m["Name"] = false
	event := []Event{NewEvent(MapEvent{m})}
	f := SelectMapBool{"Record", "Name", "Name"}

	result := f.Apply(event)
	if result[0].RecordBool("Name") {
		t.Error(result)
	}
}

func TestSelectMapInt(t *testing.T) {
	m := make(map[string]interface{})
	m["Name"] = 10
	event := []Event{NewEvent(MapEvent{m})}
	f := SelectMapInt{"Record", "Name", "Name"}

	result := f.Apply(event)
	if result[0].RecordInt("Name") != 10 {
		t.Error(result)
	}
}

func TestSelectMapFloat(t *testing.T) {
	m := make(map[string]interface{})
	m["Name"] = 10.0
	event := []Event{NewEvent(MapEvent{m})}
	f := SelectMapFloat{"Record", "Name", "Name"}

	result := f.Apply(event)
	if result[0].RecordFloat("Name") != 10.0 {
		t.Error(result)
	}
}

func TestSumMapInt(t *testing.T) {
	m := make(map[string]interface{})
	m["piyo"] = 123

	event := []Event{
		NewEvent(MapEvent{m}),
		NewEvent(MapEvent{m}),
	}
	f := SumMapInt{"Record", "piyo", "sum(Record:piyo)"}
	result := f.Apply(event)

	var test = []struct {
		index int
		sum   int
	}{
		{0, 246},
		{1, 246},
	}

	for _, tt := range test {
		if result[tt.index].Record["sum(Record:piyo)"] != tt.sum {
			t.Error(result)
		}
	}
}

func TestSumMapFloat(t *testing.T) {
	m := make(map[string]interface{})
	m["piyo"] = 12.3

	event := []Event{
		NewEvent(MapEvent{m}),
		NewEvent(MapEvent{m}),
	}
	f := SumMapFloat{"Record", "piyo", "sum(Record:piyo)"}
	result := f.Apply(event)

	var test = []struct {
		index int
		sum   float64
	}{
		{0, 24.6},
		{1, 24.6},
	}

	for _, tt := range test {
		if result[tt.index].Record["sum(Record:piyo)"] != tt.sum {
			t.Error(result)
		}
	}
}

func TestAverageMapInt(t *testing.T) {
	m := make(map[string]interface{})
	m["piyo"] = 15

	event := []Event{
		NewEvent(MapEvent{m}),
		NewEvent(MapEvent{m}),
	}
	f := AverageMapInt{"Record", "piyo", "avg(Record:piyo)"}
	result := f.Apply(event)

	var test = []struct {
		index int
		avg   float64
	}{
		{0, 15.0},
		{1, 15.0},
	}

	for _, tt := range test {
		if result[tt.index].Record["avg(Record:piyo)"] != tt.avg {
			t.Error(result)
		}
	}
}

func TestAverageMapFloat(t *testing.T) {
	m := make(map[string]interface{})
	m["piyo"] = 15.0

	event := []Event{
		NewEvent(MapEvent{m}),
		NewEvent(MapEvent{m}),
	}
	f := AverageMapFloat{"Record", "piyo", "avg(Record:piyo)"}
	result := f.Apply(event)

	var test = []struct {
		index int
		avg   float64
	}{
		{0, 15.0},
		{1, 15.0},
	}

	for _, tt := range test {
		if result[tt.index].Record["avg(Record:piyo)"] != tt.avg {
			t.Error(result)
		}
	}
}

func TestCastMapStringToInt(t *testing.T) {
	m := make(map[string]interface{})
	m["piyo"] = "123"

	event := []Event{
		NewEvent(MapEvent{m}),
		NewEvent(MapEvent{m}),
	}
	cast := CastMapStringToInt{"Record", "piyo", "cast(Record:piyo)"}
	casted := cast.Apply(event)
	if casted[0].RecordInt("cast(Record:piyo)") != 123 {
		t.Error(casted)
	}

	event = []Event{
		NewEvent(MapEvent{casted[0].Record}),
		NewEvent(MapEvent{casted[0].Record}),
	}

	sum := SumMapInt{"Record", "cast(Record:piyo)", "sum(Record:cast(Record:piyo))"}
	result := sum.Apply(event)
	if result[0].RecordInt("sum(Record:cast(Record:piyo))") != 246 {
		t.Error(result)
	}
}

func TestCastMapStringToFloat(t *testing.T) {
	m := make(map[string]interface{})
	m["piyo"] = "12.3"

	event := []Event{NewEvent(MapEvent{m})}
	cast := CastMapStringToFloat{"Record", "piyo", "cast(Record:piyo)"}
	casted := cast.Apply(event)
	if casted[0].RecordFloat("cast(Record:piyo)") != 12.3 {
		t.Error(event)
	}
}

func TestCastMapStringToBool(t *testing.T) {
	m := make(map[string]interface{})
	m["piyo"] = "false"

	event := []Event{NewEvent(MapEvent{m})}
	cast := CastMapStringToBool{"Record", "piyo", "cast(Record:piyo)"}
	casted := cast.Apply(event)
	if casted[0].RecordBool("cast(Record:piyo)") {
		t.Error(event)
	}
}
