package stream_test

import (
	"fmt"

	"github.com/itsubaki/gostream/stream"
)

func ExampleOrderBy() {
	type LogEvent struct {
		Level int
	}

	e := make([]stream.Event, 0)
	for i := 0; i < 10; i++ {
		e = append(e, stream.NewEvent(LogEvent{
			Level: i,
		}))
	}

	o := &stream.OrderBy{
		Name:  "Level",
		Index: 0,
		Desc:  false,
	}

	out := o.Apply(e)
	for _, ev := range out {
		fmt.Print(ev.Underlying)
	}

	// Output:
	// {0}{1}{2}{3}{4}{5}{6}{7}{8}{9}
}

func ExampleOrderBy_desc() {
	type LogEvent struct {
		Level int
	}

	o := &stream.OrderBy{
		Name:  "Level",
		Index: 0,
		Desc:  true,
	}

	e := make([]stream.Event, 0)
	for i := 0; i < 10; i++ {
		e = append(e, stream.NewEvent(LogEvent{
			Level: i,
		}))
	}

	out := o.Apply(e)
	for _, ev := range out {
		fmt.Print(ev.Underlying)
	}

	// Output:
	// {9}{8}{7}{6}{5}{4}{3}{2}{1}{0}
}
