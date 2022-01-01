package gostream

import "reflect"

type Where interface {
	Apply(input interface{}) bool
}

type EqualsType struct {
	Accept interface{}
}

func (f EqualsType) Apply(input interface{}) bool {
	return reflect.TypeOf(input) == reflect.TypeOf(f.Accept)
}
