package function_test

import (
	"testing"
	"time"

	"github.com/itsubaki/gostream/pkg/event"
	"github.com/itsubaki/gostream/pkg/function"
)

func TestTimeDurationBatch(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	start := time.Now()
	duration := 10 * time.Millisecond
	f := &function.TimeDurationBatch{start, start.Add(duration), duration}

	events := append(event.List(), event.New(IntEvent{"foo", 1}))
	events = f.Apply(events)
	if len(events) != 1 {
		t.Error(events)
	}
	time.Sleep(duration)

	events = f.Apply(events)
	if len(events) != 0 {
		t.Error(events)
	}
}
