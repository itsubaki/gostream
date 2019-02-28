package gocep

import (
	"fmt"
	"strings"
	"sync"
	"testing"

	"github.com/itsubaki/gocep/pkg/event"
	"github.com/itsubaki/gocep/pkg/window"
)

func BenchmarkStreamLengthWindwx4(b *testing.B) {
	s := New(b.N)
	defer s.Close()

	n := 4
	for i := 0; i < n; i++ {
		s.AddWindow(window.NewLength(10))
	}

	if len(s.Window()) != n {
		b.Error("failed.")
	}

	loop := b.N
	var wg sync.WaitGroup
	for i := 0; i < loop; i++ {
		wg.Add(1)
		go func(val int) {
			s.Input() <- fmt.Sprintf("test %d", val)
			wg.Done()
		}(i)
	}
	wg.Wait()

	for i := 0; i < n*loop; i++ {
		newest := event.Newest(<-s.Output()).Underlying.(string)
		if !strings.HasPrefix(newest, "test") {
			b.Errorf("failed %s", newest)
		}
	}
}

func TestStreamConcurrency(t *testing.T) {
	s := New()
	defer s.Close()

	n := 4
	for i := 0; i < n; i++ {
		s.AddWindow(window.NewLength(10))
	}

	if len(s.Window()) != n {
		t.Error("failed.")
	}

	loop := 100
	var wg sync.WaitGroup
	for i := 0; i < loop; i++ {
		wg.Add(1)
		go func(val int) {
			s.Input() <- fmt.Sprintf("test %d", val)
			wg.Done()
		}(i)
	}
	wg.Wait()

	for i := 0; i < n*loop; i++ {
		newest := event.Newest(<-s.Output()).Underlying.(string)
		if !strings.HasPrefix(newest, "test") {
			t.Errorf("failed %s", newest)
		}
	}
}
