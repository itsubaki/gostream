package gocep

import "testing"

func TestEqualsMapString(t *testing.T) {
	s := EqualsMapString{"Map", "foo", "bar"}

	m := make(map[string]interface{})
	m["foo"] = "bar"
	if !s.Select(NewEvent(MapEvent{"name", m})) {
		t.Error("failed.")
	}

	m["foo"] = "hoge"
	if s.Select(NewEvent(MapEvent{"name", m})) {
		t.Error("failed.")
	}
}

func TestNotEqualsMapString(t *testing.T) {
	s := NotEqualsMapString{"Map", "foo", "bar"}

	m := make(map[string]interface{})
	m["foo"] = "bar"
	if s.Select(NewEvent(MapEvent{"name", m})) {
		t.Error("failed.")
	}

	m["foo"] = "hoge"
	if !s.Select(NewEvent(MapEvent{"name", m})) {
		t.Error("failed.")
	}
}

func TestEqualsMapBool(t *testing.T) {
	s := EqualsMapBool{"Map", "foo", false}

	m := make(map[string]interface{})
	m["foo"] = false
	if !s.Select(NewEvent(MapEvent{"name", m})) {
		t.Error("failed.")
	}

	m["foo"] = true
	if s.Select(NewEvent(MapEvent{"name", m})) {
		t.Error("failed.")
	}
}

func TestNotEqualsMapBool(t *testing.T) {
	s := NotEqualsMapBool{"Map", "foo", false}

	m := make(map[string]interface{})
	m["foo"] = false
	if s.Select(NewEvent(MapEvent{"name", m})) {
		t.Error("failed.")
	}

	m["foo"] = true
	if !s.Select(NewEvent(MapEvent{"name", m})) {
		t.Error("failed.")
	}
}

func TestEqualsMapInt(t *testing.T) {
	s := EqualsMapInt{"Map", "foo", 123}

	m := make(map[string]interface{})
	m["foo"] = 123
	if !s.Select(NewEvent(MapEvent{"name", m})) {
		t.Error("failed.")
	}

	m["foo"] = 456
	if s.Select(NewEvent(MapEvent{"name", m})) {
		t.Error("failed.")
	}
}

func TestNotEqualsMapInt(t *testing.T) {
	s := NotEqualsMapInt{"Map", "foo", 123}

	m := make(map[string]interface{})
	m["foo"] = 123
	if s.Select(NewEvent(MapEvent{"name", m})) {
		t.Error("failed.")
	}

	m["foo"] = 456
	if !s.Select(NewEvent(MapEvent{"name", m})) {
		t.Error("failed.")
	}
}

func TestEqualsEqualsMapFloat(t *testing.T) {
	s := EqualsMapFloat{"Map", "foo", 12.3}

	m := make(map[string]interface{})
	m["foo"] = 12.3
	if !s.Select(NewEvent(MapEvent{"name", m})) {
		t.Error("failed.")
	}

	m["foo"] = 45.6
	if s.Select(NewEvent(MapEvent{"name", m})) {
		t.Error("failed.")
	}
}

func TestNotEqualsEqualsMapFloat(t *testing.T) {
	s := NotEqualsMapFloat{"Map", "foo", 12.3}

	m := make(map[string]interface{})
	m["foo"] = 12.3
	if s.Select(NewEvent(MapEvent{"name", m})) {
		t.Error("failed.")
	}

	m["foo"] = 45.6
	if !s.Select(NewEvent(MapEvent{"name", m})) {
		t.Error("failed.")
	}
}

func TestLargerThanMapInt(t *testing.T) {
	s := LargerThanMapInt{"Map", "foo", 100}

	m := make(map[string]interface{})
	m["foo"] = 100
	if s.Select(NewEvent(MapEvent{"name", m})) {
		t.Error("failed.")
	}

	m["foo"] = 101
	if !s.Select(NewEvent(MapEvent{"name", m})) {
		t.Error("failed.")
	}
}

func TestLargerThanMapFloat(t *testing.T) {
	s := LargerThanMapFloat{"Map", "foo", 10.0}

	m := make(map[string]interface{})
	m["foo"] = 10.0
	if s.Select(NewEvent(MapEvent{"name", m})) {
		t.Error("failed.")
	}

	m["foo"] = 10.1
	if !s.Select(NewEvent(MapEvent{"name", m})) {
		t.Error("failed.")
	}
}

func TestLessThanMapInt(t *testing.T) {
	s := LessThanMapInt{"Map", "foo", 100}

	m := make(map[string]interface{})
	m["foo"] = 101
	if s.Select(NewEvent(MapEvent{"name", m})) {
		t.Error("failed.")
	}

	m["foo"] = 99
	if !s.Select(NewEvent(MapEvent{"name", m})) {
		t.Error("failed.")
	}
}

func TestLessThanMapFloat(t *testing.T) {
	s := LessThanMapFloat{"Map", "foo", 10.0}

	m := make(map[string]interface{})
	m["foo"] = 10.1
	if s.Select(NewEvent(MapEvent{"name", m})) {
		t.Error("failed.")
	}

	m["foo"] = 9.9
	if !s.Select(NewEvent(MapEvent{"name", m})) {
		t.Error("failed.")
	}
}
