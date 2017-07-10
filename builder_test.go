package gocep

import (
	"reflect"
	"testing"
)

func TestNewStructBuilder(t *testing.T) {
	b := NewStructBuilder()
	b.Field("Name", reflect.TypeOf(""))
	b.Field("Bool", reflect.TypeOf(true))
	b.Field("Int", reflect.TypeOf(123))
	b.Field("Float", reflect.TypeOf(float64(12.3)))
	b.Field("Record", reflect.TypeOf(make(map[string]interface{})))
	b.Field("Record2", reflect.TypeOf(make(map[string]interface{})))
	strct := b.Build()

	i := strct.NewInstance()
	i.SetString("Name", "foobar")
	i.SetBool("Bool", false)
	i.SetInt("Int", 456)
	i.SetFloat("Float", 45.6)
	i.SetMapIndex("Record", "foo", "bar")
	i.SetMapIndex("Record", "val", 123)
	i.SetMapIndex("Record2", "val", 123)

	ptr := i.Pointer()
	if reflect.TypeOf(ptr).Kind() != reflect.Ptr {
		t.Error(i.Pointer())
	}

	ival := i.Interface()
	e := NewEvent(ival)
	if e.String("Name") != "foobar" {
		t.Error(e)
	}

	if e.Bool("Bool") {
		t.Error(e)
	}

	if e.Int("Int") != 456 {
		t.Error(e)
	}

	if e.Float("Float") != 45.6 {
		t.Error(e)
	}

	if e.MapString("Record", "foo") != "bar" {
		t.Error(e)
	}

	if e.MapInt("Record", "val") != 123 {
		t.Error(e)
	}

	if e.MapInt("Record2", "val") != 123 {
		t.Error(e)
	}

}
