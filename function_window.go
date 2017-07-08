package gocep

import "time"

type Length struct {
	length int
}

func (f Length) Apply(event []Event) []Event {
	if len(event) > f.length {
		event = event[1:]
	}
	return event
}

type TimeDuration struct {
	expire time.Duration
}

func (f TimeDuration) Apply(event []Event) (stream []Event) {
	for _, e := range event {
		if time.Since(e.Time) < f.expire {
			stream = append(stream, e)
		}
	}
	return stream
}

type LengthBatch struct {
	length int
	batch  []Event
}

func (f *LengthBatch) Apply(event []Event) (stream []Event) {
	f.batch = append(f.batch, event[len(event)-1])
	if len(f.batch) == f.length {
		stream, f.batch = f.batch, stream
		return stream
	}
	return stream
}

type TimeDurationBatch struct {
	start  time.Time
	end    time.Time
	expire time.Duration
}

func (f *TimeDurationBatch) Apply(event []Event) (stream []Event) {
	for {
		if time.Since(f.start) < f.expire {
			break
		}
		f.start = f.start.Add(f.expire)
		f.end = f.start.Add(f.expire)
	}

	for _, e := range event {
		if !e.Time.Before(f.start) && !e.Time.After(f.end) {
			stream = append(stream, e)
		}
	}
	return stream
}
