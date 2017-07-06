package gocep

import (
	"reflect"
	"time"
)

type Event struct {
	Time       time.Time
	Underlying interface{}
	Record     map[string]interface{}
}

func NewEvent(underlying interface{}) Event {
	return Event{time.Now(), underlying, nil}
}

func (e Event) New() Event {
	return Event{
		e.Time,
		e.Underlying,
		make(map[string]interface{}),
	}
}

func (e Event) Type() reflect.Type {
	return reflect.TypeOf(e.Underlying)
}

func (e Event) EqualsType(t reflect.Type) bool {
	return e.Type() == t
}

func (e Event) Value(name string) reflect.Value {
	return reflect.ValueOf(e.Underlying).FieldByName(name)
}

func (e Event) StringValue(name string) string {
	return e.Value(name).Interface().(string)
}

func (e Event) BoolValue(name string) bool {
	return e.Value(name).Interface().(bool)
}

func (e Event) IntValue(name string) int {
	return e.Value(name).Interface().(int)
}

func (e Event) FloatValue(name string) float64 {
	return e.Value(name).Interface().(float64)
}

func (e Event) MapValue(name, key string) reflect.Value {
	return e.Value(name).MapIndex(reflect.ValueOf(key))
}

func (e Event) MapStringValue(name, key string) string {
	return e.MapValue(name, key).Interface().(string)
}

func (e Event) MapBoolValue(name, key string) bool {
	return e.MapValue(name, key).Interface().(bool)
}

func (e Event) MapIntValue(name, key string) int {
	return e.MapValue(name, key).Interface().(int)
}

func (e Event) MapFloatValue(name, key string) float64 {
	return e.MapValue(name, key).Interface().(float64)
}

func (e Event) RecordIntValue(name string) int {
	return e.Record[name].(int)
}

func (e Event) RecordFloatValue(name string) float64 {
	return e.Record[name].(float64)
}

func (e Event) RecordBoolValue(name string) bool {
	return e.Record[name].(bool)
}
