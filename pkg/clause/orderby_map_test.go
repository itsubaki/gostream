package clause_test

import (
	"testing"

	"github.com/itsubaki/gostream/pkg/clause"
	"github.com/itsubaki/gostream/pkg/event"
)

func TestOrderByMapInt(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	events := event.List()
	for i := 10; i < 70; i = i + 10 {
		m := make(map[string]interface{})
		m["foo"] = i
		events = append(events, event.New(MapEvent{m}))
	}

	v := clause.OrderByMapInt{"Record", "foo", false}
	result := v.Apply(events)

	var cases = []struct {
		in   int
		want int
	}{
		{0, 10},
		{1, 20},
		{2, 30},
		{3, 40},
		{4, 50},
		{5, 60},
	}

	for _, c := range cases {
		if result[c.in].MapInt("Record", "foo") != c.want {
			t.Fail()
		}
	}
}

func TestOrderByMapIntReverse(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	events := event.List()

	for i := 10; i < 70; i = i + 10 {
		m := make(map[string]interface{})
		m["foo"] = i
		events = append(events, event.New(MapEvent{m}))
	}

	v := clause.OrderByMapInt{"Record", "foo", true}
	result := v.Apply(events)

	var cases = []struct {
		in   int
		want int
	}{
		{0, 60},
		{1, 50},
		{2, 40},
		{3, 30},
		{4, 20},
		{5, 10},
	}

	for _, c := range cases {
		if result[c.in].MapInt("Record", "foo") != c.want {
			t.Fail()
		}
	}
}

func TestOrderByMapFloat(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	events := event.List()
	for i := 10; i < 70; i = i + 10 {
		m := make(map[string]interface{})
		m["foo"] = float64(i)
		events = append(events, event.New(MapEvent{m}))
	}

	v := clause.OrderByMapFloat{"Record", "foo", false}
	result := v.Apply(events)

	var cases = []struct {
		in   int
		want float64
	}{
		{0, 10.0},
		{1, 20.0},
		{2, 30.0},
		{3, 40.0},
		{4, 50.0},
		{5, 60.0},
	}

	for _, c := range cases {
		if result[c.in].MapFloat("Record", "foo") != c.want {
			t.Fail()
		}
	}
}

func TestOrderByMapFloatReverse(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	events := event.List()
	for i := 10; i < 70; i = i + 10 {
		m := make(map[string]interface{})
		m["foo"] = float64(i)
		events = append(events, event.New(MapEvent{m}))
	}

	v := clause.OrderByMapFloat{"Record", "foo", true}
	result := v.Apply(events)

	var cases = []struct {
		in   int
		want float64
	}{
		{0, 60.0},
		{1, 50.0},
		{2, 40.0},
		{3, 30.0},
		{4, 20.0},
		{5, 10.0},
	}

	for _, c := range cases {
		if result[c.in].MapFloat("Record", "foo") != c.want {
			t.Fail()
		}
	}
}
