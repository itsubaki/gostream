package parser_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/itsubaki/gostream/parser"
)

func ExampleParse_length() {
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

	fmt.Println(s)

	// Output:
	// SELECT * FROM LogEvent.LENGTH(10)
}

func ExampleParse_time() {
	type LogEvent struct {
		Time    time.Time
		Level   int
		Message string
	}

	q := "SELECT * FROM LogEvent.TIME(10 MIN) ORDER BY Level DESC LIMIT 1 OFFSET 1"
	p := parser.New().Query(q).Add(LogEvent{})

	s := p.Parse()
	if len(p.Errors()) > 0 {
		fmt.Printf("%v", p.Errors())
		return
	}

	fmt.Println(s)

	// Output:
	// SELECT * FROM LogEvent.TIME(10 MIN) ORDER BY Level DESC LIMIT 1 OFFSET 1
}

func TestParse(t *testing.T) {
	type LogEvent struct {
		Time    time.Time
		Level   int
		Message string
	}

	var cases = []struct {
		in string
	}{
		{"SELECT * FROM LogEvent.LENGTH(10)"},
		{"SELECT `Time`, Level, Message FROM LogEvent.LENGTH(10)"},
	}

	p := parser.New().Add(LogEvent{})
	for _, c := range cases {
		s := p.Query(c.in).Parse()
		if len(p.Errors()) > 0 {
			t.Errorf("%v", p.Errors())
		}

		if s.String() != c.in {
			t.Errorf("want=%v, got=%v", c.in, s.String())
		}
	}
}
