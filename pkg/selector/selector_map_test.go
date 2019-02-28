package selector

import (
	"testing"

	"github.com/itsubaki/gocep/pkg/event"
)

func TestEqualsMapString(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	s := EqualsMapString{"Record", "foo", "bar"}
	m := make(map[string]interface{})

	m["foo"] = "bar"
	if !s.Select(event.New(MapEvent{m})) {
		t.Error("failed.")
	}

	m["foo"] = "hoge"
	if s.Select(event.New(MapEvent{m})) {
		t.Error("failed.")
	}
}

func TestNotEqualsMapString(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	s := NotEqualsMapString{"Record", "foo", "bar"}
	m := make(map[string]interface{})

	m["foo"] = "bar"
	if s.Select(event.New(MapEvent{m})) {
		t.Error("failed.")
	}

	m["foo"] = "hoge"
	if !s.Select(event.New(MapEvent{m})) {
		t.Error("failed.")
	}
}

func TestEqualsMapBool(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	s := EqualsMapBool{"Record", "foo", false}
	m := make(map[string]interface{})

	m["foo"] = false
	if !s.Select(event.New(MapEvent{m})) {
		t.Error("failed.")
	}

	m["foo"] = true
	if s.Select(event.New(MapEvent{m})) {
		t.Error("failed.")
	}
}

func TestNotEqualsMapBool(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	s := NotEqualsMapBool{"Record", "foo", false}
	m := make(map[string]interface{})

	m["foo"] = false
	if s.Select(event.New(MapEvent{m})) {
		t.Error("failed.")
	}

	m["foo"] = true
	if !s.Select(event.New(MapEvent{m})) {
		t.Error("failed.")
	}
}

func TestEqualsMapInt(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	s := EqualsMapInt{"Record", "foo", 123}
	m := make(map[string]interface{})

	m["foo"] = 123
	if !s.Select(event.New(MapEvent{m})) {
		t.Error("failed.")
	}

	m["foo"] = 456
	if s.Select(event.New(MapEvent{m})) {
		t.Error("failed.")
	}
}

func TestNotEqualsMapInt(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	s := NotEqualsMapInt{"Record", "foo", 123}
	m := make(map[string]interface{})

	m["foo"] = 123
	if s.Select(event.New(MapEvent{m})) {
		t.Error("failed.")
	}

	m["foo"] = 456
	if !s.Select(event.New(MapEvent{m})) {
		t.Error("failed.")
	}
}

func TestEqualsEqualsMapFloat(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	s := EqualsMapFloat{"Record", "foo", 12.3}
	m := make(map[string]interface{})

	m["foo"] = 12.3
	if !s.Select(event.New(MapEvent{m})) {
		t.Error("failed.")
	}

	m["foo"] = 45.6
	if s.Select(event.New(MapEvent{m})) {
		t.Error("failed.")
	}
}

func TestNotEqualsEqualsMapFloat(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	s := NotEqualsMapFloat{"Record", "foo", 12.3}
	m := make(map[string]interface{})

	m["foo"] = 12.3
	if s.Select(event.New(MapEvent{m})) {
		t.Error("failed.")
	}

	m["foo"] = 45.6
	if !s.Select(event.New(MapEvent{m})) {
		t.Error("failed.")
	}
}

func TestLargerThanMapInt(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	s := LargerThanMapInt{"Record", "foo", 100}
	m := make(map[string]interface{})

	m["foo"] = 100
	if s.Select(event.New(MapEvent{m})) {
		t.Error("failed.")
	}

	m["foo"] = 101
	if !s.Select(event.New(MapEvent{m})) {
		t.Error("failed.")
	}
}

func TestLargerThanMapFloat(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	s := LargerThanMapFloat{"Record", "foo", 10.0}
	m := make(map[string]interface{})

	m["foo"] = 10.0
	if s.Select(event.New(MapEvent{m})) {
		t.Error("failed.")
	}

	m["foo"] = 10.1
	if !s.Select(event.New(MapEvent{m})) {
		t.Error("failed.")
	}
}

func TestLessThanMapInt(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	s := LessThanMapInt{"Record", "foo", 100}
	m := make(map[string]interface{})

	m["foo"] = 101
	if s.Select(event.New(MapEvent{m})) {
		t.Error("failed.")
	}

	m["foo"] = 99
	if !s.Select(event.New(MapEvent{m})) {
		t.Error("failed.")
	}
}

func TestLessThanMapFloat(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	s := LessThanMapFloat{"Record", "foo", 10.0}
	m := make(map[string]interface{})

	m["foo"] = 10.1
	if s.Select(event.New(MapEvent{m})) {
		t.Error("failed.")
	}

	m["foo"] = 9.9
	if !s.Select(event.New(MapEvent{m})) {
		t.Error("failed.")
	}
}
