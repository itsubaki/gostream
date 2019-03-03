package function

import (
	"time"

	"github.com/itsubaki/gostream-core/pkg/event"
)

type Length struct {
	Length int
}

func (f *Length) Apply(events []event.Event) []event.Event {
	if len(events) > f.Length {
		events = events[1:]
	}
	return events
}

type TimeDuration struct {
	Expire time.Duration
}

func (f *TimeDuration) Apply(events []event.Event) []event.Event {
	out := event.List()
	for _, e := range events {
		if time.Since(e.Time) < f.Expire {
			out = append(out, e)
		}
	}
	return out
}

type LengthBatch struct {
	Length int
	Batch  []event.Event
}

func (f *LengthBatch) Apply(events []event.Event) []event.Event {
	f.Batch = append(f.Batch, events[len(events)-1])

	out := event.List()
	if len(f.Batch) == f.Length {
		out, f.Batch = f.Batch, out
		return out
	}

	return out
}

type TimeDurationBatch struct {
	Start  time.Time
	End    time.Time
	Expire time.Duration
}

func (f *TimeDurationBatch) Apply(events []event.Event) []event.Event {
	for {
		if time.Since(f.Start) < f.Expire {
			break
		}
		f.Start = f.Start.Add(f.Expire)
		f.End = f.Start.Add(f.Expire)
	}

	out := event.List()
	for _, e := range events {
		if !e.Time.Before(f.Start) && !e.Time.After(f.End) {
			out = append(out, e)
		}
	}

	return out
}
