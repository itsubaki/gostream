package gocep

import (
	"reflect"
	"time"
)

type InsertEvent struct {
	Record map[string]interface{}
}

type Event struct {
	Time       time.Time
	Underlying interface{}
	Record     map[string]interface{}
}

func NewEvent(underlying interface{}) Event {
	return Event{time.Now(), underlying, make(map[string]interface{})}
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

func (e Event) String(name string) string {
	return e.Value(name).Interface().(string)
}

func (e Event) Bool(name string) bool {
	return e.Value(name).Interface().(bool)
}

func (e Event) Int(name string) int {
	return e.Value(name).Interface().(int)
}

func (e Event) Float(name string) float64 {
	return e.Value(name).Interface().(float64)
}

func (e Event) Map(name, key string) reflect.Value {
	return e.Value(name).MapIndex(reflect.ValueOf(key))
}

func (e Event) MapString(name, key string) string {
	return e.Map(name, key).Interface().(string)
}

func (e Event) MapBool(name, key string) bool {
	return e.Map(name, key).Interface().(bool)
}

func (e Event) MapInt(name, key string) int {
	return e.Map(name, key).Interface().(int)
}

func (e Event) MapFloat(name, key string) float64 {
	return e.Map(name, key).Interface().(float64)
}

func (e Event) RecordString(name string) string {
	return e.Record[name].(string)
}

func (e Event) RecordBool(name string) bool {
	return e.Record[name].(bool)
}

func (e Event) RecordInt(name string) int {
	return e.Record[name].(int)
}

func (e Event) RecordFloat(name string) float64 {
	return e.Record[name].(float64)
}
