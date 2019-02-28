package function

import (
	"testing"

	"github.com/itsubaki/gocep/pkg/event"
)

func BenchmarkSumRaw128(b *testing.B) {
	type Raw struct {
		Value int
		Sum   int
	}

	events := []Raw{}
	for i := 0; i < 128; i++ {
		events = append(events, Raw{Value: i})
	}

	f := func() int {
		sum := 0
		for _, e := range events {
			sum = sum + e.Value
		}

		for _, e := range events {
			e.Sum = sum
		}

		return sum
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f()
	}

}

func BenchmarkSum128(b *testing.B) {
	type IntEvent struct {
		Name  string
		Value int
	}

	events := event.List()
	for i := 0; i < 128; i++ {
		events = append(events, event.New(IntEvent{"foo", i}))
	}

	f := func() int {
		sum := 0

		// 128 allocs
		for _, e := range events {
			sum = sum + e.Int("Value")
		}

		// 128 allocs
		for _, e := range events {
			e.Record["sum(Value)"] = sum
		}

		return sum
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f()
	}

}

func BenchmarkSumInt(b *testing.B) {
	type IntEvent struct {
		Name  string
		Value int
	}

	events := event.List()
	for i := 0; i < 1; i++ {
		events = append(events, event.New(IntEvent{"foo", i}))
	}

	f := SumInt{"Value", "sun(Value)"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.Apply(events)
	}
}

func BenchmarkSumInt64(b *testing.B) {
	type IntEvent struct {
		Name  string
		Value int
	}

	events := event.List()
	for i := 0; i < 64; i++ {
		events = append(events, event.New(IntEvent{"foo", i}))
	}

	f := SumInt{"Value", "sun(Value)"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.Apply(events)
	}
}

func BenchmarkSumInt128(b *testing.B) {
	type IntEvent struct {
		Name  string
		Value int
	}

	events := event.List()
	for i := 0; i < 128; i++ {
		events = append(events, event.New(IntEvent{"foo", i}))
	}

	f := SumInt{"Value", "sun(Value)"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.Apply(events)
	}
}

func BenchmarkAverageInt(b *testing.B) {
	type IntEvent struct {
		Name  string
		Value int
	}

	events := event.List()
	for i := 0; i < 1; i++ {
		events = append(events, event.New(IntEvent{"foo", i}))
	}

	f := AverageInt{"Value", "avg(Value)"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.Apply(events)
	}
}

func BenchmarkAverageInt128(b *testing.B) {
	type IntEvent struct {
		Name  string
		Value int
	}

	events := event.List()
	for i := 0; i < 128; i++ {
		events = append(events, event.New(IntEvent{"foo", i}))
	}

	f := AverageInt{"Value", "avg(Value)"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.Apply(events)
	}
}

func TestSelectAll(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	f := SelectAll{}
	result := f.Apply(event.List(IntEvent{"foo", 10}))

	if result[0].RecordString("Name") != "foo" {
		t.Error(result)
	}

	if result[0].RecordInt("Value") != 10 {
		t.Error(result)
	}
}

func TestSelectString(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	f := SelectString{"Name", "Name"}
	result := f.Apply(event.List(IntEvent{"foo", 10}))

	if result[0].RecordString("Name") != "foo" {
		t.Error(result)
	}
}

func TestSelectBool(t *testing.T) {
	type BoolEvent struct {
		Value bool
	}

	f := SelectBool{"Value", "Value"}
	result := f.Apply(event.List(BoolEvent{false}))

	if result[0].RecordBool("Value") {
		t.Error(result)
	}
}

func TestSelectInt(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	events := event.List(IntEvent{"Name", 10})
	f := SelectInt{"Value", "Value"}

	result := f.Apply(events)
	if result[0].RecordInt("Value") != 10 {
		t.Error(result)
	}
}

func TestSelectFloat(t *testing.T) {
	type FloatEvent struct {
		Name  string
		Value float64
	}

	events := event.List(FloatEvent{"Name", 10.0})
	f := SelectFloat{"Value", "Value"}

	result := f.Apply(events)
	if result[0].RecordFloat("Value") != 10.0 {
		t.Error(result)
	}
}

func TestSumInt(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	events := event.List(
		IntEvent{"foo", 10},
		IntEvent{"foo", 20},
	)

	f := SumInt{"Value", "sum(Value)"}
	result := f.Apply(events)

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
	type FloatEvent struct {
		Name  string
		Value float64
	}

	events := event.List(
		FloatEvent{"foo", 10.0},
		FloatEvent{"foo", 20.0},
	)

	f := SumFloat{"Value", "sum(Value)"}
	result := f.Apply(events)

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
	type IntEvent struct {
		Name  string
		Value int
	}

	events := event.List(
		IntEvent{"foo", 10},
		IntEvent{"foo", 20},
	)

	f := AverageInt{"Value", "avg(Value)"}
	result := f.Apply(events)

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
	type FloatEvent struct {
		Name  string
		Value float64
	}

	events := event.List(
		FloatEvent{"foo", 10.0},
		FloatEvent{"foo", 20.0},
	)

	f := AverageFloat{"Value", "avg(Value)"}
	result := f.Apply(events)

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
	type IntEvent struct {
		Name  string
		Value int
	}

	events := event.List(
		IntEvent{"foo", 10},
		IntEvent{"foo", 20},
	)

	f := MaxInt{"Value", "max(Value)"}
	result := f.Apply(events)

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
	type FloatEvent struct {
		Name  string
		Value float64
	}

	events := event.List(
		FloatEvent{"foo", 10},
		FloatEvent{"foo", 20},
	)

	f := MaxFloat{"Value", "max(Value)"}
	result := f.Apply(events)

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
	type IntEvent struct {
		Name  string
		Value int
	}

	events := event.List(
		IntEvent{"foo", 10},
		IntEvent{"foo", 20},
	)

	f := MinInt{"Value", "min(Value)"}
	result := f.Apply(events)

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
	type FloatEvent struct {
		Name  string
		Value float64
	}

	events := event.List(
		FloatEvent{"foo", 10},
		FloatEvent{"foo", 20},
	)

	f := MinFloat{"Value", "min(Value)"}
	result := f.Apply(events)

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

func TestMedianIntEvent(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	events := event.List(
		IntEvent{"foo", 10},
		IntEvent{"foo", 20},
	)

	f := MedianInt{"Value", "median(Value)"}
	result := f.Apply(events)

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
	type IntEvent struct {
		Name  string
		Value int
	}

	events := event.List(
		IntEvent{"foo", 10},
		IntEvent{"foo", 20},
		IntEvent{"foo", 30},
	)

	f := MedianInt{"Value", "median(Value)"}
	result := f.Apply(events)

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
	type FloatEvent struct {
		Name  string
		Value float64
	}

	events := event.List(
		FloatEvent{"foo", 10},
		FloatEvent{"foo", 20},
	)

	f := MedianFloat{"Value", "median(Value)"}
	result := f.Apply(events)

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
	type FloatEvent struct {
		Name  string
		Value float64
	}

	events := event.List(
		FloatEvent{"foo", 10},
		FloatEvent{"foo", 20},
		FloatEvent{"foo", 30},
	)

	f := MedianFloat{"Value", "median(Value)"}
	result := f.Apply(events)

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
	type IntEvent struct {
		Name  string
		Value int
	}

	events := event.List(IntEvent{"123", 10})

	f := CastStringToInt{"Name", "cast(Name)"}
	result := f.Apply(events)

	if result[0].RecordInt("cast(Name)") != 123 {
		t.Error(result)
	}
}

func TestCastStringToFloat(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	events := event.List(IntEvent{"12.3", 10})

	f := CastStringToFloat{"Name", "cast(Name)"}
	result := f.Apply(events)

	if result[0].RecordFloat("cast(Name)") != 12.3 {
		t.Error(result)
	}
}

func TestCastStringToBool(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	events := event.List(IntEvent{"false", 10})

	f := CastStringToBool{"Name", "cast(Name)"}
	result := f.Apply(events)

	if result[0].RecordBool("cast(Name)") {
		t.Error(result)
	}
}

func TestFuncEqualsInt(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	events := event.List(
		IntEvent{"foo", 10},
		IntEvent{"foo", 10},
		IntEvent{"foo", 10},
	)

	var test = []struct {
		sum      int
		expected int
	}{
		{30, 3},
		{31, 0},
	}

	for _, tt := range test {
		f := FuncEqualsInt{
			SumInt{"Value", "sum(Value)"},
			"sum(Value)",
			tt.sum,
		}
		result := f.Apply(events)
		if len(result) != tt.expected {
			t.Error(result)
		}
	}
}

func TestFuncLargerThanInt(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	events := event.List(
		IntEvent{"foo", 10},
		IntEvent{"foo", 10},
		IntEvent{"foo", 10},
	)

	var test = []struct {
		sum      int
		expected int
	}{
		{30, 0},
		{29, 3},
	}

	for _, tt := range test {
		f := FuncLargerThanInt{
			SumInt{"Value", "sum(Value)"},
			"sum(Value)",
			tt.sum,
		}
		result := f.Apply(events)
		if len(result) != tt.expected {
			t.Error(result)
		}
	}
}

func TestFuncLessThanInt(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	events := event.List(
		IntEvent{"foo", 10},
		IntEvent{"foo", 10},
		IntEvent{"foo", 10},
	)

	var test = []struct {
		sum      int
		expected int
	}{
		{31, 3},
		{30, 0},
	}

	for _, tt := range test {
		f := FuncLessThanInt{
			SumInt{"Value", "sum(Value)"},
			"sum(Value)",
			tt.sum,
		}
		result := f.Apply(events)
		if len(result) != tt.expected {
			t.Error(result)
		}
	}
}
