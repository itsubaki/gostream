package gocep

import (
	"fmt"
	"strings"
	"sync"
	"testing"
)

func TestStreamConcurrency(t *testing.T) {
	s := NewStream()
	defer s.Close()

	n := 4
	for i := 0; i < n; i++ {
		s.SetWindow(NewIdentityWindow())
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
		newest := Newest(<-s.Output()).Underlying.(string)
		if !strings.HasPrefix(newest, "test") {
			t.Errorf("failed %s", newest)
		}
	}
}
