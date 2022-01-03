package stream_test

import (
	"testing"
	"time"

	"github.com/itsubaki/gostream/stream"
)

func TestWhereString(t *testing.T) {
	type LogEvent struct {
		Time    time.Time
		Level   int
		Message string
	}

	cases := []struct {
		in   stream.Where
		want string
	}{
		{stream.From{Type: LogEvent{}}, "LogEvent"},
		{stream.LargerThan{Name: "Level", Value: 2}, "Level > 2"},
	}

	for _, c := range cases {
		if c.in.String() != c.want {
			t.Errorf("got=%v, want=%v", c.in.String(), c.want)
		}
	}
}

func TestWhere(t *testing.T) {
	type LogEvent struct {
		Time    time.Time
		Level   interface{}
		Message string
	}

	cases := []struct {
		w    stream.Where
		in   LogEvent
		want bool
	}{
		{stream.LargerThan{Name: "Level", Value: 2}, LogEvent{Level: 1}, false},
		{stream.LargerThan{Name: "Level", Value: 2}, LogEvent{Level: 2}, false},
		{stream.LargerThan{Name: "Level", Value: 2}, LogEvent{Level: 3}, true},
		{stream.LessThan{Name: "Level", Value: 2}, LogEvent{Level: 1}, true},
		{stream.LessThan{Name: "Level", Value: 2}, LogEvent{Level: 2}, false},
		{stream.LessThan{Name: "Level", Value: 2}, LogEvent{Level: 3}, false},
		{stream.Equals{Name: "Level", Value: 2}, LogEvent{Level: 1}, false},
		{stream.Equals{Name: "Level", Value: 2}, LogEvent{Level: 2}, true},
		{stream.Equals{Name: "Level", Value: 2}, LogEvent{Level: 3}, false},
		{stream.NotEquals{Name: "Level", Value: 2}, LogEvent{Level: 1}, true},
		{stream.NotEquals{Name: "Level", Value: 2}, LogEvent{Level: 2}, false},
		{stream.NotEquals{Name: "Level", Value: 2}, LogEvent{Level: 3}, true},
		{stream.LargerThan{Name: "Level", Value: 2.0}, LogEvent{Level: 1.0}, false},
		{stream.LargerThan{Name: "Level", Value: 2.0}, LogEvent{Level: 2.0}, false},
		{stream.LargerThan{Name: "Level", Value: 2.0}, LogEvent{Level: 3.0}, true},
	}

	for _, c := range cases {
		got := c.w.Apply(c.in)
		if got != c.want {
			t.Errorf("want=%v, got=%v", c.want, got)
		}
	}
}
