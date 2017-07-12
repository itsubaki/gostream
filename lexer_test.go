package gocep

import (
	"fmt"
	"strings"
	"testing"
)

func TestLexer(t *testing.T) {
	q := "select Value, count(*), avg(Value) from MyEvent.time(10 sec) where Value > 97"
	s := NewLexer(strings.NewReader(q))
	for {
		tok, lit := s.Tokenize()
		if tok == EOF {
			break
		}
		if tok == ILLEGAL {
			continue
		}
		if tok == WHITESPACE {
			continue
		}
		if tok == COMMA {
			continue
		}
		fmt.Println(tok, lit)
	}
}
