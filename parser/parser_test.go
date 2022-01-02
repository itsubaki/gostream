package parser_test

import (
	"fmt"
	"time"

	"github.com/itsubaki/gostream/parser"
)

func ExampleParse() {
	type LogEvent struct {
		Time    time.Time
		Level   int
		Message string
	}

	q := "select * from LogEvent.length(10)"
	p := parser.New().Query(q).Add(LogEvent{})

	s := p.Parse()
	if len(p.Errors()) > 0 {
		fmt.Printf("%v", p.Errors())
		return
	}

	fmt.Println(p)
	fmt.Println(s)

	// Output:
	// LogEvent
	// *stream.Length
}
