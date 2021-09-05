package function_test

import (
	"testing"

	"github.com/itsubaki/gostream/pkg/event"
	"github.com/itsubaki/gostream/pkg/function"
)

func TestEqualsMapString(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	s := function.EqualsMapString{"Record", "foo", "bar"}
	m := make(map[string]interface{})

	m["foo"] = "bar"
	if !s.Apply(event.New(MapEvent{m})) {
		t.Error("failed.")
	}

	m["foo"] = "hoge"
	if s.Apply(event.New(MapEvent{m})) {
		t.Error("failed.")
	}
}

func TestNotEqualsMapString(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	s := function.NotEqualsMapString{"Record", "foo", "bar"}
	m := make(map[string]interface{})

	m["foo"] = "bar"
	if s.Apply(event.New(MapEvent{m})) {
		t.Error("failed.")
	}

	m["foo"] = "hoge"
	if !s.Apply(event.New(MapEvent{m})) {
		t.Error("failed.")
	}
}

func TestEqualsMapBool(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	s := function.EqualsMapBool{"Record", "foo", false}
	m := make(map[string]interface{})

	m["foo"] = false
	if !s.Apply(event.New(MapEvent{m})) {
		t.Error("failed.")
	}

	m["foo"] = true
	if s.Apply(event.New(MapEvent{m})) {
		t.Error("failed.")
	}
}

func TestNotEqualsMapBool(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	s := function.NotEqualsMapBool{"Record", "foo", false}
	m := make(map[string]interface{})

	m["foo"] = false
	if s.Apply(event.New(MapEvent{m})) {
		t.Error("failed.")
	}

	m["foo"] = true
	if !s.Apply(event.New(MapEvent{m})) {
		t.Error("failed.")
	}
}

func TestEqualsMapInt(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	s := function.EqualsMapInt{"Record", "foo", 123}
	m := make(map[string]interface{})

	m["foo"] = 123
	if !s.Apply(event.New(MapEvent{m})) {
		t.Error("failed.")
	}

	m["foo"] = 456
	if s.Apply(event.New(MapEvent{m})) {
		t.Error("failed.")
	}
}

func TestNotEqualsMapInt(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	s := function.NotEqualsMapInt{"Record", "foo", 123}
	m := make(map[string]interface{})

	m["foo"] = 123
	if s.Apply(event.New(MapEvent{m})) {
		t.Error("failed.")
	}

	m["foo"] = 456
	if !s.Apply(event.New(MapEvent{m})) {
		t.Error("failed.")
	}
}

func TestEqualsEqualsMapFloat(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	s := function.EqualsMapFloat{"Record", "foo", 12.3}
	m := make(map[string]interface{})

	m["foo"] = 12.3
	if !s.Apply(event.New(MapEvent{m})) {
		t.Error("failed.")
	}

	m["foo"] = 45.6
	if s.Apply(event.New(MapEvent{m})) {
		t.Error("failed.")
	}
}

func TestNotEqualsEqualsMapFloat(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	s := function.NotEqualsMapFloat{"Record", "foo", 12.3}
	m := make(map[string]interface{})

	m["foo"] = 12.3
	if s.Apply(event.New(MapEvent{m})) {
		t.Error("failed.")
	}

	m["foo"] = 45.6
	if !s.Apply(event.New(MapEvent{m})) {
		t.Error("failed.")
	}
}

func TestLargerThanMapInt(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	s := function.LargerThanMapInt{"Record", "foo", 100}
	m := make(map[string]interface{})

	m["foo"] = 100
	if s.Apply(event.New(MapEvent{m})) {
		t.Error("failed.")
	}

	m["foo"] = 101
	if !s.Apply(event.New(MapEvent{m})) {
		t.Error("failed.")
	}
}

func TestLargerThanMapFloat(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	s := function.LargerThanMapFloat{"Record", "foo", 10.0}
	m := make(map[string]interface{})

	m["foo"] = 10.0
	if s.Apply(event.New(MapEvent{m})) {
		t.Error("failed.")
	}

	m["foo"] = 10.1
	if !s.Apply(event.New(MapEvent{m})) {
		t.Error("failed.")
	}
}

func TestLessThanMapInt(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	s := function.LessThanMapInt{"Record", "foo", 100}
	m := make(map[string]interface{})

	m["foo"] = 101
	if s.Apply(event.New(MapEvent{m})) {
		t.Error("failed.")
	}

	m["foo"] = 99
	if !s.Apply(event.New(MapEvent{m})) {
		t.Error("failed.")
	}
}

func TestLessThanMapFloat(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	s := function.LessThanMapFloat{"Record", "foo", 10.0}
	m := make(map[string]interface{})

	m["foo"] = 10.1
	if s.Apply(event.New(MapEvent{m})) {
		t.Error("failed.")
	}

	m["foo"] = 9.9
	if !s.Apply(event.New(MapEvent{m})) {
		t.Error("failed.")
	}
}
