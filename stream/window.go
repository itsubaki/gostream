package stream

import (
	"fmt"
	"time"

	"github.com/itsubaki/gostream/lexer"
)

type Window interface {
	Apply(e []Event) []Event
	String() string
}

type Length struct {
	Length int
}

func (w *Length) Apply(e []Event) []Event {
	if len(e) > w.Length {
		e = e[1:]
	}

	return e
}

func (w *Length) String() string {
	return fmt.Sprintf("LENGTH(%v)", w.Length)
}

type LengthBatch struct {
	Length int
	Batch  []Event
}

func (w *LengthBatch) Apply(e []Event) []Event {
	w.Batch = append(w.Batch, e[len(e)-1])

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

func (w *Time) Apply(e []Event) []Event {
	out := make([]Event, 0)
	for _, ev := range e {
		if time.Since(ev.Time) < w.Expire {
			out = append(out, ev)
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

func (w *TimeBatch) Apply(e []Event) []Event {
	for {
		if time.Since(w.Start) < w.Expire {
			break
		}

		w.Start = w.Start.Add(w.Expire)
		w.End = w.Start.Add(w.Expire)
	}

	out := make([]Event, 0)
	for _, ev := range e {
		if !ev.Time.Before(w.Start) && !ev.Time.After(w.End) {
			out = append(out, ev)
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
