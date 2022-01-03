package stream

import (
	"fmt"
	"reflect"
	"strings"
)

type SelectIF interface {
	Apply(e []Event) []Event
	String() string
}

type SelectAll struct{}

func (s SelectAll) Apply(e []Event) []Event {
	newest := Newest(e)

	v := reflect.ValueOf(newest.Underlying)
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		newest.ResultSet = append(newest.ResultSet, v.Field(i).Interface())
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
	newest := Newest(e)

	v := reflect.ValueOf(newest.Underlying)
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Name == strings.Trim(s.Name, "`") {
			newest.ResultSet = append(newest.ResultSet, v.Field(i).Interface())
		}
	}

	return e
}

func (s Select) String() string {
	return s.Name
}

type Distinct struct {
	Name string
}

func (s Distinct) Apply(e []Event) []Event {
	return e
}

func (s Distinct) String() string {
	return fmt.Sprintf("DISTINCT(%v)", s.Name)
}

type AVG struct {
	Name string
}

func (s AVG) Apply(e []Event) []Event {
	return e
}

func (s AVG) String() string {
	return fmt.Sprintf("AVG(%v)", s.Name)
}
