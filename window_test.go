package gocep

import (
	"testing"
	"time"
)

func BenchmarkLengthWindowAverageMap(b *testing.B) {
	w := NewLengthWindow(128, 1024)
	defer w.Close()

	w.Selector(EqualsType{MapEvent{}})
	w.Function(AverageMapInt{"Map", "Value", "avg(Value)"})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[string]interface{})
		m["Value"] = i
		w.Update(MapEvent{"foobar", m})
	}
}

func BenchmarkLengthWindowAverageInt(b *testing.B) {
	w := NewLengthWindow(128, 1024)
	defer w.Close()

	w.Selector(EqualsType{IntEvent{}})
	w.Function(AverageInt{"Value", "avg(Value)"})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Update(IntEvent{"foobar", i})
	}
}

func BenchmarkLengthWindowLargerThanMap(b *testing.B) {
	w := NewLengthWindow(128, 1024)
	defer w.Close()

	w.Selector(EqualsType{MapEvent{}})
	w.Selector(LargerThanMapInt{"Map", "Value", 100})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[string]interface{})
		m["Value"] = i
		w.Update(MapEvent{"foobar", m})
	}
}

func BenchmarkLengthWindowLargerThanInt(b *testing.B) {
	w := NewLengthWindow(128, 1024)
	defer w.Close()

	w.Selector(EqualsType{IntEvent{}})
	w.Selector(LargerThanInt{"Value", 100})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Update(IntEvent{"foobar", i})
	}
}

func BenchmarkLengthWindowSortMap(b *testing.B) {
	w := NewLengthWindow(128, 1024)
	defer w.Close()

	w.Selector(EqualsType{MapEvent{}})
	w.View(SortMapInt{"Map", "Value", false})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[string]interface{})
		m["Value"] = i
		w.Update(MapEvent{"foobar", m})
	}
}

func BenchmarkLengthWindowSortInt(b *testing.B) {
	w := NewLengthWindow(128, 1024)
	defer w.Close()

	w.Selector(EqualsType{IntEvent{}})
	w.View(SortInt{"Value", false})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Update(IntEvent{"foobar", i})
	}
}

func BenchmarkLengthWindowSortReverseMap(b *testing.B) {
	w := NewLengthWindow(128, 1024)
	defer w.Close()

	w.Selector(EqualsType{MapEvent{}})
	w.View(SortMapInt{"Map", "Value", true})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[string]interface{})
		m["Value"] = i
		w.Update(MapEvent{"foobar", m})
	}
}

func BenchmarkLengthWindowSortReverseInt(b *testing.B) {
	w := NewLengthWindow(128, 1024)
	defer w.Close()

	w.Selector(EqualsType{IntEvent{}})
	w.View(SortInt{"Value", true})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Update(IntEvent{"foobar", i})
	}
}

func TestLengthWindow(t *testing.T) {

	w := NewLengthWindow(2, 1024)
	defer w.Close()

	w.Selector(EqualsType{IntEvent{}})
	w.Selector(LargerThanInt{"Value", 1})
	w.Function(Count{"count"})
	w.View(SortInt{"Value", true})

	event := []Event{}
	for i := 0; i < 10; i++ {
		event = w.Update(IntEvent{"foo", i})
	}

	var test = []struct {
		index int
		count int
		value int
	}{
		{0, 2, 9},
		{1, 2, 8},
	}

	for _, tt := range test {
		if event[tt.index].Record["count"] != tt.count {
			t.Error(event)
		}
		if event[tt.index].IntValue("Value") != tt.value {
			t.Error(event)
		}
	}
}

func TestLengthWindowMap(t *testing.T) {

	w := NewLengthWindow(2, 1024)
	defer w.Close()

	w.Selector(EqualsType{MapEvent{}})
	w.Selector(LargerThanMapInt{"Map", "Value", 1})
	w.Function(Count{"count"})
	w.Function(AverageMapInt{"Map", "Value", "avg(Map:Value)"})
	w.View(SortMapInt{"Map", "Value", true})

	event := []Event{}
	for i := 0; i < 10; i++ {
		m := make(map[string]interface{})
		m["Value"] = i
		event = w.Update(MapEvent{"name", m})
	}

	var test = []struct {
		index int
		count int
		value int
		avg   float64
	}{
		{0, 2, 9, 8.5},
		{1, 2, 8, 8.5},
	}

	for _, tt := range test {
		if event[tt.index].Record["count"] != tt.count {
			t.Error(event)
		}
		if event[tt.index].MapIntValue("Map", "Value") != tt.value {
			t.Error(event)
		}
		if event[tt.index].Record["avg(Map:Value)"] != tt.avg {
			t.Error(event)
		}
	}
}

func TestLengthWindowListen(t *testing.T) {

	w := NewLengthWindow(2, 1024)
	defer w.Close()

	w.Selector(EqualsType{IntEvent{}})
	w.Listen("")

}

func TestLengthBatchWindow(t *testing.T) {

	w := NewLengthBatchWindow(2, 1024)
	defer w.Close()

	w.Selector(EqualsType{IntEvent{}})

	event := []Event{}
	for i := 0; i < 10; i++ {
		event = w.Update(IntEvent{"foo", i})
	}

	if event[0].IntValue("Value") != 8 {
		t.Error(event)
	}

	if event[1].IntValue("Value") != 9 {
		t.Error(event)
	}

}

func TestTimeWindow0ms(t *testing.T) {
	w := NewTimeWindow(0*time.Millisecond, 1024)
	defer w.Close()

	event := []Event{}
	for i := 0; i < 10; i++ {
		event = w.Update(IntEvent{"foo", i})
	}

	if len(event) != 0 {
		t.Error(event)
	}
}

func TestTimeWindow10ms(t *testing.T) {
	w := NewTimeWindow(1*time.Millisecond, 1024)
	defer w.Close()

	event := []Event{}
	for i := 0; i < 10; i++ {
		event = w.Update(IntEvent{"foo", i})
	}

	if len(event) == 0 {
		t.Error(event)
	}
}

func TestTimeBatchWindow10ms(t *testing.T) {
	w := NewTimeBatchWindow(4*time.Millisecond, 1024)
	defer w.Close()

	for i := 0; i < 10; i++ {
		w.Update(IntEvent{"foo", i})
	}
}

func TestLengthWindowPanic(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()

	w := NewLengthWindow(10, 16)
	defer w.Close()

	w.Selector(EqualsType{IntEvent{}})
	w.Function(AverageMapInt{"Map", "Value", "avg(Map:Value)"})
	event := w.Update(IntEvent{"foobar", 10})
	if len(event) != 0 {
		t.Error(event)
	}
}

func TestInsertIntoIntEvent(t *testing.T) {
	w := NewLengthWindow(16, 32)
	defer w.Close()
	w.Function(CastStringToInt{"Name", "c(name)"})
	e := w.Update(IntEvent{"123", 123})
	cname := e[0].RecordIntValue("c(name)")

	w2 := NewLengthWindow(16, 32)
	defer w2.Close()
	w2.Function(SumInt{"Value", "sum(name)"})
	e2 := w2.Update(IntEvent{"foobar", cname})

	if e2[0].RecordIntValue("sum(name)") != 123 {
		t.Error(e2)
	}
}

func TestInsertIntoMapEvent(t *testing.T) {
	w := NewLengthWindow(16, 32)
	defer w.Close()
	w.Function(CastMapStringToInt{"Map", "str", "cast(str)"})

	m := make(map[string]interface{})
	m["str"] = "123"
	cast := w.Update(MapEvent{"foo", m})

	w2 := NewLengthWindow(16, 32)
	defer w2.Close()
	w2.Function(SumMapInt{"Map", "cast(str)", "sum(str)"})

	for _, e := range cast {
		e := w2.Update(MapEvent{"foo", e.Record})

		for _, r := range e {
			if r.RecordIntValue("sum(str)") != 123 {
				t.Error(e)
			}
		}
	}
}
