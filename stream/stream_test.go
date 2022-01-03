package stream_test

import (
	"fmt"
	"time"

	"github.com/itsubaki/gostream/stream"
)

func ExampleStream() {
	type LogEvent struct {
		Time    time.Time
		Level   int
		Message string
	}

	s := stream.New().
		SelectAll().
		From(LogEvent{}).
		Length(10).
		OrderBy("Level", true).
		Limit(10, 5)

	fmt.Println(s)

	// Output:
	// SELECT * FROM LogEvent.LENGTH(10) ORDER BY Level DESC LIMIT 10 OFFSET 5
}
