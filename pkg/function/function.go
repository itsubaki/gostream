package function

import (
	"math"
	"strconv"

	"github.com/itsubaki/gostream/pkg/event"
)

type Function interface {
	Apply(events []event.Event) []event.Event
}

type SelectAll struct {
}

func (f SelectAll) Apply(events []event.Event) []event.Event {
	e := event.Newest(events)

	t := e.ValueType()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		v := e.Value(f.Name)
		e.Record[f.Name] = v.Interface()
	}

	return events
}

type SelectString struct {
	Name string
	As   string
}

func (f SelectString) Apply(events []event.Event) []event.Event {
	e := event.Newest(events)
	e.Record[f.As] = e.String(f.Name)

	return events
}

type SelectBool struct {
	Name string
	As   string
}

func (f SelectBool) Apply(events []event.Event) []event.Event {
	e := event.Newest(events)
	e.Record[f.As] = e.Bool(f.Name)

	return events
}

type SelectInt struct {
	Name string
	As   string
}

func (f SelectInt) Apply(events []event.Event) []event.Event {
	e := event.Newest(events)
	e.Record[f.As] = e.Int(f.Name)

	return events
}

type SelectFloat struct {
	Name string
	As   string
}

func (f SelectFloat) Apply(events []event.Event) []event.Event {
	e := event.Newest(events)
	e.Record[f.As] = e.Float(f.Name)

	return events
}

type Count struct {
	As string
}

func (f Count) Apply(events []event.Event) []event.Event {
	e := event.Newest(events)
	e.Record[f.As] = len(events)

	return events
}

type SumInt struct {
	Name string
	As   string
}

func (f SumInt) Apply(events []event.Event) []event.Event {
	sum := 0
	for _, e := range events {
		sum = sum + e.Int(f.Name)
	}

	e := event.Newest(events)
	e.Record[f.As] = sum

	return events
}

type SumFloat struct {
	Name string
	As   string
}

func (f SumFloat) Apply(events []event.Event) []event.Event {
	var sum float64
	for _, e := range events {
		sum = sum + e.Float(f.Name)
	}

	e := event.Newest(events)
	e.Record[f.As] = sum

	return events
}

type AverageInt struct {
	Name string
	As   string
}

func (f AverageInt) Apply(events []event.Event) []event.Event {
	sum := 0
	for _, e := range events {
		sum = sum + e.Int(f.Name)
	}
	avg := float64(sum) / float64(len(events))

	e := event.Newest(events)
	e.Record[f.As] = avg

	return events
}

type AverageFloat struct {
	Name string
	As   string
}

func (f AverageFloat) Apply(events []event.Event) []event.Event {
	var sum float64
	for _, e := range events {
		sum = sum + e.Float(f.Name)
	}
	avg := sum / float64(len(events))

	e := event.Newest(events)
	e.Record[f.As] = avg

	return events
}

type MaxInt struct {
	Name string
	As   string
}

func (f MaxInt) Apply(events []event.Event) []event.Event {
	max := math.MinInt8
	for _, e := range events {
		val := e.Int(f.Name)
		if val > max {
			max = val
		}
	}

	e := event.Newest(events)
	e.Record[f.As] = max

	return events
}

type MaxFloat struct {
	Name string
	As   string
}

func (f MaxFloat) Apply(events []event.Event) []event.Event {
	max := events[0].Float(f.Name)
	for _, e := range events {
		max = math.Max(max, e.Float(f.Name))
	}

	e := event.Newest(events)
	e.Record[f.As] = max

	return events
}

type MinInt struct {
	Name string
	As   string
}

func (f MinInt) Apply(events []event.Event) []event.Event {
	min := math.MaxInt8
	for _, e := range events {
		val := e.Int(f.Name)
		if val < min {
			min = val
		}
	}

	e := event.Newest(events)
	e.Record[f.As] = min

	return events
}

type MinFloat struct {
	Name string
	As   string
}

func (f MinFloat) Apply(events []event.Event) []event.Event {
	min := events[0].Float(f.Name)
	for _, e := range events {
		min = math.Min(min, e.Float(f.Name))
	}

	e := event.Newest(events)
	e.Record[f.As] = min

	return events
}

type MedianInt struct {
	Name string
	As   string
}

func (f MedianInt) Apply(events []event.Event) []event.Event {
	values := []int{}
	for _, e := range events {
		values = append(values, e.Int(f.Name))
	}

	center := len(values) / 2
	median := float64(values[center])
	if len(values)%2 == 0 {
		median = float64(values[center-1]+values[center]) / float64(2)
	}

	e := event.Newest(events)
	e.Record[f.As] = median

	return events
}

type MedianFloat struct {
	Name string
	As   string
}

func (f MedianFloat) Apply(events []event.Event) []event.Event {
	values := []float64{}
	for _, e := range events {
		values = append(values, e.Float(f.Name))
	}

	center := len(values) / 2
	median := values[center]
	if len(values)%2 == 0 {
		median = (values[center-1] + values[center]) / float64(2)
	}

	e := event.Newest(events)
	e.Record[f.As] = median

	return events
}

type CastStringToInt struct {
	Name string
	As   string
}

func (f CastStringToInt) Apply(events []event.Event) []event.Event {
	e := event.Newest(events)
	str := e.String(f.Name)
	val, _ := strconv.Atoi(str)
	e.Record[f.As] = val

	return events
}

type CastStringToFloat struct {
	Name string
	As   string
}

func (f CastStringToFloat) Apply(events []event.Event) []event.Event {
	e := event.Newest(events)
	str := e.String(f.Name)
	val, _ := strconv.ParseFloat(str, 64)
	e.Record[f.As] = val

	return events
}

type CastStringToBool struct {
	Name string
	As   string
}

func (f CastStringToBool) Apply(events []event.Event) []event.Event {
	e := event.Newest(events)
	str := e.String(f.Name)
	val, _ := strconv.ParseBool(str)
	e.Record[f.As] = val

	return events
}

type FuncEqualsInt struct {
	Function Function
	Name     string
	Value    int
}

func (f FuncEqualsInt) Apply(events []event.Event) []event.Event {
	e := f.Function.Apply(events)
	val := e[len(e)-1].RecordInt(f.Name)
	if val == f.Value {
		return events
	}

	return event.List()
}

type FuncLargerThanInt struct {
	Function Function
	Name     string
	Value    int
}

func (f FuncLargerThanInt) Apply(events []event.Event) []event.Event {
	e := f.Function.Apply(events)
	val := e[len(e)-1].RecordInt(f.Name)
	if val > f.Value {
		return events
	}

	return event.List()
}

type FuncLessThanInt struct {
	Function Function
	Name     string
	Value    int
}

func (f FuncLessThanInt) Apply(events []event.Event) []event.Event {
	e := f.Function.Apply(events)
	val := e[len(e)-1].RecordInt(f.Name)
	if val < f.Value {
		return events
	}

	return event.List()
}
