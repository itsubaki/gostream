package gostream_test

import (
	"fmt"
	"strings"
	"time"

	"github.com/itsubaki/gostream"
	"github.com/itsubaki/gostream/lexer"
)

func ExampleParse() {
	type LogEvent struct {
		Time    time.Time
		Level   int
		Message string
	}

	r := make(gostream.Registry)
	r.Add(LogEvent{})

	q := "select * from LogEvent.length(10)"
	p := gostream.NewParser(lexer.New(strings.NewReader(q)), r)
	s := p.Parse()

	if len(p.Errors()) > 0 {
		fmt.Printf("%v", p.Errors())
		return
	}

	fmt.Println(s)

	// Output:
	// *gostream.Length
}
