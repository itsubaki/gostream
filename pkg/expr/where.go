package expr

import (
	"reflect"

	"github.com/itsubaki/gostream/pkg/event"
)

type Where interface {
	Apply(e event.Event) bool
}

type Or []Where

func (f Or) Apply(e event.Event) bool {
	for _, s := range f {
		if s.Apply(e) {
			return true
		}
	}
	return false
}

type And []Where

func (f And) Apply(e event.Event) bool {
	for _, s := range f {
		if !s.Apply(e) {
			return false
		}
	}
	return true
}

type EqualsType struct {
	Accept interface{}
}

func (f EqualsType) Apply(e event.Event) bool {
	return e.EqualsType(reflect.TypeOf(f.Accept))
}

type NotEqualsType struct {
	Accept interface{}
}

func (f NotEqualsType) Apply(e event.Event) bool {
	return !e.EqualsType(reflect.TypeOf(f.Accept))
}

type EqualsString struct {
	Name  string
	Value string
}

func (f EqualsString) Apply(e event.Event) bool {
	return e.String(f.Name) == f.Value
}

type EqualsBool struct {
	Name  string
	Value bool
}

func (f EqualsBool) Apply(e event.Event) bool {
	return e.Bool(f.Name) == f.Value
}

type EqualsInt struct {
	Name  string
	Value int
}

func (f EqualsInt) Apply(e event.Event) bool {
	return e.Int(f.Name) == f.Value
}

type EqualsFloat struct {
	Name  string
	Value float64
}

func (f EqualsFloat) Apply(e event.Event) bool {
	return e.Float(f.Name) == f.Value
}

type NotEqualsString struct {
	Name  string
	Value string
}

func (f NotEqualsString) Apply(e event.Event) bool {
	return e.String(f.Name) != f.Value
}

type NotEqualsBool struct {
	Name  string
	Value bool
}

func (f NotEqualsBool) Apply(e event.Event) bool {
	return e.Bool(f.Name) != f.Value
}

type NotEqualsInt struct {
	Name  string
	Value int
}

func (f NotEqualsInt) Apply(e event.Event) bool {
	return e.Int(f.Name) != f.Value
}

type NotEqualsFloat struct {
	Name  string
	Value float64
}

func (f NotEqualsFloat) Apply(e event.Event) bool {
	return e.Float(f.Name) != f.Value
}

type LargerThanInt struct {
	Name  string
	Value int
}

func (f LargerThanInt) Apply(e event.Event) bool {
	return e.Int(f.Name) > f.Value
}

type LargerThanFloat struct {
	Name  string
	Value float64
}

func (f LargerThanFloat) Apply(e event.Event) bool {
	return e.Float(f.Name) > f.Value
}

type LessThanInt struct {
	Name  string
	Value int
}

func (f LessThanInt) Apply(e event.Event) bool {
	return e.Int(f.Name) < f.Value
}

type LessThanFloat struct {
	Name  string
	Value float64
}

func (f LessThanFloat) Apply(e event.Event) bool {
	return e.Float(f.Name) < f.Value
}
