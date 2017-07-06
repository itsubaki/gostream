package gocep

import (
	"testing"
	"time"
)

func BenchmarkLengthWindowAverageMap(b *testing.B) {
	w := NewLengthWindow(128, 1024)
	defer w.Close()

	w.Selector(EqualsType{MapEvent{}})
	w.Function(AverageMapInt{"Map", "Value"})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[string]interface{})
		m["Value"] = i
		w.Update(NewEvent(MapEvent{"foobar", m}))
	}
}

func BenchmarkLengthWindowAverageInt(b *testing.B) {
	w := NewLengthWindow(128, 1024)
	defer w.Close()

	w.Selector(EqualsType{IntEvent{}})
	w.Function(AverageInt{"Value"})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Update(NewEvent(IntEvent{"foobar", i}))
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
		w.Update(NewEvent(MapEvent{"foobar", m}))
	}
}

func BenchmarkLengthWindowLargerThanInt(b *testing.B) {
	w := NewLengthWindow(128, 1024)
	defer w.Close()

	w.Selector(EqualsType{IntEvent{}})
	w.Selector(LargerThanInt{"Value", 100})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Update(NewEvent(IntEvent{"foobar", i}))
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
		w.Update(NewEvent(MapEvent{"foobar", m}))
	}
}

func BenchmarkLengthWindowSortInt(b *testing.B) {
	w := NewLengthWindow(128, 1024)
	defer w.Close()

	w.Selector(EqualsType{IntEvent{}})
	w.View(SortInt{"Value", false})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Update(NewEvent(IntEvent{"foobar", i}))
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
		w.Update(NewEvent(MapEvent{"foobar", m}))
	}
}

func BenchmarkLengthWindowSortReverseInt(b *testing.B) {
	w := NewLengthWindow(128, 1024)
	defer w.Close()

	w.Selector(EqualsType{IntEvent{}})
	w.View(SortInt{"Value", true})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Update(NewEvent(IntEvent{"foobar", i}))
	}
}

func TestLengthWindow(t *testing.T) {

	w := NewLengthWindow(2, 1024)
	defer w.Close()

	w.Selector(EqualsType{IntEvent{}})
	w.Selector(LargerThanInt{"Value", 1})
	w.Function(Count{})
	w.View(SortInt{"Value", true})

	event := []Event{}
	for i := 0; i < 10; i++ {
		event = w.Update(NewEvent(IntEvent{"foo", i}))
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
	w.Function(Count{})
	w.Function(AverageMapInt{"Map", "Value"})
	w.View(SortMapInt{"Map", "Value", true})

	event := []Event{}
	for i := 0; i < 10; i++ {
		m := make(map[string]interface{})
		m["Value"] = i
		event = w.Update(NewEvent(MapEvent{"name", m}))
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
	w.Listen(NewEvent("").New())

}

func TestLengthBatchWindow(t *testing.T) {

	w := NewLengthBatchWindow(2, 1024)
	defer w.Close()

	w.Selector(EqualsType{IntEvent{}})

	event := []Event{}
	for i := 0; i < 10; i++ {
		event = w.Update(NewEvent(IntEvent{"foo", i}))
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
		event = w.Update(NewEvent(IntEvent{"foo", i}))
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
		event = w.Update(NewEvent(IntEvent{"foo", i}))
	}

	if len(event) == 0 {
		t.Error(event)
	}
}

func TestTimeBatchWindow10ms(t *testing.T) {
	w := NewTimeBatchWindow(4*time.Millisecond, 1024)
	defer w.Close()

	for i := 0; i < 10; i++ {
		w.Update(NewEvent(IntEvent{"foo", i}))
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
	w.Function(AverageMapInt{"Map", "Value"})
	event := w.Update(NewEvent(IntEvent{"foobar", 10}))
	if len(event) != 0 {
		t.Error(event)
	}
}

func TestInsertInto(t *testing.T) {
	w := NewLengthWindow(16, 32)
	defer w.Close()
	w.Function(CastMapStringToInt{"Map", "piyo"})

	m := make(map[string]interface{})
	m["piyo"] = "123"
	e := w.Update(NewEvent(MapEvent{"foobar", m}))

	w2 := NewLengthWindow(16, 32)
	defer w2.Close()
	w2.Function(SumMapInt{"Map", "cast(Map:piyo)"})
	e2 := w2.Update(NewEvent(MapEvent{"foobar", e[0].Record}))

	if e2[0].RecordIntValue("sum(Map:cast(Map:piyo))") != 123 {
		t.Error(e2)
	}
}
