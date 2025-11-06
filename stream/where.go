package stream

import (
	"fmt"
	"reflect"
)

var (
	_ Where = (*From)(nil)
	_ Where = (*LargerThan)(nil)
	_ Where = (*LessThan)(nil)
	_ Where = (*Equal)(nil)
	_ Where = (*NotEqual)(nil)
	_ Where = (*And)(nil)
)

type Where interface {
	Apply(input any) bool
	String() string
}

type From struct {
	Type any
}

func (w From) Apply(input any) bool {
	return reflect.TypeOf(input) == reflect.TypeOf(w.Type)
}

func (w From) String() string {
	return reflect.TypeOf(w.Type).Name()
}

type LargerThan struct {
	Name  string
	Value any
}

func (w LargerThan) Apply(input any) bool {
	v := reflect.ValueOf(input).FieldByName(w.Name).Interface()

	switch val := w.Value.(type) {
	case int:
		return v.(int) > val
	case int32:
		return v.(int32) > val
	case int64:
		return v.(int64) > val
	case float32:
		return v.(float32) > val
	case float64:
		return v.(float64) > val
	}

	return true
}

func (w LargerThan) String() string {
	return fmt.Sprintf("%v > %v", w.Name, w.Value)
}

type LessThan struct {
	Name  string
	Value any
}

func (w LessThan) Apply(input any) bool {
	v := reflect.ValueOf(input).FieldByName(w.Name).Interface()

	switch val := w.Value.(type) {
	case int:
		return v.(int) < val
	case int32:
		return v.(int32) < val
	case int64:
		return v.(int64) < val
	case float32:
		return v.(float32) < val
	case float64:
		return v.(float64) < val
	}

	return true
}

func (w LessThan) String() string {
	return fmt.Sprintf("%v < %v", w.Name, w.Value)
}

type Equal struct {
	Name  string
	Value any
}

func (w Equal) Apply(input any) bool {
	v := reflect.ValueOf(input).FieldByName(w.Name).Interface()

	switch val := w.Value.(type) {
	case int:
		return v.(int) == val
	case int32:
		return v.(int32) == val
	case int64:
		return v.(int64) == val
	case float32:
		return v.(float32) == val
	case float64:
		return v.(float64) == val
	case string:
		return v.(string) == val
	}

	return true
}

func (w Equal) String() string {
	return fmt.Sprintf("%v = %v", w.Name, w.Value)
}

type NotEqual struct {
	Name  string
	Value any
}

func (w NotEqual) Apply(input any) bool {
	v := reflect.ValueOf(input).FieldByName(w.Name).Interface()

	switch val := w.Value.(type) {
	case int:
		return v.(int) != val
	case int32:
		return v.(int32) != val
	case int64:
		return v.(int64) != val
	case float32:
		return v.(float32) != val
	case float64:
		return v.(float64) != val
	case string:
		return v.(string) != val
	}

	return true
}

func (w NotEqual) String() string {
	return fmt.Sprintf("%v != %v", w.Name, w.Value)
}

type And struct {
	Lhs Where
	Rhs Where
}

func (w And) Apply(input any) bool {
	return w.Lhs.Apply(input) && w.Rhs.Apply(input)
}

func (w And) String() string {
	return fmt.Sprintf("%v AND %v", w.Lhs, w.Rhs)
}
