package gocep

import "reflect"

type StructBuilder struct {
	field []reflect.StructField
	strct reflect.Type
	index map[string]int
	imap  []int
}

func NewStructBuilder() *StructBuilder {
	return &StructBuilder{}
}

func (b *StructBuilder) Field(fname string, ftype reflect.Type) {
	b.field = append(
		b.field,
		reflect.StructField{
			Name: fname,
			Type: ftype,
		})
}

func (b *StructBuilder) Build() {
	b.strct = reflect.StructOf(b.field)
	b.index = make(map[string]int)
	for i := 0; i < b.strct.NumField(); i++ {
		f := b.strct.Field(i)
		b.index[f.Name] = i
		if f.Type.Kind() == reflect.Map {
			b.imap = append(b.imap, i)
		}
	}
}

func (b *StructBuilder) NewInstance() *Instance {
	instance := reflect.New(b.strct).Elem()
	mtype := reflect.TypeOf(make(map[string]interface{}))
	for _, i := range b.imap {
		mval := reflect.MakeMap(mtype)
		instance.Field(i).Set(mval)
	}
	return &Instance{instance, b.index}
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

func (inst *Instance) Build() interface{} {
	return inst.internal.Interface()
}

//func (inst *Instance) Pointer() interface{} {
//	return inst.internal.Addr().Interface()
//}
