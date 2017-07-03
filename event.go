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

func (e Event) TypeEquals(t reflect.Type) bool {
	return e.Type() == t
}

func (e Event) Value(name string) reflect.Value {
	return reflect.ValueOf(e.Underlying).FieldByName(name)
}

func (e Event) StringValue(name string) string {
	return e.Value(name).String()
}

func (e Event) BoolValue(name string) bool {
	return e.Value(name).Bool()
}

func (e Event) IntValue(name string) int {
	return int(e.Value(name).Int())
}

func (e Event) Float32Value(name string) float32 {
	return float32(e.Value(name).Float())
}

func (e Event) Float64Value(name string) float64 {
	return float64(e.Value(name).Float())
}
