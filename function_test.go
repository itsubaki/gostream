package gocep

import (
	"testing"
	"time"
)

func TestTimeDurationBatch(t *testing.T) {
	start := time.Now()
	duration := 10 * time.Millisecond
	f := &TimeDurationBatch{start, start.Add(duration), duration}

	event := []Event{}
	event = append(event, NewEvent(IntEvent{"foo", 1}))
	event = f.Apply(event)
	if len(event) != 1 {
		t.Error(event)
	}
	time.Sleep(duration)

	event = f.Apply(event)
	if len(event) != 0 {
		t.Error(event)
	}
}

func TestSumInt(t *testing.T) {
	event := []Event{}
	event = append(event, NewEvent(IntEvent{"foo", 10}).New())
	event = append(event, NewEvent(IntEvent{"foo", 20}).New())

	f := SumInt{"Value"}
	result := f.Apply(event)

	var test = []struct {
		index int
		sum   int
	}{
		{0, 30},
		{1, 30},
	}

	for _, tt := range test {
		if result[tt.index].Record["sum(Value)"] != tt.sum {
			t.Error(result)
		}
	}
}

func TestSumFloat(t *testing.T) {
	event := []Event{}
	event = append(event, NewEvent(FloatEvent{"foo", 10.0}).New())
	event = append(event, NewEvent(FloatEvent{"foo", 20.0}).New())

	f := SumFloat{"Value"}
	result := f.Apply(event)

	var test = []struct {
		index int
		sum   float64
	}{
		{0, 30.0},
		{1, 30.0},
	}

	for _, tt := range test {
		if result[tt.index].Record["sum(Value)"] != tt.sum {
			t.Error(result)
		}
	}
}

func TestAverageInt(t *testing.T) {
	event := []Event{}
	event = append(event, NewEvent(IntEvent{"foo", 10}).New())
	event = append(event, NewEvent(IntEvent{"foo", 20}).New())

	f := AverageInt{"Value"}
	result := f.Apply(event)

	var test = []struct {
		index int
		avg   float64
	}{
		{0, 15.0},
		{1, 15.0},
	}

	for _, tt := range test {
		if result[tt.index].Record["avg(Value)"] != tt.avg {
			t.Error(result)
		}
	}
}

func TestAverageFloat(t *testing.T) {
	event := []Event{}
	event = append(event, NewEvent(FloatEvent{"foo", 10.0}).New())
	event = append(event, NewEvent(FloatEvent{"foo", 20.0}).New())

	f := AverageFloat{"Value"}
	result := f.Apply(event)

	var test = []struct {
		index int
		avg   float64
	}{
		{0, 15.0},
		{1, 15.0},
	}

	for _, tt := range test {
		if result[tt.index].Record["avg(Value)"] != tt.avg {
			t.Error(result)
		}
	}
}

func TestMaxInt(t *testing.T) {
	event := []Event{}
	event = append(event, NewEvent(IntEvent{"foo", 10}).New())
	event = append(event, NewEvent(IntEvent{"foo", 20}).New())

	f := MaxInt{"Value"}
	result := f.Apply(event)

	var test = []struct {
		index int
		max   int
	}{
		{0, 20},
		{1, 20},
	}

	for _, tt := range test {
		if result[tt.index].Record["max(Value)"] != tt.max {
			t.Error(result)
		}
	}
}

func TestMaxFloat(t *testing.T) {
	event := []Event{}
	event = append(event, NewEvent(FloatEvent{"foo", 10}).New())
	event = append(event, NewEvent(FloatEvent{"foo", 20}).New())

	f := MaxFloat{"Value"}
	result := f.Apply(event)

	var test = []struct {
		index int
		max   float64
	}{
		{0, 20.0},
		{1, 20.0},
	}

	for _, tt := range test {
		if result[tt.index].Record["max(Value)"] != tt.max {
			t.Error(result)
		}
	}
}

func TestMinInt(t *testing.T) {
	event := []Event{}
	event = append(event, NewEvent(IntEvent{"foo", 10}).New())
	event = append(event, NewEvent(IntEvent{"foo", 20}).New())

	f := MinInt{"Value"}
	result := f.Apply(event)

	var test = []struct {
		index int
		min   int
	}{
		{0, 10},
		{1, 10},
	}

	for _, tt := range test {
		if result[tt.index].Record["min(Value)"] != tt.min {
			t.Error(result)
		}
	}
}

func TestMinFloat(t *testing.T) {
	event := []Event{}
	event = append(event, NewEvent(FloatEvent{"foo", 10}).New())
	event = append(event, NewEvent(FloatEvent{"foo", 20}).New())

	f := MinFloat{"Value"}
	result := f.Apply(event)

	var test = []struct {
		index int
		max   float64
	}{
		{0, 10.0},
		{1, 10.0},
	}

	for _, tt := range test {
		if result[tt.index].Record["min(Value)"] != tt.max {
			t.Error(result)
		}
	}
}

func TestMedianIntEven(t *testing.T) {
	event := []Event{}
	event = append(event, NewEvent(IntEvent{"foo", 10}).New())
	event = append(event, NewEvent(IntEvent{"foo", 20}).New())

	f := MedianInt{"Value"}
	result := f.Apply(event)

	var test = []struct {
		index  int
		median float64
	}{
		{0, 15.0},
		{1, 15.0},
	}

	for _, tt := range test {
		if result[tt.index].Record["median(Value)"] != tt.median {
			t.Error(result)
		}
	}
}

func TestMedianIntOdd(t *testing.T) {
	event := []Event{}
	event = append(event, NewEvent(IntEvent{"foo", 10}).New())
	event = append(event, NewEvent(IntEvent{"foo", 20}).New())
	event = append(event, NewEvent(IntEvent{"foo", 30}).New())

	f := MedianInt{"Value"}
	result := f.Apply(event)

	var test = []struct {
		index  int
		median float64
	}{
		{0, 20.0},
		{1, 20.0},
	}

	for _, tt := range test {
		if result[tt.index].Record["median(Value)"] != tt.median {
			t.Error(result)
		}
	}
}

func TestMedianFloatEven(t *testing.T) {
	event := []Event{}
	event = append(event, NewEvent(FloatEvent{"foo", 10}).New())
	event = append(event, NewEvent(FloatEvent{"foo", 20}).New())

	f := MedianFloat{"Value"}
	result := f.Apply(event)

	var test = []struct {
		index  int
		median float64
	}{
		{0, 15.0},
		{1, 15.0},
	}

	for _, tt := range test {
		if result[tt.index].Record["median(Value)"] != tt.median {
			t.Error(result)
		}
	}
}

func TestMedianFloatOdd(t *testing.T) {
	event := []Event{}
	event = append(event, NewEvent(FloatEvent{"foo", 10}).New())
	event = append(event, NewEvent(FloatEvent{"foo", 20}).New())
	event = append(event, NewEvent(FloatEvent{"foo", 30}).New())

	f := MedianFloat{"Value"}
	result := f.Apply(event)

	var test = []struct {
		index  int
		median float64
	}{
		{0, 20},
		{1, 20},
	}

	for _, tt := range test {
		if result[tt.index].Record["median(Value)"] != tt.median {
			t.Error(result)
		}
	}
}
