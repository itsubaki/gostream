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

type Average struct {
	Name string
}

func (s Average) Apply(e []Event) []Event {
	var sum float64
	for _, ev := range e {
		v := reflect.ValueOf(ev.Underlying)
		t := v.Type()
		for i := 0; i < t.NumField(); i++ {
			val := v.Field(i).Interface()
			switch val := val.(type) {
			case int:
				sum += float64(val)
			case int32:
				sum += float64(val)
			case int64:
				sum += float64(val)
			case float32:
				sum += float64(val)
			case float64:
				sum += val
			}
		}
	}

	newest := Newest(e)
	newest.ResultSet = append(newest.ResultSet, sum/float64(len(e)))
	return e
}

func (s Average) String() string {
	return fmt.Sprintf("AVG(%v)", s.Name)
}
