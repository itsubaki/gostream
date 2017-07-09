package gocep

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEventBuilder(t *testing.T) {
	b := NewEventBuilder()
	tp := b.Build()

	rf := reflect.New(tp).Elem()
	rf.Field(0).SetString("foobar")
	rf.Field(1).Set(reflect.MakeMap(reflect.TypeOf(make(map[string]interface{}))))
	rf.Field(1).SetMapIndex(reflect.ValueOf("foo"), reflect.ValueOf("bar"))
	rf.Field(1).SetMapIndex(reflect.ValueOf("val"), reflect.ValueOf(123))

	fmt.Println(reflect.TypeOf(MapEvent{}))
	fmt.Println(reflect.TypeOf(rf))
	fmt.Println(rf.FieldByName("Name"))
	fmt.Println(rf.FieldByName("Record"))
	fmt.Println(rf.FieldByName("Record").MapIndex(reflect.ValueOf("foo")))
	fmt.Println(rf.FieldByName("Record").MapIndex(reflect.ValueOf("val")))
}
