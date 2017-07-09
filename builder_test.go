package gocep

import (
	"reflect"
	"testing"
)

func TestNewStructBuilder(t *testing.T) {
	b := NewStructBuilder()
	b.Field("Name", reflect.TypeOf(""))
	b.Field("Record", reflect.TypeOf(make(map[string]interface{})))
	tpe := b.Build()

	rf := reflect.New(tpe).Elem()
	rf.Field(0).SetString("foobar")
	rf.Field(1).Set(reflect.MakeMap(reflect.TypeOf(make(map[string]interface{}))))
	rf.Field(1).SetMapIndex(reflect.ValueOf("foo"), reflect.ValueOf("bar"))
	rf.Field(1).SetMapIndex(reflect.ValueOf("val"), reflect.ValueOf(123))

	e := NewEvent(rf.Interface())
	if e.String("Name") != "foobar" {
		t.Error(e)
	}
	if e.MapString("Record", "foo") != "bar" {
		t.Error(e)
	}
	if e.MapInt("Record", "val") != 123 {
		t.Error(e)
	}

	//	fmt.Printf("%+v\n", &MapEvent{})
	//	fmt.Println(rf.Addr().Interface())
	//	fmt.Printf("%+v\n", rf.Addr().Interface())
	//	fmt.Println("-------")

}
