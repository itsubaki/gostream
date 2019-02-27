package gocep

import "reflect"

type Or []Selector

func (f Or) Select(e Event) bool {
	for _, s := range f {
		if s.Select(e) {
			return true
		}
	}
	return false
}

type And []Selector

func (f And) Select(e Event) bool {
	for _, s := range f {
		if !s.Select(e) {
			return false
		}
	}
	return true
}

type Selector interface {
	Select(e Event) bool
}

type EqualsType struct {
	Accept interface{}
}

func (f EqualsType) Select(e Event) bool {
	return e.EqualsType(reflect.TypeOf(f.Accept))
}

type NotEqualsType struct {
	Accept interface{}
}

func (f NotEqualsType) Select(e Event) bool {
	return !e.EqualsType(reflect.TypeOf(f.Accept))
}

type EqualsString struct {
	Name  string
	Value string
}

func (f EqualsString) Select(e Event) bool {
	return e.String(f.Name) == f.Value
}

type EqualsBool struct {
	Name  string
	Value bool
}

func (f EqualsBool) Select(e Event) bool {
	return e.Bool(f.Name) == f.Value
}

type EqualsInt struct {
	Name  string
	Value int
}

func (f EqualsInt) Select(e Event) bool {
	return e.Int(f.Name) == f.Value
}

type EqualsFloat struct {
	Name  string
	Value float64
}

func (f EqualsFloat) Select(e Event) bool {
	return e.Float(f.Name) == f.Value
}

type NotEqualsString struct {
	Name  string
	Value string
}

func (f NotEqualsString) Select(e Event) bool {
	return e.String(f.Name) != f.Value
}

type NotEqualsBool struct {
	Name  string
	Value bool
}

func (f NotEqualsBool) Select(e Event) bool {
	return e.Bool(f.Name) != f.Value
}

type NotEqualsInt struct {
	Name  string
	Value int
}

func (f NotEqualsInt) Select(e Event) bool {
	return e.Int(f.Name) != f.Value
}

type NotEqualsFloat struct {
	Name  string
	Value float64
}

func (f NotEqualsFloat) Select(e Event) bool {
	return e.Float(f.Name) != f.Value
}

type LargerThanInt struct {
	Name  string
	Value int
}

func (f LargerThanInt) Select(e Event) bool {
	return e.Int(f.Name) > f.Value
}

type LargerThanFloat struct {
	Name  string
	Value float64
}

func (f LargerThanFloat) Select(e Event) bool {
	return e.Float(f.Name) > f.Value
}

type LessThanInt struct {
	Name  string
	Value int
}

func (f LessThanInt) Select(e Event) bool {
	return e.Int(f.Name) < f.Value
}

type LessThanFloat struct {
	Name  string
	Value float64
}

func (f LessThanFloat) Select(e Event) bool {
	return e.Float(f.Name) < f.Value
}
