package parser

import (
	"fmt"
	"reflect"
)

func FieldType(event interface{}, fieldname string) reflect.Type {
	v := reflect.ValueOf(event)
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Name == fieldname {
			return f.Type
		}
	}

	panic(fmt.Sprintf("field name=%s not found in event= %v", fieldname, event))
}

func IntField(event interface{}, fieldname string) bool {
	switch FieldType(event, fieldname) {
	case
		reflect.TypeOf(int(0)),
		reflect.TypeOf(int8(0)),
		reflect.TypeOf(int16(0)),
		reflect.TypeOf(int32(0)),
		reflect.TypeOf(int64(0)):
		return true
	default:
		return false
	}
}

func FloatField(event interface{}, fieldname string) bool {
	switch FieldType(event, fieldname) {
	case
		reflect.TypeOf(float32(0)),
		reflect.TypeOf(float64(0)):
		return true
	default:
		return false
	}
}
