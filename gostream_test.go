package gostream_test

import (
	"fmt"
	"testing"
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

	fmt.Println(s)

	// Output:
	// SELECT * FROM IDENT(LogEvent) . LENGTH ( INT(10) )
	// SELECT * FROM LogEvent.LENGTH(10)
}

func TestGoStreamLength(t *testing.T) {
	type LogEvent struct {
		Time    time.Time
		Level   int
		Message string
	}

	s, err := gostream.New().
		Add(LogEvent{}).
		Query("select * from LogEvent.length(3)")
	if err != nil {
		fmt.Printf("query: %v", err)
		return
	}
	defer s.Close()

	for i := 0; i < 10; i++ {
		s.Input() <- LogEvent{}
		out := <-s.Output()
		if len(out) > 3 {
			t.Errorf("len(out)=%v", len(out))
		}
	}
}
