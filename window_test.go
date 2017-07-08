package gocep

import (
	"testing"
	"time"
)

func BenchmarkLengthWindowAverageMap(b *testing.B) {
	w := NewLengthWindow(128, 1024)
	defer w.Close()

	w.Selector(EqualsType{MapEvent{}})
	w.Function(AverageMapInt{"Record", "Value", "avg(Value)"})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[string]interface{})
		m["Value"] = i
		w.Update(MapEvent{m})
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
	w.Selector(LargerThanMapInt{"Record", "Value", 100})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[string]interface{})
		m["Value"] = i
		w.Update(MapEvent{m})
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
	w.View(SortMapInt{"Record", "Value", false})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[string]interface{})
		m["Value"] = i
		w.Update(MapEvent{m})
	}
}

func BenchmarkLengthWindowSortInt(b *testing.B) {
	w := NewLengthWindow(128, 1024)
	defer w.Close()

	w.Selector(EqualsType{IntEvent{}})
	w.View(OrderByInt{"Value", false})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Update(IntEvent{"foobar", i})
	}
}

func BenchmarkLengthWindowSortReverseMap(b *testing.B) {
	w := NewLengthWindow(128, 1024)
	defer w.Close()

	w.Selector(EqualsType{MapEvent{}})
	w.View(SortMapInt{"Record", "Value", true})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[string]interface{})
		m["Value"] = i
		w.Update(MapEvent{m})
	}
}

func BenchmarkLengthWindowSortReverseInt(b *testing.B) {
	w := NewLengthWindow(128, 1024)
	defer w.Close()

	w.Selector(EqualsType{IntEvent{}})
	w.View(OrderByInt{"Value", true})

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
	w.View(OrderByInt{"Value", true})

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
		if event[tt.index].Int("Value") != tt.value {
			t.Error(event)
		}
	}
}

func TestLengthWindowMap(t *testing.T) {

	w := NewLengthWindow(2, 1024)
	defer w.Close()

	w.Selector(EqualsType{MapEvent{}})
	w.Selector(LargerThanMapInt{"Record", "Value", 1})
	w.Function(Count{"count"})
	w.Function(AverageMapInt{"Record", "Value", "avg(Record:Value)"})
	w.View(SortMapInt{"Record", "Value", true})

	event := []Event{}
	for i := 0; i < 10; i++ {
		m := make(map[string]interface{})
		m["Value"] = i
		event = w.Update(MapEvent{m})
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
		if event[tt.index].MapInt("Record", "Value") != tt.value {
			t.Error(event)
		}
		if event[tt.index].Record["avg(Record:Value)"] != tt.avg {
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

	if event[0].Int("Value") != 8 {
		t.Error(event)
	}

	if event[1].Int("Value") != 9 {
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
	// IntEvent and Map Function -> panic!!
	w.Function(AverageMapInt{"Record", "Value", "avg(Record:Value)"})
	event := w.Update(IntEvent{"foobar", 10})
	if len(event) != 0 {
		t.Error(event)
	}
}
