package gocep

import "reflect"

type StructBuilder struct {
	field []reflect.StructField
}

func NewStructBuilder() *StructBuilder {
	return &StructBuilder{}
}

func (b *StructBuilder) Field(name string, tpe reflect.Type) {
	b.field = append(
		b.field,
		reflect.StructField{
			Name: name,
			Type: tpe,
		})
}

func (b *StructBuilder) Build() reflect.Type {
	return reflect.StructOf(b.field)
}
