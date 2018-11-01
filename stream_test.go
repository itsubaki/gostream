package gocep

import (
	"testing"
)

func BenchmarkLengthWindowNoFunction128Stream(b *testing.B) {
	s := NewStream(b.N)
	defer s.Close()

	w := NewLengthWindow(128)
	w.SetSelector(EqualsType{MapEvent{}})
	s.SetWindow(w)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[string]interface{})
		m["Value"] = i
		s.Input() <- MapEvent{m}
	}

	for i := 0; i < b.N; i++ {
		<-s.Output()
	}

}

func BenchmarkLengthWindowSumIntStream(b *testing.B) {
	s := NewStream(b.N)
	defer s.Close()

	w := NewLengthWindow(1)
	w.SetSelector(EqualsType{IntEvent{}})
	w.SetFunction(SumInt{"Value", "sum(Value)"})
	s.SetWindow(w)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Input() <- IntEvent{"foobar", i}
	}

	for i := 0; i < b.N; i++ {
		<-s.Output()
	}
}

func BenchmarkLengthWindowSumInt64Stream(b *testing.B) {
	s := NewStream(b.N)
	defer s.Close()

	w := NewLengthWindow(64)
	w.SetSelector(EqualsType{IntEvent{}})
	w.SetFunction(SumInt{"Value", "sum(Value)"})
	s.SetWindow(w)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Input() <- IntEvent{"foobar", i}
	}

	for i := 0; i < b.N; i++ {
		<-s.Output()
	}
}

func BenchmarkLengthWindowSumInt128Stream(b *testing.B) {
	s := NewStream(b.N)
	defer s.Close()

	w := NewLengthWindow(128)
	w.SetSelector(EqualsType{IntEvent{}})
	w.SetFunction(SumInt{"Value", "sum(Value)"})
	s.SetWindow(w)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Input() <- IntEvent{"foobar", i}
	}

	for i := 0; i < b.N; i++ {
		<-s.Output()
	}
}

func BenchmarkLengthWindowSumInt256Stream(b *testing.B) {
	s := NewStream(b.N)
	defer s.Close()

	w := NewLengthWindow(256)
	w.SetSelector(EqualsType{IntEvent{}})
	w.SetFunction(SumInt{"Value", "sum(Value)"})
	s.SetWindow(w)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Input() <- IntEvent{"foobar", i}
	}

	for i := 0; i < b.N; i++ {
		<-s.Output()
	}
}

func TestStream(t *testing.T) {
	s := NewStream()
	defer s.Close()

	n := 2
	for i := 0; i < n; i++ {
		s.SetWindow(NewIdentityWindow())
	}

	if len(s.Window()) != n {
		t.Error("failed.")
	}

	s.Input() <- "test"

	for i := 0; i < n; i++ {
		if Oldest(<-s.Output()).Underlying != "test" {
			t.Error("failed")
		}
	}
}
