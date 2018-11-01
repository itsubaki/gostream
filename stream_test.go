package gocep

import (
	"testing"
)

func BenchmarkLengthWindowSumIntStream(b *testing.B) {
	s := NewStream(b.N)
	defer s.Close()

	w := NewLengthWindow(b.N)
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
