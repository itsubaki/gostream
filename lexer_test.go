package gocep

import (
	"strings"
	"testing"
)

func TestLexer(t *testing.T) {
	q := "select Value, count(*), avg(Value), sum(Value) from MyEvent.time(10 sec) where Value > 97"
	lexer := NewLexer(strings.NewReader(q))

	var test = []struct {
		token   Token
		literal string
	}{
		{SELECT, "select"},
		{WHITESPACE, " "},
		{LITERAL, "Value"},
		{COMMA, ","},
		{WHITESPACE, " "},
		{COUNT, "count"},
		{ILLEGAL, "("},
		{ASTERISK, "*"},
		{ILLEGAL, ")"},
		{COMMA, ","},
		{WHITESPACE, " "},
		{AVG, "avg"},
		{ILLEGAL, "("},
		{LITERAL, "Value"},
		{ILLEGAL, ")"},
		{COMMA, ","},
		{WHITESPACE, " "},
		{SUM, "sum"},
		{ILLEGAL, "("},
		{LITERAL, "Value"},
		{ILLEGAL, ")"},
		{WHITESPACE, " "},
		{FROM, "from"},
		{WHITESPACE, " "},
		{LITERAL, "MyEvent"},
		{DOT, "."},
		{TIME, "time"},
		{ILLEGAL, "("},
		{LITERAL, "10"},
		{WHITESPACE, " "},
		{SEC, "sec"},
		{ILLEGAL, ")"},
		{WHITESPACE, " "},
		{WHERE, "where"},
		{WHITESPACE, " "},
		{LITERAL, "Value"},
		{WHITESPACE, " "},
		{LARGER, ">"},
		{WHITESPACE, " "},
		{LITERAL, "97"},
	}

	for _, tt := range test {
		token, literal := lexer.Tokenize()
		if token != tt.token || literal != tt.literal {
			t.Error(token, literal)
		}
	}
}
