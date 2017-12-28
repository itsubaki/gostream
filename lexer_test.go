package gocep

import (
	"strings"
	"testing"
)

func TestLexerTokenize(t *testing.T) {
	q := "select Value, count(*), avg(Value), sum(Value) from MyEvent.time(10 sec) where Value > 97"
	lexer := NewLexer(strings.NewReader(q))

	var test = []struct {
		token   Token
		literal string
	}{
		{SELECT, "select"},
		{WHITESPACE, " "},
		{IDENTIFIER, "Value"},
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
		{IDENTIFIER, "Value"},
		{ILLEGAL, ")"},
		{COMMA, ","},
		{WHITESPACE, " "},
		{SUM, "sum"},
		{ILLEGAL, "("},
		{IDENTIFIER, "Value"},
		{ILLEGAL, ")"},
		{WHITESPACE, " "},
		{FROM, "from"},
		{WHITESPACE, " "},
		{IDENTIFIER, "MyEvent"},
		{DOT, "."},
		{TIME, "time"},
		{ILLEGAL, "("},
		{IDENTIFIER, "10"},
		{WHITESPACE, " "},
		{SEC, "sec"},
		{ILLEGAL, ")"},
		{WHITESPACE, " "},
		{WHERE, "where"},
		{WHITESPACE, " "},
		{IDENTIFIER, "Value"},
		{WHITESPACE, " "},
		{LARGER, ">"},
		{WHITESPACE, " "},
		{IDENTIFIER, "97"},
	}

	for _, tt := range test {
		token, literal := lexer.Tokenize()
		if token != tt.token || literal != tt.literal {
			t.Error(token, literal)
		}
	}
}

func TestLexerTokenizeIgnoreSpaceTimeWindow(t *testing.T) {
	q := "select count(*) from LogEvent.time(10 sec) where Level > 2"
	lexer := NewLexer(strings.NewReader(q))

	var test = []struct {
		token   Token
		literal string
	}{
		{SELECT, "select"},
		{COUNT, "count"},
		{ILLEGAL, "("},
		{ASTERISK, "*"},
		{ILLEGAL, ")"},
		{FROM, "from"},
		{IDENTIFIER, "LogEvent"},
		{DOT, "."},
		{TIME, "time"},
		{ILLEGAL, "("},
		{IDENTIFIER, "10"},
		{SEC, "sec"},
		{ILLEGAL, ")"},
		{WHERE, "where"},
		{IDENTIFIER, "Level"},
		{LARGER, ">"},
		{IDENTIFIER, "2"},
	}

	for _, tt := range test {
		token, literal := lexer.TokenizeIgnoreWhiteSpace()
		if token != tt.token || literal != tt.literal {
			t.Error(token, literal)
		}
	}
}

func TestLexerTokenizeIgnoreSpace(t *testing.T) {
	q := "select Value, count(*), avg(Value), sum(Value) from MyEvent.time(10 sec) where Value > 97"
	lexer := NewLexer(strings.NewReader(q))

	var test = []struct {
		token   Token
		literal string
	}{
		{SELECT, "select"},
		{IDENTIFIER, "Value"},
		{COMMA, ","},
		{COUNT, "count"},
		{ILLEGAL, "("},
		{ASTERISK, "*"},
		{ILLEGAL, ")"},
		{COMMA, ","},
		{AVG, "avg"},
		{ILLEGAL, "("},
		{IDENTIFIER, "Value"},
		{ILLEGAL, ")"},
		{COMMA, ","},
		{SUM, "sum"},
		{ILLEGAL, "("},
		{IDENTIFIER, "Value"},
		{ILLEGAL, ")"},
		{FROM, "from"},
		{IDENTIFIER, "MyEvent"},
		{DOT, "."},
		{TIME, "time"},
		{ILLEGAL, "("},
		{IDENTIFIER, "10"},
		{SEC, "sec"},
		{ILLEGAL, ")"},
		{WHERE, "where"},
		{IDENTIFIER, "Value"},
		{LARGER, ">"},
		{IDENTIFIER, "97"},
	}

	for _, tt := range test {
		token, literal := lexer.TokenizeIgnoreWhiteSpace()
		if token != tt.token || literal != tt.literal {
			t.Error(token, literal)
		}
	}
}
