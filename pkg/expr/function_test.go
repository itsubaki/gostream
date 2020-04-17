package expr

import (
	"testing"

	"github.com/itsubaki/gostream/pkg/event"
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
	events := f.Apply(event.List(IntEvent{"foo", 10}))

	if events[0].RecordString("Name") != "foo" {
		t.Error(events)
	}

	if events[0].RecordInt("Value") != 10 {
		t.Error(events)
	}
}

func TestSelectString(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	f := SelectString{"Name", "Name"}
	events := f.Apply(event.List(IntEvent{"foo", 10}))

	if events[0].RecordString("Name") != "foo" {
		t.Error(events)
	}
}

func TestSelectBool(t *testing.T) {
	type BoolEvent struct {
		Value bool
	}

	f := SelectBool{"Value", "Value"}

	events := event.List()

	events = append(events, event.New(BoolEvent{false}))
	events = f.Apply(events)

	if events[0].RecordBool("Value") {
		t.Error(events)
	}
}

func TestSelectInt(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	f := SelectInt{"Value", "Value"}

	events := event.List()

	events = append(events, event.New(IntEvent{"Name", 10}))
	events = f.Apply(events)

	if events[0].RecordInt("Value") != 10 {
		t.Error(events)
	}
}

func TestSelectFloat(t *testing.T) {
	type FloatEvent struct {
		Name  string
		Value float64
	}

	f := SelectFloat{"Value", "Value"}

	events := event.List()

	events = append(events, event.New(FloatEvent{"Name", 10.0}))
	events = f.Apply(events)

	if events[0].RecordFloat("Value") != 10.0 {
		t.Error(events)
	}
}

func TestSumInt(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	f := SumInt{"Value", "sum(Value)"}

	events := event.List()

	events = append(events, event.New(IntEvent{"foo", 30}))
	events = f.Apply(events)

	events = append(events, event.New(IntEvent{"foo", 30}))
	events = f.Apply(events)

	var test = []struct {
		index int
		sum   int
	}{
		{0, 30},
		{1, 60},
	}

	for _, tt := range test {
		if events[tt.index].Record["sum(Value)"] != tt.sum {
			t.Error(events)
		}
	}
}

func TestSumFloat(t *testing.T) {
	type FloatEvent struct {
		Name  string
		Value float64
	}

	f := SumFloat{"Value", "sum(Value)"}

	events := event.List()

	events = append(events, event.New(FloatEvent{"foo", 30}))
	events = f.Apply(events)

	events = append(events, event.New(FloatEvent{"foo", 30}))
	events = f.Apply(events)

	var test = []struct {
		index int
		sum   float64
	}{
		{0, 30.0},
		{1, 60.0},
	}

	for _, tt := range test {
		if events[tt.index].Record["sum(Value)"] != tt.sum {
			t.Error(events)
		}
	}
}

func TestAverageInt(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	f := AverageInt{"Value", "avg(Value)"}

	events := event.List()

	events = append(events, event.New(IntEvent{"foo", 10}))
	events = f.Apply(events)

	events = append(events, event.New(IntEvent{"foo", 20}))
	events = f.Apply(events)

	var test = []struct {
		index int
		avg   float64
	}{
		{0, 10},
		{1, 15},
	}

	for _, tt := range test {
		if events[tt.index].Record["avg(Value)"] != tt.avg {
			t.Errorf("%v %v\n", events[tt.index].Record["avg(Value)"], tt.avg)
		}
	}
}

func TestAverageFloat(t *testing.T) {
	type FloatEvent struct {
		Name  string
		Value float64
	}

	f := AverageFloat{"Value", "avg(Value)"}

	events := event.List()

	events = append(events, event.New(FloatEvent{"foo", 10}))
	events = f.Apply(events)

	events = append(events, event.New(FloatEvent{"foo", 20}))
	events = f.Apply(events)

	var test = []struct {
		index int
		avg   float64
	}{
		{0, 10.0},
		{1, 15.0},
	}

	for _, tt := range test {
		if events[tt.index].Record["avg(Value)"] != tt.avg {
			t.Errorf("%v %v\n", events[tt.index].Record["avg(Value)"], tt.avg)
		}
	}
}

func TestMaxInt(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	f := MaxInt{"Value", "max(Value)"}

	events := event.List()

	events = append(events, event.New(IntEvent{"foo", 10}))
	events = f.Apply(events)

	events = append(events, event.New(IntEvent{"foo", 20}))
	events = f.Apply(events)

	var test = []struct {
		index int
		max   int
	}{
		{0, 10},
		{1, 20},
	}

	for _, tt := range test {
		if events[tt.index].Record["max(Value)"] != tt.max {
			t.Error(events)
		}
	}
}

func TestMaxFloat(t *testing.T) {
	type FloatEvent struct {
		Name  string
		Value float64
	}

	f := MaxFloat{"Value", "max(Value)"}

	events := event.List()

	events = append(events, event.New(FloatEvent{"foo", 10}))
	events = f.Apply(events)

	events = append(events, event.New(FloatEvent{"foo", 20}))
	events = f.Apply(events)

	var test = []struct {
		index int
		max   float64
	}{
		{0, 10.0},
		{1, 20.0},
	}

	for _, tt := range test {
		if events[tt.index].Record["max(Value)"] != tt.max {
			t.Error(events)
		}
	}
}

func TestMinInt(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	f := MinInt{"Value", "min(Value)"}

	events := event.List()

	events = append(events, event.New(IntEvent{"foo", 10}))
	events = f.Apply(events)

	events = append(events, event.New(IntEvent{"foo", 20}))
	events = f.Apply(events)

	var test = []struct {
		index int
		min   int
	}{
		{0, 10},
		{1, 10},
	}

	for _, tt := range test {
		if events[tt.index].Record["min(Value)"] != tt.min {
			t.Error(events)
		}
	}
}

func TestMinFloat(t *testing.T) {
	type FloatEvent struct {
		Name  string
		Value float64
	}

	f := MinFloat{"Value", "min(Value)"}

	events := event.List()

	events = append(events, event.New(FloatEvent{"foo", 10}))
	events = f.Apply(events)

	events = append(events, event.New(FloatEvent{"foo", 20}))
	events = f.Apply(events)

	var test = []struct {
		index int
		max   float64
	}{
		{0, 10.0},
		{1, 10.0},
	}

	for _, tt := range test {
		if events[tt.index].Record["min(Value)"] != tt.max {
			t.Error(events)
		}
	}
}

func TestMedianIntEvent(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	f := MedianInt{"Value", "med(Value)"}

	events := event.List()

	events = append(events, event.New(IntEvent{"foo", 10}))
	events = f.Apply(events)

	events = append(events, event.New(IntEvent{"foo", 20}))
	events = f.Apply(events)

	var test = []struct {
		index  int
		median float64
	}{
		{0, 10.0},
		{1, 15.0},
	}

	for _, tt := range test {
		if events[tt.index].Record["med(Value)"] != tt.median {
			t.Error(events)
		}
	}
}

func TestMedianIntOdd(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	f := MedianInt{"Value", "med(Value)"}

	events := event.List()

	events = append(events, event.New(IntEvent{"foo", 10}))
	events = f.Apply(events)

	events = append(events, event.New(IntEvent{"foo", 20}))
	events = f.Apply(events)

	events = append(events, event.New(IntEvent{"foo", 30}))
	events = f.Apply(events)

	var test = []struct {
		index  int
		median float64
	}{
		{0, 10.0},
		{1, 15.0},
		{2, 20.0},
	}

	for _, tt := range test {
		if events[tt.index].Record["med(Value)"] != tt.median {
			t.Error(events)
		}
	}
}

func TestMedianFloatEven(t *testing.T) {
	type FloatEvent struct {
		Name  string
		Value float64
	}

	f := MedianFloat{"Value", "med(Value)"}

	events := event.List()

	events = append(events, event.New(FloatEvent{"foo", 10}))
	events = f.Apply(events)

	events = append(events, event.New(FloatEvent{"foo", 20}))
	events = f.Apply(events)

	var test = []struct {
		index  int
		median float64
	}{
		{0, 10.0},
		{1, 15.0},
	}

	for _, tt := range test {
		if events[tt.index].Record["med(Value)"] != tt.median {
			t.Error(events)
		}
	}
}

func TestMedianFloatOdd(t *testing.T) {
	type FloatEvent struct {
		Name  string
		Value float64
	}

	f := MedianFloat{"Value", "med(Value)"}

	events := event.List()

	events = append(events, event.New(FloatEvent{"foo", 10}))
	events = f.Apply(events)

	events = append(events, event.New(FloatEvent{"foo", 20}))
	events = f.Apply(events)

	events = append(events, event.New(FloatEvent{"foo", 30}))
	events = f.Apply(events)

	var test = []struct {
		index  int
		median float64
	}{
		{0, 10},
		{1, 15},
		{2, 20},
	}

	for _, tt := range test {
		if events[tt.index].Record["med(Value)"] != tt.median {
			t.Error(events)
		}
	}
}

func TestCastStringToInt(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	f := CastStringToInt{"Name", "cast(Name)"}

	events := event.List(IntEvent{"123", 10})
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

	f := CastStringToFloat{"Name", "cast(Name)"}

	events := event.List(IntEvent{"12.3", 10})
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

	f := CastStringToBool{"Name", "cast(Name)"}

	events := event.List(IntEvent{"false", 10})
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
