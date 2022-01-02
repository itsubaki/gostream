package gostream_test

import (
	"fmt"
	"time"

	"github.com/itsubaki/gostream"
)

func ExampleGoStream_Query() {
	type LogEvent struct {
		Time    time.Time
		Level   int
		Message string
	}

	s, err := gostream.
		New(&gostream.Option{
			Verbose: true,
		}).
		Add(LogEvent{}).
		Query("select * from LogEvent.length(10)")
	if err != nil {
		fmt.Printf("query: %v", err)
		return
	}
	defer s.Close()

	// Output:
	// SELECT * FROM IDENT(LogEvent) . LENGTH ( INT(10) )
}
