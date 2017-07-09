package gocep

import (
	"math"
	"reflect"
	"strconv"
)

type Function interface {
	Apply(event []Event) []Event
}

type SelectAll struct {
}

func (f SelectAll) Apply(event []Event) []Event {
	for _, e := range event {
		t := reflect.ValueOf(e.Underlying).Type()
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			v := e.Value(f.Name)
			e.Record[f.Name] = v.Interface()
		}
	}
	return event
}

type SelectString struct {
	Name string
	As   string
}

func (f SelectString) Apply(event []Event) []Event {
	for _, e := range event {
		e.Record[f.As] = e.String(f.Name)
	}
	return event
}

type SelectBool struct {
	Name string
	As   string
}

func (f SelectBool) Apply(event []Event) []Event {
	for _, e := range event {
		e.Record[f.As] = e.Bool(f.Name)
	}
	return event
}

type SelectInt struct {
	Name string
	As   string
}

func (f SelectInt) Apply(event []Event) []Event {
	for _, e := range event {
		e.Record[f.As] = e.Int(f.Name)
	}
	return event
}

type SelectFloat struct {
	Name string
	As   string
}

func (f SelectFloat) Apply(event []Event) []Event {
	for _, e := range event {
		e.Record[f.As] = e.Float(f.Name)
	}
	return event
}

type Count struct {
	As string
}

func (f Count) Apply(event []Event) []Event {
	count := len(event)
	for _, e := range event {
		e.Record[f.As] = count
	}
	return event
}

type SumInt struct {
	Name string
	As   string
}

func (f SumInt) Apply(event []Event) []Event {
	sum := 0
	for _, e := range event {
		sum = sum + e.Int(f.Name)
	}

	for _, e := range event {
		e.Record[f.As] = sum
	}

	return event
}

type SumFloat struct {
	Name string
	As   string
}

func (f SumFloat) Apply(event []Event) []Event {
	var sum float64
	for _, e := range event {
		sum = sum + e.Float(f.Name)
	}

	for _, e := range event {
		e.Record[f.As] = sum
	}

	return event
}

type AverageInt struct {
	Name string
	As   string
}

func (f AverageInt) Apply(event []Event) []Event {
	sum := 0
	for _, e := range event {
		sum = sum + e.Int(f.Name)
	}
	length := len(event)
	avg := float64(sum) / float64(length)

	for _, e := range event {
		e.Record[f.As] = avg
	}

	return event
}

type AverageFloat struct {
	Name string
	As   string
}

func (f AverageFloat) Apply(event []Event) []Event {
	var sum float64
	for _, e := range event {
		sum = sum + e.Float(f.Name)
	}
	length := len(event)
	avg := sum / float64(length)

	for _, e := range event {
		e.Record[f.As] = avg
	}

	return event
}

type MaxInt struct {
	Name string
	As   string
}

func (f MaxInt) Apply(event []Event) []Event {
	max := math.MinInt8
	for _, e := range event {
		val := e.Int(f.Name)
		if val > max {
			max = val
		}
	}

	for _, e := range event {
		e.Record[f.As] = max
	}

	return event
}

type MaxFloat struct {
	Name string
	As   string
}

func (f MaxFloat) Apply(event []Event) []Event {
	max := event[0].Float(f.Name)
	for _, e := range event {
		max = math.Max(max, e.Float(f.Name))
	}

	for _, e := range event {
		e.Record[f.As] = max
	}

	return event
}

type MinInt struct {
	Name string
	As   string
}

func (f MinInt) Apply(event []Event) []Event {
	min := math.MaxInt8
	for _, e := range event {
		val := e.Int(f.Name)
		if val < min {
			min = val
		}
	}

	for _, e := range event {
		e.Record[f.As] = min
	}

	return event
}

type MinFloat struct {
	Name string
	As   string
}

func (f MinFloat) Apply(event []Event) []Event {
	min := event[0].Float(f.Name)
	for _, e := range event {
		min = math.Min(min, e.Float(f.Name))
	}

	for _, e := range event {
		e.Record[f.As] = min
	}

	return event
}

type MedianInt struct {
	Name string
	As   string
}

func (f MedianInt) Apply(event []Event) []Event {
	values := []int{}
	for _, e := range event {
		values = append(values, e.Int(f.Name))
	}

	center := len(values) / 2
	median := float64(values[center])
	if len(values)%2 == 0 {
		median = float64(values[center-1]+values[center]) / float64(2)
	}

	for _, e := range event {
		e.Record[f.As] = median
	}

	return event
}

type MedianFloat struct {
	Name string
	As   string
}

func (f MedianFloat) Apply(event []Event) []Event {
	values := []float64{}
	for _, e := range event {
		values = append(values, e.Float(f.Name))
	}

	center := len(values) / 2
	median := values[center]
	if len(values)%2 == 0 {
		median = float64(values[center-1]+values[center]) / float64(2)
	}

	for _, e := range event {
		e.Record[f.As] = median
	}

	return event
}

type CastStringToInt struct {
	Name string
	As   string
}

func (f CastStringToInt) Apply(event []Event) []Event {
	for _, e := range event {
		str := e.String(f.Name)
		val, _ := strconv.Atoi(str)
		e.Record[f.As] = val
	}

	return event
}

type CastStringToFloat struct {
	Name string
	As   string
}

func (f CastStringToFloat) Apply(event []Event) []Event {
	for _, e := range event {
		str := e.String(f.Name)
		val, _ := strconv.ParseFloat(str, 64)
		e.Record[f.As] = val
	}

	return event
}

type CastStringToBool struct {
	Name string
	As   string
}

func (f CastStringToBool) Apply(event []Event) []Event {
	for _, e := range event {
		str := e.String(f.Name)
		val, _ := strconv.ParseBool(str)
		e.Record[f.As] = val
	}

	return event
}
