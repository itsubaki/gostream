package gostream

import "time"

type Function interface {
	Apply(events []Event) []Event
}

type Length struct {
	Length int
}

func (f *Length) Apply(events []Event) []Event {
	if len(events) > f.Length {
		events = events[1:]
	}

	return events
}

type LengthBatch struct {
	Length int
	Batch  []Event
}

func (f *LengthBatch) Apply(events []Event) []Event {
	f.Batch = append(f.Batch, events[len(events)-1])

	out := make([]Event, 0)
	if len(f.Batch) == f.Length {
		out, f.Batch = f.Batch, out
		return out
	}

	return out
}

type Time struct {
	Expire time.Duration
}

func (f *Time) Apply(events []Event) []Event {
	out := make([]Event, 0)
	for _, e := range events {
		if time.Since(e.Time) < f.Expire {
			out = append(out, e)
		}
	}

	return out
}

type TimeBatch struct {
	Start  time.Time
	End    time.Time
	Expire time.Duration
}

func (f *TimeBatch) Apply(events []Event) []Event {
	for {
		if time.Since(f.Start) < f.Expire {
			break
		}
		f.Start = f.Start.Add(f.Expire)
		f.End = f.Start.Add(f.Expire)
	}

	out := make([]Event, 0)
	for _, e := range events {
		if !e.Time.Before(f.Start) && !e.Time.After(f.End) {
			out = append(out, e)
		}
	}

	return out
}
