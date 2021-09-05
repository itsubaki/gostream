package function

import "github.com/itsubaki/gostream/pkg/event"

type EqualsMapString struct {
	Name  string
	Key   string
	Value string
}

func (f EqualsMapString) Apply(e event.Event) bool {
	return e.MapString(f.Name, f.Key) == f.Value
}

type EqualsMapBool struct {
	Name  string
	Key   string
	Value bool
}

func (f EqualsMapBool) Apply(e event.Event) bool {
	return e.MapBool(f.Name, f.Key) == f.Value
}

type EqualsMapInt struct {
	Name  string
	Key   string
	Value int
}

func (f EqualsMapInt) Apply(e event.Event) bool {
	return e.MapInt(f.Name, f.Key) == f.Value
}

type EqualsMapFloat struct {
	Name  string
	Key   string
	Value float64
}

func (f EqualsMapFloat) Apply(e event.Event) bool {
	return e.MapFloat(f.Name, f.Key) == f.Value
}

type NotEqualsMapString struct {
	Name  string
	Key   string
	Value string
}

func (f NotEqualsMapString) Apply(e event.Event) bool {
	return e.MapString(f.Name, f.Key) != f.Value
}

type NotEqualsMapBool struct {
	Name  string
	Key   string
	Value bool
}

func (f NotEqualsMapBool) Apply(e event.Event) bool {
	return e.MapBool(f.Name, f.Key) != f.Value
}

type NotEqualsMapInt struct {
	Name  string
	Key   string
	Value int
}

func (f NotEqualsMapInt) Apply(e event.Event) bool {
	return e.MapInt(f.Name, f.Key) != f.Value
}

type NotEqualsMapFloat struct {
	Name  string
	Key   string
	Value float64
}

func (f NotEqualsMapFloat) Apply(e event.Event) bool {
	return e.MapFloat(f.Name, f.Key) != f.Value
}

type LargerThanMapInt struct {
	Name  string
	Key   string
	Value int
}

func (f LargerThanMapInt) Apply(e event.Event) bool {
	return e.MapInt(f.Name, f.Key) > f.Value
}

type LargerThanMapFloat struct {
	Name  string
	Key   string
	Value float64
}

func (f LargerThanMapFloat) Apply(e event.Event) bool {
	return e.MapFloat(f.Name, f.Key) > f.Value
}

type LessThanMapInt struct {
	Name  string
	Key   string
	Value int
}

func (f LessThanMapInt) Apply(e event.Event) bool {
	return e.MapInt(f.Name, f.Key) < f.Value
}

type LessThanMapFloat struct {
	Name  string
	Key   string
	Value float64
}

func (f LessThanMapFloat) Apply(e event.Event) bool {
	return e.MapFloat(f.Name, f.Key) < f.Value
}
