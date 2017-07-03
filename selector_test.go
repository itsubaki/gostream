package gocep

import "testing"

func TestEqualsType(t *testing.T) {
	e0 := IntEvent{"foo", 1}
	e1 := IntEvent{"foo", 1}

	s := EqualsType{e0}
	if !s.Select(NewEvent(e1)) {
		t.Error("failed.")
	}
}

func TestNotEqualsType(t *testing.T) {
	e0 := IntEvent{"foo", 1}
	e1 := IntEvent{"foo", 1}

	s := NotEqualsType{e0}
	if s.Select(NewEvent(e1)) {
		t.Error("failed.")
	}
}

func TestEqualsString(t *testing.T) {
	s := EqualsString{"Name", "foo"}

	e0 := IntEvent{"foo", 1}
	if !s.Select(NewEvent(e0)) {
		t.Error("failed.")
	}

	e1 := IntEvent{"bar", 1}
	if s.Select(NewEvent(e1)) {
		t.Error("failed.")
	}
}

func TestEqualsBool(t *testing.T) {
	s := EqualsBool{"Value", true}

	e0 := BoolEvent{true}
	if !s.Select(NewEvent(e0)) {
		t.Error("failed.")
	}

	e1 := BoolEvent{false}
	if s.Select(NewEvent(e1)) {
		t.Error("failed.")
	}
}

func TestEqualsInt(t *testing.T) {
	s := EqualsInt{"Value", 1}

	e0 := IntEvent{"foo", 1}
	if !s.Select(NewEvent(e0)) {
		t.Error("failed.")
	}

	e1 := IntEvent{"foo", 2}
	if s.Select(NewEvent(e1)) {
		t.Error("failed.")
	}
}

func TestEqualsFloat(t *testing.T) {
	s := EqualsFloat{"Value", 1.0}

	e0 := FloatEvent{"foo", 1.0}
	if !s.Select(NewEvent(e0)) {
		t.Error("failed.")
	}

	e1 := FloatEvent{"foo", 2.0}
	if s.Select(NewEvent(e1)) {
		t.Error("failed.")
	}
}

func TestNotEqualsString(t *testing.T) {
	s := NotEqualsString{"Name", "foo"}

	e0 := IntEvent{"foo", 1}
	if s.Select(NewEvent(e0)) {
		t.Error("failed.")
	}

	e1 := IntEvent{"bar", 1}
	if !s.Select(NewEvent(e1)) {
		t.Error("failed.")
	}
}

func TestNotEqualsBool(t *testing.T) {
	s := NotEqualsBool{"Value", true}

	e0 := BoolEvent{true}
	if s.Select(NewEvent(e0)) {
		t.Error("failed.")
	}

	e1 := BoolEvent{false}
	if !s.Select(NewEvent(e1)) {
		t.Error("failed.")
	}
}

func TestNotEqualsInt(t *testing.T) {
	s := NotEqualsInt{"Value", 1}

	e0 := IntEvent{"foo", 1}
	if s.Select(NewEvent(e0)) {
		t.Error("failed.")
	}

	e1 := IntEvent{"foo", 2}
	if !s.Select(NewEvent(e1)) {
		t.Error("failed.")
	}
}

func TestNotEqualsFloat(t *testing.T) {
	s := NotEqualsFloat{"Value", 1.0}

	e0 := FloatEvent{"foo", 1.0}
	if s.Select(NewEvent(e0)) {
		t.Error("failed.")
	}

	e1 := FloatEvent{"foo", 2.0}
	if !s.Select(NewEvent(e1)) {
		t.Error("failed.")
	}
}

func TestLargerThanInt(t *testing.T) {
	s := LargerThanInt{"Value", 10}

	e0 := IntEvent{"foo", 10}
	if s.Select(NewEvent(e0)) {
		t.Error("failed.")
	}

	e1 := IntEvent{"bar", 11}
	if !s.Select(NewEvent(e1)) {
		t.Error("failed.")
	}
}

func TestLargerThanFloat(t *testing.T) {
	s := LargerThanFloat{"Value", 10.0}

	e0 := FloatEvent{"foo", 10.0}
	if s.Select(NewEvent(e0)) {
		t.Error("failed.")
	}

	e1 := FloatEvent{"bar", 10.1}
	if !s.Select(NewEvent(e1)) {
		t.Error("failed.")
	}
}

func TestLessThanInt(t *testing.T) {
	s := LessThanInt{"Value", 10}

	e0 := IntEvent{"foo", 10}
	if s.Select(NewEvent(e0)) {
		t.Error("failed.")
	}

	e1 := IntEvent{"bar", 9}
	if !s.Select(NewEvent(e1)) {
		t.Error("failed.")
	}
}

func TestLessThanFloat(t *testing.T) {
	s := LessThanFloat{"Value", 10.0}

	e0 := FloatEvent{"foo", 10.0}
	if s.Select(NewEvent(e0)) {
		t.Error("failed.")
	}

	e1 := FloatEvent{"bar", 9.9}
	if !s.Select(NewEvent(e1)) {
		t.Error("failed.")
	}
}
