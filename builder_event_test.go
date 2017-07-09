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

	fmt.Println(reflect.TypeOf(&MapEvent{}))
	fmt.Printf("%+v\n", &MapEvent{})
	fmt.Println(rf.Addr().Interface())
	fmt.Printf("%+v\n", rf.Addr().Interface())
}
