package lexer_test

import (
	"strings"
	"testing"

	"github.com/itsubaki/gostream/pkg/lexer"
)

func TestLexerTokenizeMin(t *testing.T) {
	q := "select min(Level) from LogEvent.time(10 sec) where Level > 2"
	lex := lexer.New(strings.NewReader(q))

	var test = []struct {
		token   lexer.Token
		literal string
	}{
		{lexer.SELECT, "select"},
		{lexer.MIN, "min"},
		{lexer.LPAREN, "("},
		{lexer.IDENTIFIER, "Level"},
		{lexer.RPAREN, ")"},
		{lexer.FROM, "from"},
		{lexer.IDENTIFIER, "LogEvent"},
		{lexer.DOT, "."},
		{lexer.TIME, "time"},
		{lexer.LPAREN, "("},
		{lexer.IDENTIFIER, "10"},
		{lexer.SEC, "sec"},
		{lexer.RPAREN, ")"},
		{lexer.WHERE, "where"},
		{lexer.IDENTIFIER, "Level"},
		{lexer.LARGER, ">"},
		{lexer.IDENTIFIER, "2"},
	}

	for _, tt := range test {
		token, literal := lex.TokenizeIgnoreWhiteSpace()
		if token != tt.token || literal != tt.literal {
			t.Error(token, literal)
		}
	}
}

func TestLexerTokenizeFloat(t *testing.T) {
	q := "select count(*) from LogEvent.time(10 sec) where Level > 2.5"
	lex := lexer.New(strings.NewReader(q))

	var test = []struct {
		token   lexer.Token
		literal string
	}{
		{lexer.SELECT, "select"},
		{lexer.COUNT, "count"},
		{lexer.LPAREN, "("},
		{lexer.ASTERISK, "*"},
		{lexer.RPAREN, ")"},
		{lexer.FROM, "from"},
		{lexer.IDENTIFIER, "LogEvent"},
		{lexer.DOT, "."},
		{lexer.TIME, "time"},
		{lexer.LPAREN, "("},
		{lexer.IDENTIFIER, "10"},
		{lexer.SEC, "sec"},
		{lexer.RPAREN, ")"},
		{lexer.WHERE, "where"},
		{lexer.IDENTIFIER, "Level"},
		{lexer.LARGER, ">"},
		{lexer.IDENTIFIER, "2"},
		{lexer.DOT, "."},
		{lexer.IDENTIFIER, "5"},
	}

	for _, tt := range test {
		token, literal := lex.TokenizeIgnoreWhiteSpace()
		if token != tt.token || literal != tt.literal {
			t.Error(token, literal)
		}
	}
}

func TestLexerTokenize(t *testing.T) {
	q := "select Value, count(*), avg(Value), sum(Value) from MyEvent.time(10 sec) where Value > 97"
	lex := lexer.New(strings.NewReader(q))

	var test = []struct {
		token   lexer.Token
		literal string
	}{
		{lexer.SELECT, "select"},
		{lexer.WHITESPACE, " "},
		{lexer.IDENTIFIER, "Value"},
		{lexer.COMMA, ","},
		{lexer.WHITESPACE, " "},
		{lexer.COUNT, "count"},
		{lexer.LPAREN, "("},
		{lexer.ASTERISK, "*"},
		{lexer.RPAREN, ")"},
		{lexer.COMMA, ","},
		{lexer.WHITESPACE, " "},
		{lexer.AVG, "avg"},
		{lexer.LPAREN, "("},
		{lexer.IDENTIFIER, "Value"},
		{lexer.RPAREN, ")"},
		{lexer.COMMA, ","},
		{lexer.WHITESPACE, " "},
		{lexer.SUM, "sum"},
		{lexer.LPAREN, "("},
		{lexer.IDENTIFIER, "Value"},
		{lexer.RPAREN, ")"},
		{lexer.WHITESPACE, " "},
		{lexer.FROM, "from"},
		{lexer.WHITESPACE, " "},
		{lexer.IDENTIFIER, "MyEvent"},
		{lexer.DOT, "."},
		{lexer.TIME, "time"},
		{lexer.LPAREN, "("},
		{lexer.IDENTIFIER, "10"},
		{lexer.WHITESPACE, " "},
		{lexer.SEC, "sec"},
		{lexer.RPAREN, ")"},
		{lexer.WHITESPACE, " "},
		{lexer.WHERE, "where"},
		{lexer.WHITESPACE, " "},
		{lexer.IDENTIFIER, "Value"},
		{lexer.WHITESPACE, " "},
		{lexer.LARGER, ">"},
		{lexer.WHITESPACE, " "},
		{lexer.IDENTIFIER, "97"},
	}

	for _, tt := range test {
		token, literal := lex.Tokenize()
		if token != tt.token || literal != tt.literal {
			t.Error(token, literal)
		}
	}
}

func TestLexerTokenizeIgnoreSpaceTimeWindow(t *testing.T) {
	q := "select count(*) from LogEvent.time(10 sec) where Level > 2"
	lex := lexer.New(strings.NewReader(q))

	var test = []struct {
		token   lexer.Token
		literal string
	}{
		{lexer.SELECT, "select"},
		{lexer.COUNT, "count"},
		{lexer.LPAREN, "("},
		{lexer.ASTERISK, "*"},
		{lexer.RPAREN, ")"},
		{lexer.FROM, "from"},
		{lexer.IDENTIFIER, "LogEvent"},
		{lexer.DOT, "."},
		{lexer.TIME, "time"},
		{lexer.LPAREN, "("},
		{lexer.IDENTIFIER, "10"},
		{lexer.SEC, "sec"},
		{lexer.RPAREN, ")"},
		{lexer.WHERE, "where"},
		{lexer.IDENTIFIER, "Level"},
		{lexer.LARGER, ">"},
		{lexer.IDENTIFIER, "2"},
	}

	for _, tt := range test {
		token, literal := lex.TokenizeIgnoreWhiteSpace()
		if token != tt.token || literal != tt.literal {
			t.Error(token, literal)
		}
	}
}

func TestLexerTokenizeIgnoreSpace(t *testing.T) {
	q := "select Value, count(*), avg(Value), sum(Value) from MyEvent.time(10 sec) where Value > 97"
	lex := lexer.New(strings.NewReader(q))

	var test = []struct {
		token   lexer.Token
		literal string
	}{
		{lexer.SELECT, "select"},
		{lexer.IDENTIFIER, "Value"},
		{lexer.COMMA, ","},
		{lexer.COUNT, "count"},
		{lexer.LPAREN, "("},
		{lexer.ASTERISK, "*"},
		{lexer.RPAREN, ")"},
		{lexer.COMMA, ","},
		{lexer.AVG, "avg"},
		{lexer.LPAREN, "("},
		{lexer.IDENTIFIER, "Value"},
		{lexer.RPAREN, ")"},
		{lexer.COMMA, ","},
		{lexer.SUM, "sum"},
		{lexer.LPAREN, "("},
		{lexer.IDENTIFIER, "Value"},
		{lexer.RPAREN, ")"},
		{lexer.FROM, "from"},
		{lexer.IDENTIFIER, "MyEvent"},
		{lexer.DOT, "."},
		{lexer.TIME, "time"},
		{lexer.LPAREN, "("},
		{lexer.IDENTIFIER, "10"},
		{lexer.SEC, "sec"},
		{lexer.RPAREN, ")"},
		{lexer.WHERE, "where"},
		{lexer.IDENTIFIER, "Value"},
		{lexer.LARGER, ">"},
		{lexer.IDENTIFIER, "97"},
	}

	for _, tt := range test {
		token, literal := lex.TokenizeIgnoreWhiteSpace()
		if token != tt.token || literal != tt.literal {
			t.Error(token, literal)
		}
	}
}
