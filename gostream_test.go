package gostream_test

import (
	"time"

	"github.com/itsubaki/gostream/pkg/parser"
	"github.com/itsubaki/gostream/pkg/window"
)

func Example_parser() {
	type MyEvent struct {
		Name  string
		Value int
	}

	p := parser.New()
	p.Register("MyEvent", MyEvent{})

	query := "select * from MyEvent.length(10)"
	statement, err := p.Parse(query)
	if err != nil {
		return
	}

	window := statement.New()
	defer window.Close()

	// Output:
}

func Example_timeWindow() {
	type LogEvent struct {
		Time    time.Time
		Level   int
		Message string
	}

	w := window.NewTime(LogEvent{}, 10*time.Second)
	defer w.Close()

	w.Where().LargerThan().Int("Level", 2)
	w.Function().Count()

	// Output:
}

func Example_lengthWindow() {
	type MyEvent struct {
		Name  string
		Value int
	}

	w := window.NewLength(MyEvent{}, 10)
	defer w.Close()

	w.Function().Average().Int("Value")
	w.Function().Sum().Int("Value")

	// Output:
}

func Example_view() {
	type MyEvent struct {
		Name  string
		Value int
	}

	w := window.NewTime(MyEvent{}, 10*time.Millisecond)
	defer w.Close()

	w.Where().LargerThan().Int("Value", 97)
	w.Function().Select().String("Name")
	w.Function().Select().Int("Value")
	w.OrderBy().Desc().Int("Value")
	w.Limit(10).Offset(5)

	// Output:
}
