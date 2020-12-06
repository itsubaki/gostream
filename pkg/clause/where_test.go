package clause_test

import (
	"testing"

	"github.com/itsubaki/gostream/pkg/clause"
	"github.com/itsubaki/gostream/pkg/event"
)

func BenchmarkEqualsType(b *testing.B) {
	type IntEvent struct {
		Name  string
		Value int
	}

	e0 := IntEvent{"foo", 1}
	e1 := IntEvent{"foo", 1}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		clause.EqualsType{e0}.Apply(event.New(e1))
	}
}

func TestEqualsType(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	e0 := IntEvent{"foo", 1}
	e1 := IntEvent{"foo", 1}

	s := clause.EqualsType{e0}
	if !s.Apply(event.New(e1)) {
		t.Error("failed.")
	}
}

func TestNotEqualsType(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	e0 := IntEvent{"foo", 1}
	e1 := IntEvent{"foo", 1}

	s := clause.NotEqualsType{e0}
	if s.Apply(event.New(e1)) {
		t.Error("failed.")
	}
}

func TestEqualsString(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	s := clause.EqualsString{"Name", "foo"}

	e0 := IntEvent{"foo", 1}
	if !s.Apply(event.New(e0)) {
		t.Error("failed.")
	}

	e1 := IntEvent{"bar", 1}
	if s.Apply(event.New(e1)) {
		t.Error("failed.")
	}
}

func TestEqualsBool(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	type BoolEvent struct {
		Value bool
	}

	s := clause.EqualsBool{"Value", true}

	e0 := BoolEvent{true}
	if !s.Apply(event.New(e0)) {
		t.Error("failed.")
	}

	e1 := BoolEvent{false}
	if s.Apply(event.New(e1)) {
		t.Error("failed.")
	}
}

func TestEqualsInt(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	s := clause.EqualsInt{"Value", 1}

	e0 := IntEvent{"foo", 1}
	if !s.Apply(event.New(e0)) {
		t.Error("failed.")
	}

	e1 := IntEvent{"foo", 2}
	if s.Apply(event.New(e1)) {
		t.Error("failed.")
	}
}

func TestEqualsFloat(t *testing.T) {
	type FloatEvent struct {
		Name  string
		Value float64
	}

	s := clause.EqualsFloat{"Value", 1.0}

	e0 := FloatEvent{"foo", 1.0}
	if !s.Apply(event.New(e0)) {
		t.Error("failed.")
	}

	e1 := FloatEvent{"foo", 2.0}
	if s.Apply(event.New(e1)) {
		t.Error("failed.")
	}
}

func TestNotEqualsString(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	s := clause.NotEqualsString{"Name", "foo"}

	e0 := IntEvent{"foo", 1}
	if s.Apply(event.New(e0)) {
		t.Error("failed.")
	}

	e1 := IntEvent{"bar", 1}
	if !s.Apply(event.New(e1)) {
		t.Error("failed.")
	}
}

func TestNotEqualsBool(t *testing.T) {
	type BoolEvent struct {
		Value bool
	}

	s := clause.NotEqualsBool{"Value", true}

	e0 := BoolEvent{true}
	if s.Apply(event.New(e0)) {
		t.Error("failed.")
	}

	e1 := BoolEvent{false}
	if !s.Apply(event.New(e1)) {
		t.Error("failed.")
	}
}

func TestNotEqualsInt(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	s := clause.NotEqualsInt{"Value", 1}

	e0 := IntEvent{"foo", 1}
	if s.Apply(event.New(e0)) {
		t.Error("failed.")
	}

	e1 := IntEvent{"foo", 2}
	if !s.Apply(event.New(e1)) {
		t.Error("failed.")
	}
}

func TestNotEqualsFloat(t *testing.T) {
	type FloatEvent struct {
		Name  string
		Value float64
	}

	s := clause.NotEqualsFloat{"Value", 1.0}

	e0 := FloatEvent{"foo", 1.0}
	if s.Apply(event.New(e0)) {
		t.Error("failed.")
	}

	e1 := FloatEvent{"foo", 2.0}
	if !s.Apply(event.New(e1)) {
		t.Error("failed.")
	}
}

func TestLargerThanInt(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	s := clause.LargerThanInt{"Value", 10}

	e0 := IntEvent{"foo", 10}
	if s.Apply(event.New(e0)) {
		t.Error("failed.")
	}

	e1 := IntEvent{"bar", 11}
	if !s.Apply(event.New(e1)) {
		t.Error("failed.")
	}
}

func TestLargerThanFloat(t *testing.T) {
	type FloatEvent struct {
		Name  string
		Value float64
	}

	s := clause.LargerThanFloat{"Value", 10.0}

	e0 := FloatEvent{"foo", 10.0}
	if s.Apply(event.New(e0)) {
		t.Error("failed.")
	}

	e1 := FloatEvent{"bar", 10.1}
	if !s.Apply(event.New(e1)) {
		t.Error("failed.")
	}
}

func TestLessThanInt(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	s := clause.LessThanInt{"Value", 10}

	e0 := IntEvent{"foo", 10}
	if s.Apply(event.New(e0)) {
		t.Error("failed.")
	}

	e1 := IntEvent{"bar", 9}
	if !s.Apply(event.New(e1)) {
		t.Error("failed.")
	}
}

func TestLessThanFloat(t *testing.T) {
	type FloatEvent struct {
		Name  string
		Value float64
	}

	s := clause.LessThanFloat{"Value", 10.0}

	e0 := FloatEvent{"foo", 10.0}
	if s.Apply(event.New(e0)) {
		t.Error("failed.")
	}

	e1 := FloatEvent{"bar", 9.9}
	if !s.Apply(event.New(e1)) {
		t.Error("failed.")
	}
}
