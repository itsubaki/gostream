package gocep

import "reflect"

type Selector interface {
	Select(e Event) bool
}

type EqualsType struct {
	Accept interface{}
}

func (f EqualsType) Select(e Event) bool {
	return e.TypeEquals(reflect.TypeOf(f.Accept))
}

type NotEqualsType struct {
	Accept interface{}
}

func (f NotEqualsType) Select(e Event) bool {
	return !e.TypeEquals(reflect.TypeOf(f.Accept))
}

type EqualsString struct {
	Name  string
	Value string
}

func (f EqualsString) Select(e Event) bool {
	return e.StringValue(f.Name) == f.Value
}

type EqualsBool struct {
	Name  string
	Value bool
}

func (f EqualsBool) Select(e Event) bool {
	return e.BoolValue(f.Name) == f.Value
}

type EqualsInt struct {
	Name  string
	Value int
}

func (f EqualsInt) Select(e Event) bool {
	return e.IntValue(f.Name) == f.Value
}

type EqualsFloat struct {
	Name  string
	Value float32
}

func (f EqualsFloat) Select(e Event) bool {
	return e.Float32Value(f.Name) == f.Value
}

type NotEqualsString struct {
	Name  string
	Value string
}

func (f NotEqualsString) Select(e Event) bool {
	return e.StringValue(f.Name) != f.Value
}

type NotEqualsBool struct {
	Name  string
	Value bool
}

func (f NotEqualsBool) Select(e Event) bool {
	return e.BoolValue(f.Name) != f.Value
}

type NotEqualsInt struct {
	Name  string
	Value int
}

func (f NotEqualsInt) Select(e Event) bool {
	return e.IntValue(f.Name) != f.Value
}

type NotEqualsFloat struct {
	Name  string
	Value float32
}

func (f NotEqualsFloat) Select(e Event) bool {
	return e.Float32Value(f.Name) != f.Value
}

type LargerThanInt struct {
	Name  string
	Value int
}

func (f LargerThanInt) Select(e Event) bool {
	return e.IntValue(f.Name) > f.Value
}

type LargerThanFloat struct {
	Name  string
	Value float32
}

func (f LargerThanFloat) Select(e Event) bool {
	return e.Float32Value(f.Name) > f.Value
}

type LessThanInt struct {
	Name  string
	Value int
}

func (f LessThanInt) Select(e Event) bool {
	return e.IntValue(f.Name) < f.Value
}

type LessThanFloat struct {
	Name  string
	Value float32
}

func (f LessThanFloat) Select(e Event) bool {
	return e.Float32Value(f.Name) < f.Value
}
