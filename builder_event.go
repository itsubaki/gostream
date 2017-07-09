package gocep

import "reflect"

type EevntBuilder struct {
}

func NewEventBuilder() *EevntBuilder {
	return &EevntBuilder{}
}

func (b *EevntBuilder) Build() reflect.Type {
	return reflect.StructOf(
		[]reflect.StructField{
			reflect.StructField{
				Name: "Name",
				Type: reflect.TypeOf(""),
			},
			reflect.StructField{
				Name: "Record",
				Type: reflect.TypeOf(make(map[string]interface{})),
			},
		})
}
