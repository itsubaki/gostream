package stream

import (
	"reflect"
	"strings"
)

var (
	_ Selector = (*SelectAll)(nil)
	_ Selector = (*Select)(nil)
)

type Selector interface {
	Apply(e []Event) []Event
	String() string
}

type SelectAll struct{}

func (s SelectAll) Apply(e []Event) []Event {
	v := reflect.ValueOf(e[len(e)-1].Underlying)
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		e[len(e)-1].ResultSet = append(e[len(e)-1].ResultSet, v.Field(i).Interface())
	}

	return e
}

func (s SelectAll) String() string {
	return "*"
}

type Select struct {
	Name string
}

func (s Select) Apply(e []Event) []Event {
	v := reflect.ValueOf(e[len(e)-1].Underlying)
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Name == strings.Trim(s.Name, "`") {
			e[len(e)-1].ResultSet = append(e[len(e)-1].ResultSet, v.Field(i).Interface())
		}
	}

	return e
}

func (s Select) String() string {
	return s.Name
}
