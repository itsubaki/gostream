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
