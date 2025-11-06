package stream

import (
	"fmt"
	"reflect"
)

var (
	_ Aggeregator = (*Average)(nil)
	_ Aggeregator = (*Sum)(nil)
	_ Aggeregator = (*Count)(nil)
	_ Aggeregator = (*Max)(nil)
	_ Aggeregator = (*Min)(nil)
	_ Aggeregator = (*Distinct)(nil)
)

type Aggeregator interface {
	Apply(e []Event) []Event
	String() string
}

type Average struct {
	Name string
}

func (s Average) Apply(e []Event) []Event {
	var sum float64
	for _, ev := range e {
		v := reflect.ValueOf(ev.Underlying)
		for i := 0; i < v.Type().NumField(); i++ {
			if v.Type().Field(i).Name != s.Name {
				continue
			}

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

	e[len(e)-1].ResultSet = append(e[len(e)-1].ResultSet, sum/float64(len(e)))
	return e
}

func (s Average) String() string {
	return fmt.Sprintf("AVG(%v)", s.Name)
}

type Sum struct {
	Name string
}

func (s Sum) Apply(e []Event) []Event {
	var sum float64
	for _, ev := range e {
		v := reflect.ValueOf(ev.Underlying)
		for i := 0; i < v.Type().NumField(); i++ {
			if v.Type().Field(i).Name != s.Name {
				continue
			}

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

	e[len(e)-1].ResultSet = append(e[len(e)-1].ResultSet, sum)
	return e
}

func (s Sum) String() string {
	return fmt.Sprintf("SUM(%v)", s.Name)
}

type Count struct {
	Name string
}

func (s Count) Apply(e []Event) []Event {
	e[len(e)-1].ResultSet = append(e[len(e)-1].ResultSet, len(e))
	return e
}

func (s Count) String() string {
	return fmt.Sprintf("COUNT(%v)", s.Name)
}

type Max struct {
	Name string
}

func (s Max) Apply(e []Event) []Event {
	var max float64
	for _, ev := range e {
		v := reflect.ValueOf(ev.Underlying)
		for i := 0; i < v.Type().NumField(); i++ {
			if v.Type().Field(i).Name != s.Name {
				continue
			}

			val := v.Field(i).Interface()
			switch val := val.(type) {
			case int:
				if float64(val) > max {
					max = float64(val)
				}
			case int32:
				if float64(val) > max {
					max = float64(val)
				}
			case int64:
				if float64(val) > max {
					max = float64(val)
				}
			case float32:
				if float64(val) > max {
					max = float64(val)
				}
			case float64:
				if val > max {
					max = val
				}
			}
		}
	}

	e[len(e)-1].ResultSet = append(e[len(e)-1].ResultSet, max)
	return e
}

func (s Max) String() string {
	return fmt.Sprintf("MAX(%v)", s.Name)
}

type Min struct {
	Name string
}

func (s Min) Apply(e []Event) []Event {
	var min float64
	for _, ev := range e {
		v := reflect.ValueOf(ev.Underlying)
		for i := 0; i < v.Type().NumField(); i++ {
			if v.Type().Field(i).Name != s.Name {
				continue
			}

			val := v.Field(i).Interface()
			switch val := val.(type) {
			case int:
				if float64(val) < min {
					min = float64(val)
				}
			case int32:
				if float64(val) < min {
					min = float64(val)
				}
			case int64:
				if float64(val) < min {
					min = float64(val)
				}
			case float32:
				if float64(val) < min {
					min = float64(val)
				}
			case float64:
				if val < min {
					min = val
				}
			}
		}
	}

	e[len(e)-1].ResultSet = append(e[len(e)-1].ResultSet, min)
	return e
}

func (s Min) String() string {
	return fmt.Sprintf("MIN(%v)", s.Name)
}

type Distinct struct {
	Name string
}

func (s Distinct) Apply(e []Event) []Event {
	dist := make(map[interface{}]int)
	for i, ev := range e {
		v := reflect.ValueOf(ev.Underlying)
		for j := 0; j < v.Type().NumField(); j++ {
			if v.Type().Field(j).Name != s.Name {
				continue
			}

			dist[v.Field(j).Interface()] = i
		}
	}

	out := make([]Event, 0)
	for _, v := range dist {
		out = append(out, e[v])
	}

	return out
}

func (s Distinct) String() string {
	return fmt.Sprintf("DISTINCT(%v)", s.Name)
}
