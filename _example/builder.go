package _example

import (
	"fmt"
	"reflect"

	"github.com/itsubaki/gostream/pkg/builder"
)

func Builder() {
	b := builder.New()
	b.SetField("Name", reflect.TypeOf(""))
	b.SetField("Value", reflect.TypeOf(0))
	s := b.Build()

	i := s.NewInstance()
	i.SetString("Name", "foobar")
	i.SetInt("Value", 123)

	fmt.Printf("%#v\n", i.Value())
	fmt.Printf("%#v\n", i.Pointer())
}
