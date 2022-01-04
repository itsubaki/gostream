package stream_test

import (
	"fmt"

	"github.com/itsubaki/gostream/stream"
)

func ExampleAverage() {
	type LogEvent struct {
		Level int
	}

	e := make([]stream.Event, 0)
	for i := 0; i < 10; i++ {
		e = append(e, stream.NewEvent(LogEvent{
			Level: i,
		}))
	}

	avg := &stream.Average{Name: "Level"}
	out := avg.Apply(e)

	fmt.Println(out[len(out)-1].ResultSet)

	// Output:
	// [4.5]
}

func ExampleSum() {
	type LogEvent struct {
		Level int
	}

	e := make([]stream.Event, 0)
	for i := 0; i < 10; i++ {
		e = append(e, stream.NewEvent(LogEvent{
			Level: i,
		}))
	}

	s := &stream.Sum{Name: "Level"}
	out := s.Apply(e)

	fmt.Println(out[len(out)-1].ResultSet)

	// Output:
	// [45]
}

func ExampleCount() {
	type LogEvent struct {
		Level int
	}

	e := make([]stream.Event, 0)
	for i := 0; i < 10; i++ {
		e = append(e, stream.NewEvent(LogEvent{
			Level: i,
		}))
	}

	c := &stream.Count{Name: "Level"}
	out := c.Apply(e)

	fmt.Println(out[len(out)-1].ResultSet)

	// Output:
	// [10]
}

func ExampleMax() {
	type LogEvent struct {
		Level int
	}

	e := make([]stream.Event, 0)
	for i := 0; i < 10; i++ {
		e = append(e, stream.NewEvent(LogEvent{
			Level: i,
		}))
	}

	m := &stream.Max{Name: "Level"}
	out := m.Apply(e)

	fmt.Println(out[len(out)-1].ResultSet)

	// Output:
	// [9]
}

func ExampleMin() {
	type LogEvent struct {
		Level int
	}

	e := make([]stream.Event, 0)
	for i := 0; i < 10; i++ {
		e = append(e, stream.NewEvent(LogEvent{
			Level: i,
		}))
	}

	m := &stream.Min{Name: "Level"}
	out := m.Apply(e)

	fmt.Println(out[len(out)-1].ResultSet)

	// Output:
	// [0]
}
