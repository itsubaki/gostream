package gocep

import (
	"math"
	"time"
)

type Function interface {
	Apply(event []Event) []Event
}

type Length struct {
	length int
}

func (f Length) Apply(event []Event) []Event {
	if len(event) > f.length {
		event = event[1:]
	}
	return event
}

type TimeDuration struct {
	expire time.Duration
}

func (f TimeDuration) Apply(event []Event) (stream []Event) {
	for _, e := range event {
		if time.Since(e.Time) < f.expire {
			stream = append(stream, e)
		}
	}
	return stream
}

type LengthBatch struct {
	length int
	batch  []Event
}

func (f *LengthBatch) Apply(event []Event) (stream []Event) {
	f.batch = append(f.batch, event[len(event)-1])
	if len(f.batch) == f.length {
		stream, f.batch = f.batch, stream
		return stream
	}
	return stream
}

type TimeDurationBatch struct {
	start  time.Time
	end    time.Time
	expire time.Duration
}

func (f *TimeDurationBatch) Apply(event []Event) (stream []Event) {
	for {
		if time.Since(f.start) < f.expire {
			break
		}
		f.start = f.start.Add(f.expire)
		f.end = f.start.Add(f.expire)
	}

	for _, e := range event {
		if !e.Time.Before(f.start) && !e.Time.After(f.end) {
			stream = append(stream, e)
		}
	}
	return stream
}

type Count struct{}

func (f Count) Apply(event []Event) []Event {
	count := len(event)
	for _, e := range event {
		e.Record["count"] = count
	}
	return event
}

type SumInt struct {
	Name string
}

func (f SumInt) Apply(event []Event) []Event {
	sum := 0
	for _, e := range event {
		sum = sum + e.IntValue(f.Name)
	}

	for _, e := range event {
		e.Record["sum("+f.Name+")"] = sum
	}

	return event
}

type SumFloat struct {
	Name string
}

func (f SumFloat) Apply(event []Event) []Event {
	var sum float64
	for _, e := range event {
		sum = sum + e.FloatValue(f.Name)
	}

	for _, e := range event {
		e.Record["sum("+f.Name+")"] = sum
	}

	return event
}

type AverageInt struct {
	Name string
}

func (f AverageInt) Apply(event []Event) []Event {
	sum := 0
	for _, e := range event {
		sum = sum + e.IntValue(f.Name)
	}
	length := len(event)
	avg := float64(sum) / float64(length)

	for _, e := range event {
		e.Record["avg("+f.Name+")"] = avg
	}

	return event
}

type AverageFloat struct {
	Name string
}

func (f AverageFloat) Apply(event []Event) []Event {
	var sum float64
	for _, e := range event {
		sum = sum + e.FloatValue(f.Name)
	}
	length := len(event)
	avg := float64(sum) / float64(length)

	for _, e := range event {
		e.Record["avg("+f.Name+")"] = avg
	}

	return event
}

type MaxInt struct {
	Name string
}

func (f MaxInt) Apply(event []Event) []Event {
	max := math.MinInt8
	for _, e := range event {
		val := e.IntValue(f.Name)
		if val > max {
			max = val
		}
	}

	for _, e := range event {
		e.Record["max("+f.Name+")"] = max
	}

	return event
}

type MaxFloat struct {
	Name string
}

func (f MaxFloat) Apply(event []Event) []Event {
	max := event[0].FloatValue(f.Name)
	for _, e := range event {
		max = math.Max(max, e.FloatValue(f.Name))
	}

	for _, e := range event {
		e.Record["max("+f.Name+")"] = float64(max)
	}

	return event
}

type MinInt struct {
	Name string
}

func (f MinInt) Apply(event []Event) []Event {
	min := math.MaxInt8
	for _, e := range event {
		val := e.IntValue(f.Name)
		if val < min {
			min = val
		}
	}

	for _, e := range event {
		e.Record["min("+f.Name+")"] = min
	}

	return event
}

type MinFloat struct {
	Name string
}

func (f MinFloat) Apply(event []Event) []Event {
	min := event[0].FloatValue(f.Name)
	for _, e := range event {
		min = math.Min(min, e.FloatValue(f.Name))
	}

	for _, e := range event {
		e.Record["min("+f.Name+")"] = float64(min)
	}

	return event
}

type MedianInt struct {
	Name string
}

func (f MedianInt) Apply(event []Event) []Event {
	values := []int{}
	for _, e := range event {
		values = append(values, e.IntValue(f.Name))
	}

	center := len(values) / 2
	median := float64(values[center])
	if len(values)%2 == 0 {
		median = float64(values[center-1]+values[center]) / float64(2)
	}

	for _, e := range event {
		e.Record["median("+f.Name+")"] = median
	}

	return event
}

type MedianFloat struct {
	Name string
}

func (f MedianFloat) Apply(event []Event) []Event {
	values := []float64{}
	for _, e := range event {
		values = append(values, e.FloatValue(f.Name))
	}

	center := len(values) / 2
	median := float64(values[center])
	if len(values)%2 == 0 {
		median = float64(values[center-1]+values[center]) / float64(2)
	}

	for _, e := range event {
		e.Record["median("+f.Name+")"] = median
	}

	return event
}
