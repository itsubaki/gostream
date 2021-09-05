package function_test

import (
	"testing"

	"github.com/itsubaki/gostream/pkg/event"
	"github.com/itsubaki/gostream/pkg/function"
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

	f := function.SumInt{"Value", "sun(Value)"}
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

	f := function.SumInt{"Value", "sun(Value)"}
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

	f := function.SumInt{"Value", "sun(Value)"}
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

	f := function.AverageInt{"Value", "avg(Value)"}
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

	f := function.AverageInt{"Value", "avg(Value)"}
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

	f := function.SelectAll{}
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

	f := function.SelectString{"Name", "Name"}
	events := f.Apply(event.List(IntEvent{"foo", 10}))

	if events[0].RecordString("Name") != "foo" {
		t.Error(events)
	}
}

func TestSelectBool(t *testing.T) {
	type BoolEvent struct {
		Value bool
	}

	f := function.SelectBool{"Value", "Value"}

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

	f := function.SelectInt{"Value", "Value"}

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

	f := function.SelectFloat{"Value", "Value"}

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

	f := function.SumInt{"Value", "sum(Value)"}

	events := event.List()
	events = append(events, event.New(IntEvent{"foo", 30}))
	events = f.Apply(events)
	events = append(events, event.New(IntEvent{"foo", 30}))
	events = f.Apply(events)

	var cases = []struct {
		in   int
		want int
	}{
		{0, 30},
		{1, 60},
	}

	for _, c := range cases {
		if events[c.in].Record["sum(Value)"] != c.want {
			t.Fail()
		}
	}
}

func TestSumFloat(t *testing.T) {
	type FloatEvent struct {
		Name  string
		Value float64
	}

	f := function.SumFloat{"Value", "sum(Value)"}

	events := event.List()
	events = append(events, event.New(FloatEvent{"foo", 30}))
	events = f.Apply(events)
	events = append(events, event.New(FloatEvent{"foo", 30}))
	events = f.Apply(events)

	var cases = []struct {
		in   int
		want float64
	}{
		{0, 30.0},
		{1, 60.0},
	}

	for _, c := range cases {
		if events[c.in].Record["sum(Value)"] != c.want {
			t.Fail()
		}
	}
}

func TestAverageInt(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	f := function.AverageInt{"Value", "avg(Value)"}

	events := event.List()
	events = append(events, event.New(IntEvent{"foo", 10}))
	events = f.Apply(events)
	events = append(events, event.New(IntEvent{"foo", 20}))
	events = f.Apply(events)

	var cases = []struct {
		in   int
		want float64
	}{
		{0, 10},
		{1, 15},
	}

	for _, c := range cases {
		if events[c.in].Record["avg(Value)"] != c.want {
			t.Fail()
		}
	}
}

func TestAverageFloat(t *testing.T) {
	type FloatEvent struct {
		Name  string
		Value float64
	}

	f := function.AverageFloat{"Value", "avg(Value)"}

	events := event.List()
	events = append(events, event.New(FloatEvent{"foo", 10}))
	events = f.Apply(events)
	events = append(events, event.New(FloatEvent{"foo", 20}))
	events = f.Apply(events)

	var cases = []struct {
		in   int
		want float64
	}{
		{0, 10.0},
		{1, 15.0},
	}

	for _, c := range cases {
		if events[c.in].Record["avg(Value)"] != c.want {
			t.Fail()
		}
	}
}

func TestMaxInt(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	f := function.MaxInt{"Value", "max(Value)"}

	events := event.List()
	events = append(events, event.New(IntEvent{"foo", 10}))
	events = f.Apply(events)
	events = append(events, event.New(IntEvent{"foo", 20}))
	events = f.Apply(events)

	var cases = []struct {
		in   int
		want int
	}{
		{0, 10},
		{1, 20},
	}

	for _, c := range cases {
		if events[c.in].Record["max(Value)"] != c.want {
			t.Fail()
		}
	}
}

func TestMaxFloat(t *testing.T) {
	type FloatEvent struct {
		Name  string
		Value float64
	}

	f := function.MaxFloat{"Value", "max(Value)"}

	events := event.List()
	events = append(events, event.New(FloatEvent{"foo", 10}))
	events = f.Apply(events)
	events = append(events, event.New(FloatEvent{"foo", 20}))
	events = f.Apply(events)

	var cases = []struct {
		in   int
		want float64
	}{
		{0, 10.0},
		{1, 20.0},
	}

	for _, c := range cases {
		if events[c.in].Record["max(Value)"] != c.want {
			t.Fail()
		}
	}
}

func TestMinInt(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	f := function.MinInt{"Value", "min(Value)"}

	events := event.List()
	events = append(events, event.New(IntEvent{"foo", 10}))
	events = f.Apply(events)
	events = append(events, event.New(IntEvent{"foo", 20}))
	events = f.Apply(events)

	var cases = []struct {
		in   int
		want int
	}{
		{0, 10},
		{1, 10},
	}

	for _, c := range cases {
		if events[c.in].Record["min(Value)"] != c.want {
			t.Fail()
		}
	}
}

func TestMinFloat(t *testing.T) {
	type FloatEvent struct {
		Name  string
		Value float64
	}

	f := function.MinFloat{"Value", "min(Value)"}

	events := event.List()
	events = append(events, event.New(FloatEvent{"foo", 10}))
	events = f.Apply(events)
	events = append(events, event.New(FloatEvent{"foo", 20}))
	events = f.Apply(events)

	var cases = []struct {
		in   int
		want float64
	}{
		{0, 10.0},
		{1, 10.0},
	}

	for _, c := range cases {
		if events[c.in].Record["min(Value)"] != c.want {
			t.Fail()
		}
	}
}

func TestMedianIntEvent(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	f := function.MedianInt{"Value", "med(Value)"}

	events := event.List()
	events = append(events, event.New(IntEvent{"foo", 10}))
	events = f.Apply(events)
	events = append(events, event.New(IntEvent{"foo", 20}))
	events = f.Apply(events)

	var cases = []struct {
		in   int
		want float64
	}{
		{0, 10.0},
		{1, 15.0},
	}

	for _, c := range cases {
		if events[c.in].Record["med(Value)"] != c.want {
			t.Fail()
		}
	}
}

func TestMedianIntOdd(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	f := function.MedianInt{"Value", "med(Value)"}

	events := event.List()
	events = append(events, event.New(IntEvent{"foo", 10}))
	events = f.Apply(events)
	events = append(events, event.New(IntEvent{"foo", 20}))
	events = f.Apply(events)
	events = append(events, event.New(IntEvent{"foo", 30}))
	events = f.Apply(events)

	var cases = []struct {
		in   int
		want float64
	}{
		{0, 10.0},
		{1, 15.0},
		{2, 20.0},
	}

	for _, c := range cases {
		if events[c.in].Record["med(Value)"] != c.want {
			t.Fail()
		}
	}
}

func TestMedianFloatEven(t *testing.T) {
	type FloatEvent struct {
		Name  string
		Value float64
	}

	f := function.MedianFloat{"Value", "med(Value)"}

	events := event.List()
	events = append(events, event.New(FloatEvent{"foo", 10}))
	events = f.Apply(events)
	events = append(events, event.New(FloatEvent{"foo", 20}))
	events = f.Apply(events)

	var cases = []struct {
		in   int
		want float64
	}{
		{0, 10.0},
		{1, 15.0},
	}

	for _, c := range cases {
		if events[c.in].Record["med(Value)"] != c.want {
			t.Fail()
		}
	}
}

func TestMedianFloatOdd(t *testing.T) {
	type FloatEvent struct {
		Name  string
		Value float64
	}

	f := function.MedianFloat{"Value", "med(Value)"}

	events := event.List()
	events = append(events, event.New(FloatEvent{"foo", 10}))
	events = f.Apply(events)
	events = append(events, event.New(FloatEvent{"foo", 20}))
	events = f.Apply(events)
	events = append(events, event.New(FloatEvent{"foo", 30}))
	events = f.Apply(events)

	var cases = []struct {
		in   int
		want float64
	}{
		{0, 10},
		{1, 15},
		{2, 20},
	}

	for _, c := range cases {
		if events[c.in].Record["med(Value)"] != c.want {
			t.Fail()
		}
	}
}

func TestCastStringToInt(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	f := function.CastStringToInt{"Name", "cast(Name)"}

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

	f := function.CastStringToFloat{"Name", "cast(Name)"}

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

	f := function.CastStringToBool{"Name", "cast(Name)"}

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

	var cases = []struct {
		in   int
		want int
	}{
		{30, 3},
		{31, 0},
	}

	for _, c := range cases {
		f := function.FuncEqualsInt{
			function.SumInt{"Value", "sum(Value)"},
			"sum(Value)",
			c.in,
		}

		got := f.Apply(events)
		if len(got) != c.want {
			t.Fail()
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

	var cases = []struct {
		in   int
		want int
	}{
		{30, 0},
		{29, 3},
	}

	for _, c := range cases {
		f := function.FuncLargerThanInt{
			function.SumInt{"Value", "sum(Value)"},
			"sum(Value)",
			c.in,
		}

		got := f.Apply(events)
		if len(got) != c.want {
			t.Fail()
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

	var cases = []struct {
		in   int
		want int
	}{
		{31, 3},
		{30, 0},
	}

	for _, c := range cases {
		f := function.FuncLessThanInt{
			function.SumInt{"Value", "sum(Value)"},
			"sum(Value)",
			c.in,
		}

		got := f.Apply(events)
		if len(got) != c.want {
			t.Fail()
		}
	}
}

func TestDistinctString(t *testing.T) {
	type StringEvent struct {
		Name string
	}

	events := event.List(
		StringEvent{"foo"},
		StringEvent{"bar"},
		StringEvent{"foo"},
	)

	f := function.DistinctString{
		Name: "Name",
		As:   "distinct",
	}

	result := f.Apply(events)
	newest := event.Newest(result)
	distinct := newest.RecordStringSlice("distinct")
	if distinct[0] != "foo" {
		t.Error(distinct)
	}
	if distinct[1] != "bar" {
		t.Error(distinct)
	}
}
