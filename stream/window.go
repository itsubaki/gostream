package stream

import (
	"fmt"
	"time"

	"github.com/itsubaki/gostream/lexer"
)

type Window interface {
	Apply(events []Event) []Event
	String() string
}

type Length struct {
	Length int
}

func (w *Length) Apply(events []Event) []Event {
	if len(events) > w.Length {
		events = events[1:]
	}

	return events
}

func (w *Length) String() string {
	return fmt.Sprintf("LENGTH(%v)", w.Length)
}

type LengthBatch struct {
	Length int
	Batch  []Event
}

func (w *LengthBatch) Apply(events []Event) []Event {
	w.Batch = append(w.Batch, events[len(events)-1])

	out := make([]Event, 0)
	if len(w.Batch) == w.Length {
		out, w.Batch = w.Batch, out
		return out
	}

	return out
}

func (w *LengthBatch) String() string {
	return fmt.Sprintf("LENGTH_BATCH(%v)", w.Length)
}

type Time struct {
	Expire time.Duration
	Unit   lexer.Token
}

func (w *Time) Apply(events []Event) []Event {
	out := make([]Event, 0)
	for _, e := range events {
		if time.Since(e.Time) < w.Expire {
			out = append(out, e)
		}
	}

	return out
}

func (w *Time) String() string {
	v := w.Expire.Seconds()
	if w.Unit == lexer.MIN {
		v = w.Expire.Minutes()
	}
	if w.Unit == lexer.HOUR {
		v = w.Expire.Hours()
	}

	return fmt.Sprintf("TIME(%v %v)", v, lexer.Tokens[w.Unit])
}

type TimeBatch struct {
	Start  time.Time
	End    time.Time
	Expire time.Duration
	Unit   lexer.Token
}

func (w *TimeBatch) Apply(events []Event) []Event {
	for {
		if time.Since(w.Start) < w.Expire {
			break
		}

		w.Start = w.Start.Add(w.Expire)
		w.End = w.Start.Add(w.Expire)
	}

	out := make([]Event, 0)
	for _, e := range events {
		if !e.Time.Before(w.Start) && !e.Time.After(w.End) {
			out = append(out, e)
		}
	}

	return out
}

func (w *TimeBatch) String() string {
	v := w.Expire.Seconds()
	if w.Unit == lexer.MIN {
		v = w.Expire.Minutes()
	}
	if w.Unit == lexer.HOUR {
		v = w.Expire.Hours()
	}

	return fmt.Sprintf("TIME_BATCH(%v %v)", v, lexer.Tokens[w.Unit])
}
