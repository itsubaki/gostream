package expr

import (
	"testing"

	"github.com/itsubaki/gostream/pkg/event"
)

func BenchmarkSumMapInt(b *testing.B) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	f := SumMapInt{"Record", "piyo", "sum(Record:piyo)"}

	events := event.List()
	for i := 0; i < 1; i++ {
		m := make(map[string]interface{})
		m["piyo"] = i
		events = append(events, event.New(MapEvent{m}))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.Apply(events)
	}
}

func BenchmarkSumMapInt128(b *testing.B) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	f := SumMapInt{"Record", "piyo", "sum(Record:piyo)"}

	events := event.List()
	for i := 0; i < 128; i++ {
		m := make(map[string]interface{})
		m["piyo"] = i
		events = append(events, event.New(MapEvent{m}))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.Apply(events)
	}
}

func BenchmarkAverageMapInt(b *testing.B) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	f := AverageMapInt{"Record", "piyo", "avg(Record:piyo)"}

	events := event.List()
	for i := 0; i < 1; i++ {
		m := make(map[string]interface{})
		m["piyo"] = i
		events = append(events, event.New(MapEvent{m}))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.Apply(events)
	}
}

func BenchmarkAverageMapInt128(b *testing.B) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	f := AverageMapInt{"Record", "piyo", "avg(Record:piyo)"}

	events := event.List()
	for i := 0; i < 128; i++ {
		m := make(map[string]interface{})
		m["piyo"] = i
		events = append(events, event.New(MapEvent{m}))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.Apply(events)
	}
}

func TestSelectMapAll(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	m := make(map[string]interface{})
	m["Name"] = "foo"

	f := SelectMapAll{"Record"}

	events := event.List(MapEvent{m})
	result := f.Apply(events)

	if result[0].RecordString("Name") != "foo" {
		t.Error(result)
	}

}

func TestSelectMapString(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	m := make(map[string]interface{})
	m["Name"] = "foo"

	f := SelectMapString{"Record", "Name", "Name"}

	events := event.List(MapEvent{m})
	result := f.Apply(events)

	if result[0].RecordString("Name") != "foo" {
		t.Error(result)
	}
}

func TestSelectMapBool(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	m := make(map[string]interface{})
	m["Name"] = false

	f := SelectMapBool{"Record", "Name", "Name"}

	events := event.List(MapEvent{m})
	result := f.Apply(events)

	if result[0].RecordBool("Name") {
		t.Error(result)
	}
}

func TestSelectMapInt(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	m := make(map[string]interface{})
	m["Name"] = 10

	f := SelectMapInt{"Record", "Name", "Name"}

	events := event.List(MapEvent{m})
	result := f.Apply(events)

	if result[0].RecordInt("Name") != 10 {
		t.Error(result)
	}
}

func TestSelectMapFloat(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	m := make(map[string]interface{})
	m["Name"] = 10.0

	f := SelectMapFloat{"Record", "Name", "Name"}

	events := event.List(MapEvent{m})
	result := f.Apply(events)

	if result[0].RecordFloat("Name") != 10.0 {
		t.Error(result)
	}
}

func TestSumMapInt(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	m := make(map[string]interface{})
	m["piyo"] = 123

	f := SumMapInt{"Record", "piyo", "sum(Record:piyo)"}

	events := event.List()

	events = append(events, event.New(MapEvent{m}))
	events = f.Apply(events)

	events = append(events, event.New(MapEvent{m}))
	events = f.Apply(events)

	var test = []struct {
		index int
		sum   int
	}{
		{0, 123},
		{1, 246},
	}

	for _, tt := range test {
		if events[tt.index].Record["sum(Record:piyo)"] != tt.sum {
			t.Error(events)
		}
	}
}

func TestSumMapFloat(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	m := make(map[string]interface{})
	m["piyo"] = 12.3

	f := SumMapFloat{"Record", "piyo", "sum(Record:piyo)"}

	events := event.List()

	events = append(events, event.New(MapEvent{m}))
	events = f.Apply(events)

	events = append(events, event.New(MapEvent{m}))
	events = f.Apply(events)

	var test = []struct {
		index int
		sum   float64
	}{
		{0, 12.3},
		{1, 24.6},
	}

	for _, tt := range test {
		if events[tt.index].Record["sum(Record:piyo)"] != tt.sum {
			t.Error(events)
		}
	}
}

func TestAverageMapInt(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	m := make(map[string]interface{})
	m["piyo"] = 15

	f := AverageMapInt{"Record", "piyo", "avg(Record:piyo)"}

	events := event.List()

	events = append(events, event.New(MapEvent{m}))
	events = f.Apply(events)

	events = append(events, event.New(MapEvent{m}))
	events = f.Apply(events)

	var test = []struct {
		index int
		avg   float64
	}{
		{0, 15.0},
		{1, 15.0},
	}

	for _, tt := range test {
		if events[tt.index].Record["avg(Record:piyo)"] != tt.avg {
			t.Error(events)
		}
	}
}

func TestAverageMapFloat(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	m := make(map[string]interface{})
	m["piyo"] = 15.0

	f := AverageMapFloat{"Record", "piyo", "avg(Record:piyo)"}

	events := event.List()

	events = append(events, event.New(MapEvent{m}))
	events = f.Apply(events)

	events = append(events, event.New(MapEvent{m}))
	events = f.Apply(events)

	var test = []struct {
		index int
		avg   float64
	}{
		{0, 15.0},
		{1, 15.0},
	}

	for _, tt := range test {
		if events[tt.index].Record["avg(Record:piyo)"] != tt.avg {
			t.Error(events)
		}
	}
}

func TestCastMapStringToInt(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	m := make(map[string]interface{})
	m["piyo"] = "123"

	events := event.List(MapEvent{m})
	cast := CastMapStringToInt{"Record", "piyo", "cast(Record:piyo)"}
	casted := cast.Apply(events)
	if casted[0].RecordInt("cast(Record:piyo)") != 123 {
		t.Error(casted)
	}

	events = event.List(MapEvent{casted[0].Record})
	sum := SumMapInt{"Record", "cast(Record:piyo)", "sum(Record:cast(Record:piyo))"}
	result := sum.Apply(events)
	if result[0].RecordInt("sum(Record:cast(Record:piyo))") != 123 {
		t.Error(result)
	}
}

func TestCastMapStringToFloat(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	m := make(map[string]interface{})
	m["piyo"] = "12.3"

	events := event.List(MapEvent{m})
	cast := CastMapStringToFloat{"Record", "piyo", "cast(Record:piyo)"}
	casted := cast.Apply(events)
	if casted[0].RecordFloat("cast(Record:piyo)") != 12.3 {
		t.Error(events)
	}
}

func TestCastMapStringToBool(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	m := make(map[string]interface{})
	m["piyo"] = "false"

	events := event.List(MapEvent{m})
	cast := CastMapStringToBool{"Record", "piyo", "cast(Record:piyo)"}
	casted := cast.Apply(events)
	if casted[0].RecordBool("cast(Record:piyo)") {
		t.Error(events)
	}
}
