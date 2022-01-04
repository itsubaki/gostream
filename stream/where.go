package stream

import (
	"fmt"
	"reflect"
)

type Where interface {
	Apply(input interface{}) bool
	String() string
}

type From struct {
	Type interface{}
}

func (w From) Apply(input interface{}) bool {
	return reflect.TypeOf(input) == reflect.TypeOf(w.Type)
}

func (w From) String() string {
	return reflect.TypeOf(w.Type).Name()
}

type LargerThan struct {
	Name  string
	Value interface{}
}

func (w LargerThan) Apply(input interface{}) bool {
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
	Value interface{}
}

func (w LessThan) Apply(input interface{}) bool {
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

type Equals struct {
	Name  string
	Value interface{}
}

func (w Equals) Apply(input interface{}) bool {
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

func (w Equals) String() string {
	return fmt.Sprintf("%v = %v", w.Name, w.Value)
}

type NotEquals struct {
	Name  string
	Value interface{}
}

func (w NotEquals) Apply(input interface{}) bool {
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

func (w NotEquals) String() string {
	return fmt.Sprintf("%v != %v", w.Name, w.Value)
}

type AND struct {
	Lhs Where
	Rhs Where
}

func (w AND) Apply(input interface{}) bool {
	return w.Lhs.Apply(input) && w.Rhs.Apply(input)
}

func (w AND) String() string {
	return fmt.Sprintf("%v AND %v", w.Lhs, w.Rhs)
}
