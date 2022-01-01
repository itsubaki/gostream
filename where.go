package gostream

import "reflect"

type Where interface {
	Apply(input interface{}) bool
}

type Accept struct {
	Type interface{}
}

func (w Accept) Apply(input interface{}) bool {
	return reflect.TypeOf(input) == reflect.TypeOf(w.Type)
}
