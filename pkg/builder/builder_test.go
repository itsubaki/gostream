package builder

import (
	"reflect"
	"testing"

	"github.com/itsubaki/gostream-core/pkg/event"
)

func TestStructBuilder(t *testing.T) {
	b := New()
	b.SetField("Name", reflect.TypeOf(""))

	b.SetField("Bool", reflect.TypeOf(true))
	b.SetField("Int", reflect.TypeOf(123))
	b.SetField("Float", reflect.TypeOf(float64(12.3)))
	b.SetField("Record", reflect.TypeOf(make(map[string]interface{})))
	b.SetField("Record2", reflect.TypeOf(make(map[string]interface{})))
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

	val := i.Value()
	e := event.New(val)
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
