package clause_test

import (
	"testing"

	"github.com/itsubaki/gostream/pkg/clause"
	"github.com/itsubaki/gostream/pkg/event"
)

func BenchmarkSumMapInt(b *testing.B) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	f := clause.SumMapInt{"Record", "piyo", "sum(Record:piyo)"}

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

	f := clause.SumMapInt{"Record", "piyo", "sum(Record:piyo)"}

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

	f := clause.AverageMapInt{"Record", "piyo", "avg(Record:piyo)"}

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

	f := clause.AverageMapInt{"Record", "piyo", "avg(Record:piyo)"}

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

	f := clause.SelectMapAll{"Record"}

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

	f := clause.SelectMapString{"Record", "Name", "Name"}

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

	f := clause.SelectMapBool{"Record", "Name", "Name"}

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

	f := clause.SelectMapInt{"Record", "Name", "Name"}

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

	f := clause.SelectMapFloat{"Record", "Name", "Name"}

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

	f := clause.SumMapInt{"Record", "piyo", "sum(Record:piyo)"}

	events := event.List()

	events = append(events, event.New(MapEvent{m}))
	events = f.Apply(events)

	events = append(events, event.New(MapEvent{m}))
	events = f.Apply(events)

	var cases = []struct {
		in   int
		want int
	}{
		{0, 123},
		{1, 246},
	}

	for _, c := range cases {
		if events[c.in].Record["sum(Record:piyo)"] != c.want {
			t.Fail()
		}
	}
}

func TestSumMapFloat(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	m := make(map[string]interface{})
	m["piyo"] = 12.3

	f := clause.SumMapFloat{"Record", "piyo", "sum(Record:piyo)"}

	events := event.List()

	events = append(events, event.New(MapEvent{m}))
	events = f.Apply(events)

	events = append(events, event.New(MapEvent{m}))
	events = f.Apply(events)

	var cases = []struct {
		in   int
		want float64
	}{
		{0, 12.3},
		{1, 24.6},
	}

	for _, c := range cases {
		if events[c.in].Record["sum(Record:piyo)"] != c.want {
			t.Fail()
		}
	}
}

func TestAverageMapInt(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	m := make(map[string]interface{})
	m["piyo"] = 15

	f := clause.AverageMapInt{"Record", "piyo", "avg(Record:piyo)"}

	events := event.List()

	events = append(events, event.New(MapEvent{m}))
	events = f.Apply(events)

	events = append(events, event.New(MapEvent{m}))
	events = f.Apply(events)

	var cases = []struct {
		in   int
		want float64
	}{
		{0, 15.0},
		{1, 15.0},
	}

	for _, c := range cases {
		if events[c.in].Record["avg(Record:piyo)"] != c.want {
			t.Fail()
		}
	}
}

func TestAverageMapFloat(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	m := make(map[string]interface{})
	m["piyo"] = 15.0

	f := clause.AverageMapFloat{"Record", "piyo", "avg(Record:piyo)"}

	events := event.List()

	events = append(events, event.New(MapEvent{m}))
	events = f.Apply(events)

	events = append(events, event.New(MapEvent{m}))
	events = f.Apply(events)

	var cases = []struct {
		in   int
		want float64
	}{
		{0, 15.0},
		{1, 15.0},
	}

	for _, c := range cases {
		if events[c.in].Record["avg(Record:piyo)"] != c.want {
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
	cast := clause.CastMapStringToInt{"Record", "piyo", "cast(Record:piyo)"}
	casted := cast.Apply(events)
	if casted[0].RecordInt("cast(Record:piyo)") != 123 {
		t.Error(casted)
	}

	events = event.List(MapEvent{casted[0].Record})
	sum := clause.SumMapInt{"Record", "cast(Record:piyo)", "sum(Record:cast(Record:piyo))"}
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
	cast := clause.CastMapStringToFloat{"Record", "piyo", "cast(Record:piyo)"}
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
	cast := clause.CastMapStringToBool{"Record", "piyo", "cast(Record:piyo)"}
	casted := cast.Apply(events)
	if casted[0].RecordBool("cast(Record:piyo)") {
		t.Error(events)
	}
}
