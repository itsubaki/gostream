package _example

import (
	"fmt"
	"time"

	"github.com/itsubaki/gostream/pkg/clause"
	"github.com/itsubaki/gostream/pkg/event"
	"github.com/itsubaki/gostream/pkg/stream"
)

func TimeWindow() {
	type LogEvent struct {
		Time    time.Time
		Level   int
		Message string
	}

	w := stream.NewTime(LogEvent{}, 10*time.Second)
	defer w.Close()

	w.SetWhere(
		clause.LargerThanInt{
			Name:  "Level",
			Value: 2,
		},
	)

	w.SetFunction(
		clause.Count{
			As: "count",
		},
	)

	go func() {
		for {
			newest := event.Newest(<-w.Output())
			if newest.Int("count") > 10 {
				fmt.Println("Notify!")
			}
		}
	}()

	w.Input() <- LogEvent{
		Time:    time.Now(),
		Level:   1,
		Message: "this is text log.",
	}
}

func LengthWindow() {
	type MyEvent struct {
		Name  string
		Value int
	}

	w := stream.NewLength(MyEvent{}, 10)
	defer w.Close()

	w.SetFunction(
		clause.AverageInt{
			Name: "Value",
			As:   "avg(Value)",
		},
		clause.SumInt{
			Name: "Value",
			As:   "sum(Value)",
		},
	)
}

func View() {
	type MyEvent struct {
		Name  string
		Value int
	}

	w := stream.NewTime(MyEvent{}, 10*time.Millisecond)
	defer w.Close()

	w.SetWhere(
		clause.LargerThanInt{
			Name:  "Value",
			Value: 97,
		},
	)
	w.SetFunction(
		clause.SelectString{
			Name: "Name",
			As:   "n",
		},
		clause.SelectInt{
			Name: "Value",
			As:   "v",
		},
	)
	w.SetOrderBy(
		clause.OrderByInt{
			Name:    "Value",
			Reverse: true,
		},
	)
	w.SetLimit(
		clause.Limit{
			Limit:  10,
			Offset: 5,
		})

	go func() {
		for {
			fmt.Println(<-w.Output())
		}
	}()

	for i := 0; i < 100; i++ {
		w.Input() <- MyEvent{
			Name:  "name",
			Value: i,
		}
	}
}
