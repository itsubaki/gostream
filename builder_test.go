package gocep

import (
	"reflect"
	"testing"
)

func TestNewStructBuilder(t *testing.T) {
	b := NewStructBuilder()
	b.Field("Name", reflect.TypeOf(""))
	b.Field("Record", reflect.TypeOf(make(map[string]interface{})))
	s := b.Build()

	i := reflect.New(s).Elem()
	i.Field(0).SetString("foobar")
	i.Field(1).Set(reflect.MakeMap(reflect.TypeOf(make(map[string]interface{}))))
	i.Field(1).SetMapIndex(reflect.ValueOf("foo"), reflect.ValueOf("bar"))
	i.Field(1).SetMapIndex(reflect.ValueOf("val"), reflect.ValueOf(123))

	e := NewEvent(i.Interface())
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
	//	fmt.Println(i.Addr().Interface())
	//	fmt.Printf("%+v\n", i.Addr().Interface())
	//	fmt.Println("-------")

}
