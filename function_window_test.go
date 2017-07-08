package gocep

import (
	"testing"
	"time"
)

func TestTimeDurationBatch(t *testing.T) {
	start := time.Now()
	duration := 10 * time.Millisecond
	f := &TimeDurationBatch{start, start.Add(duration), duration}

	event := []Event{}
	event = append(event, NewEvent(IntEvent{"foo", 1}))
	event = f.Apply(event)
	if len(event) != 1 {
		t.Error(event)
	}
	time.Sleep(duration)

	event = f.Apply(event)
	if len(event) != 0 {
		t.Error(event)
	}
}
