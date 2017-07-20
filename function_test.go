package gocep

import (
	"testing"
)

func TestSelectAll(t *testing.T) {
	event := []Event{NewEvent(IntEvent{"foo", 10})}
	f := SelectAll{}

	result := f.Apply(event)
	if result[0].RecordString("Name") != "foo" {
		t.Error(result)
	}
	if result[0].RecordInt("Value") != 10 {
		t.Error(result)
	}
}

func TestSelectString(t *testing.T) {
	event := []Event{NewEvent(IntEvent{"foo", 10})}
	f := SelectString{"Name", "Name"}

	result := f.Apply(event)
	if result[0].RecordString("Name") != "foo" {
		t.Error(result)
	}
}

func TestSelectBool(t *testing.T) {
	event := []Event{NewEvent(BoolEvent{false})}
	f := SelectBool{"Value", "Value"}

	result := f.Apply(event)
	if result[0].RecordBool("Value") {
		t.Error(result)
	}
}

func TestSelectInt(t *testing.T) {
	event := []Event{NewEvent(IntEvent{"Name", 10})}
	f := SelectInt{"Value", "Value"}

	result := f.Apply(event)
	if result[0].RecordInt("Value") != 10 {
		t.Error(result)
	}
}

func TestSelectFloat(t *testing.T) {
	event := []Event{NewEvent(FloatEvent{"Name", 10.0})}
	f := SelectFloat{"Value", "Value"}

	result := f.Apply(event)
	if result[0].RecordFloat("Value") != 10.0 {
		t.Error(result)
	}
}

func TestSumInt(t *testing.T) {
	event := []Event{}
	event = append(event, NewEvent(IntEvent{"foo", 10}))
	event = append(event, NewEvent(IntEvent{"foo", 20}))

	f := SumInt{"Value", "sum(Value)"}
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
	event = append(event, NewEvent(FloatEvent{"foo", 10.0}))
	event = append(event, NewEvent(FloatEvent{"foo", 20.0}))

	f := SumFloat{"Value", "sum(Value)"}
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
	event = append(event, NewEvent(IntEvent{"foo", 10}))
	event = append(event, NewEvent(IntEvent{"foo", 20}))

	f := AverageInt{"Value", "avg(Value)"}
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
	event = append(event, NewEvent(FloatEvent{"foo", 10.0}))
	event = append(event, NewEvent(FloatEvent{"foo", 20.0}))

	f := AverageFloat{"Value", "avg(Value)"}
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
	event = append(event, NewEvent(IntEvent{"foo", 10}))
	event = append(event, NewEvent(IntEvent{"foo", 20}))

	f := MaxInt{"Value", "max(Value)"}
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
	event = append(event, NewEvent(FloatEvent{"foo", 10}))
	event = append(event, NewEvent(FloatEvent{"foo", 20}))

	f := MaxFloat{"Value", "max(Value)"}
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
	event = append(event, NewEvent(IntEvent{"foo", 10}))
	event = append(event, NewEvent(IntEvent{"foo", 20}))

	f := MinInt{"Value", "min(Value)"}
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
	event = append(event, NewEvent(FloatEvent{"foo", 10}))
	event = append(event, NewEvent(FloatEvent{"foo", 20}))

	f := MinFloat{"Value", "min(Value)"}
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
	event = append(event, NewEvent(IntEvent{"foo", 10}))
	event = append(event, NewEvent(IntEvent{"foo", 20}))

	f := MedianInt{"Value", "median(Value)"}
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
	event = append(event, NewEvent(IntEvent{"foo", 10}))
	event = append(event, NewEvent(IntEvent{"foo", 20}))
	event = append(event, NewEvent(IntEvent{"foo", 30}))

	f := MedianInt{"Value", "median(Value)"}
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
	event = append(event, NewEvent(FloatEvent{"foo", 10}))
	event = append(event, NewEvent(FloatEvent{"foo", 20}))

	f := MedianFloat{"Value", "median(Value)"}
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
	event = append(event, NewEvent(FloatEvent{"foo", 10}))
	event = append(event, NewEvent(FloatEvent{"foo", 20}))
	event = append(event, NewEvent(FloatEvent{"foo", 30}))

	f := MedianFloat{"Value", "median(Value)"}
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

func TestCastStringToInt(t *testing.T) {
	event := []Event{NewEvent(IntEvent{"123", 10})}

	f := CastStringToInt{"Name", "cast(Name)"}
	result := f.Apply(event)

	if result[0].RecordInt("cast(Name)") != 123 {
		t.Error(result)
	}
}

func TestCastStringToFloat(t *testing.T) {
	event := []Event{NewEvent(IntEvent{"12.3", 10})}

	f := CastStringToFloat{"Name", "cast(Name)"}
	result := f.Apply(event)

	if result[0].RecordFloat("cast(Name)") != 12.3 {
		t.Error(result)
	}
}

func TestCastStringToBool(t *testing.T) {
	event := []Event{NewEvent(IntEvent{"false", 10})}

	f := CastStringToBool{"Name", "cast(Name)"}
	result := f.Apply(event)

	if result[0].RecordBool("cast(Name)") {
		t.Error(result)
	}
}

func TestHavingLargerThanInt(t *testing.T) {
	event := []Event{
		NewEvent(IntEvent{"foo", 10}),
		NewEvent(IntEvent{"foo", 10}),
		NewEvent(IntEvent{"foo", 10}),
	}

	var test = []struct {
		sum      int
		expected int
	}{
		{30, 0},
		{29, 3},
	}

	for _, tt := range test {
		f := HavingLargerThanInt{
			SumInt{"Value", "sum(Value)"},
			"sum(Value)",
			tt.sum,
		}
		result := f.Apply(event)
		if len(result) != tt.expected {
			t.Error(result)
		}
	}
}

func TestHavingLessThanInt(t *testing.T) {
	event := []Event{
		NewEvent(IntEvent{"foo", 10}),
		NewEvent(IntEvent{"foo", 10}),
		NewEvent(IntEvent{"foo", 10}),
	}

	var test = []struct {
		sum      int
		expected int
	}{
		{31, 3},
		{30, 0},
	}

	for _, tt := range test {
		f := HavingLessThanInt{
			SumInt{"Value", "sum(Value)"},
			"sum(Value)",
			tt.sum,
		}
		result := f.Apply(event)
		if len(result) != tt.expected {
			t.Error(result)
		}
	}
}
