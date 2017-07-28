package gocep

import (
	"reflect"
)

type StructBuilder struct {
	field []reflect.StructField
}

func NewStructBuilder() *StructBuilder {
	return &StructBuilder{}
}

func (b *StructBuilder) SetField(fname string, ftype reflect.Type) {
	b.field = append(
		b.field,
		reflect.StructField{
			Name: fname,
			Type: ftype,
		})
}

func (b *StructBuilder) Build() Struct {
	strct := reflect.StructOf(b.field)
	index := make(map[string]int)
	imap := []int{}
	for i := 0; i < strct.NumField(); i++ {
		f := strct.Field(i)
		index[f.Name] = i
		if f.Type.Kind() == reflect.Map {
			imap = append(imap, i)
		}
	}
	return Struct{strct, index, imap}
}

type Struct struct {
	strct reflect.Type
	index map[string]int
	imap  []int
}

func (s *Struct) NewInstance() *Instance {
	instance := reflect.New(s.strct).Elem()
	mtype := reflect.TypeOf(make(map[string]interface{}))
	for _, i := range s.imap {
		mval := reflect.MakeMap(mtype)
		instance.Field(i).Set(mval)
	}
	return &Instance{instance, s.index}
}

type Instance struct {
	internal reflect.Value
	index    map[string]int
}

func (i *Instance) Field(name string) reflect.Value {
	return i.internal.Field(i.index[name])
}

func (i *Instance) SetString(name, value string) {
	i.Field(name).SetString(value)
}

func (i *Instance) SetBool(name string, value bool) {
	i.Field(name).SetBool(value)
}

func (i *Instance) SetInt(name string, value int) {
	i.Field(name).SetInt(int64(value))
}

func (i *Instance) SetFloat(name string, value float64) {
	i.Field(name).SetFloat(value)
}

func (i *Instance) SetMapIndex(name, key string, value interface{}) {
	refkey := reflect.ValueOf(key)
	refval := reflect.ValueOf(value)
	i.Field(name).SetMapIndex(refkey, refval)
}

func (i *Instance) Value() interface{} {
	return i.internal.Interface()
}

func (i *Instance) Pointer() interface{} {
	return i.internal.Addr().Interface()
}
