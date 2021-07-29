package lexer_test

import (
	"strings"
	"testing"

	"github.com/itsubaki/gostream/pkg/lexer"
)

func TestLexerIgnore(t *testing.T) {
	type token struct {
		token lexer.Token
		str   string
	}

	var cases = []struct {
		in   string
		want []token
	}{
		{
			in: "select min(Level) from LogEvent.time(10 sec) where Level > 2",
			want: []token{
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
			},
		},
		{
			in: "select count(*) from LogEvent.time(10 sec) where Level > 2.5",
			want: []token{
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
			},
		},
		{
			in: "select count(*) from LogEvent.time(10 sec) where Level > 2",
			want: []token{
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
			},
		},
		{
			in: "select Value, count(*), avg(Value), sum(Value) from MyEvent.time(10 sec) where Value > 97",
			want: []token{
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
			},
		},
	}

	for _, c := range cases {
		lex := lexer.New(strings.NewReader(c.in))
		for _, w := range c.want {
			token, str := lex.TokenizeIgnore(lexer.WHITESPACE)
			if token != w.token || str != w.str {
				t.Fail()
			}
		}
	}
}

func TestLexerTokenize(t *testing.T) {
	type token struct {
		token lexer.Token
		str   string
	}

	var cases = []struct {
		in   string
		want []token
	}{
		{
			in: "select Value, count(*), avg(Value), sum(Value) from MyEvent.time(10 sec) where Value > 97",
			want: []token{
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
			},
		},
	}

	for _, c := range cases {
		lex := lexer.New(strings.NewReader(c.in))
		for _, w := range c.want {
			token, str := lex.Tokenize()
			if token != w.token || str != w.str {
				t.Fail()
			}
		}
	}
}
