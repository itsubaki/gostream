package lexer_test

import (
	"strings"
	"testing"

	"github.com/itsubaki/gostream/lexer"
)

func TestLexer(t *testing.T) {
	type Token struct {
		token   lexer.Token
		literal string
	}

	var cases = []struct {
		in   string
		want []Token
	}{
		{
			in: "select * from LogEvent.length(10) where Level > 2",
			want: []Token{
				{lexer.SELECT, "select"},
				{lexer.ASTERISK, "*"},
				{lexer.FROM, "from"},
				{lexer.IDENT, "LogEvent"},
				{lexer.DOT, "."},
				{lexer.LENGTH, "length"},
				{lexer.LPAREN, "("},
				{lexer.INT, "10"},
				{lexer.RPAREN, ")"},
				{lexer.WHERE, "where"},
				{lexer.IDENT, "Level"},
				{lexer.LARGER, ">"},
				{lexer.INT, "2"},
			},
		},
		{
			in: "select * from LogEvent.time(10 sec)",
			want: []Token{
				{lexer.SELECT, "select"},
				{lexer.ASTERISK, "*"},
				{lexer.FROM, "from"},
				{lexer.IDENT, "LogEvent"},
				{lexer.DOT, "."},
				{lexer.TIME, "time"},
				{lexer.LPAREN, "("},
				{lexer.INT, "10"},
				{lexer.SEC, "sec"},
				{lexer.RPAREN, ")"},
			},
		},
		{
			in: "select * from LogEvent.length_batch(10)",
			want: []Token{
				{lexer.SELECT, "select"},
				{lexer.ASTERISK, "*"},
				{lexer.FROM, "from"},
				{lexer.IDENT, "LogEvent"},
				{lexer.DOT, "."},
				{lexer.LENGTH_BATCH, "length_batch"},
				{lexer.LPAREN, "("},
				{lexer.INT, "10"},
				{lexer.RPAREN, ")"},
			},
		},
		{
			in: "select * from LogEvent.time_batch(10 sec)",
			want: []Token{
				{lexer.SELECT, "select"},
				{lexer.ASTERISK, "*"},
				{lexer.FROM, "from"},
				{lexer.IDENT, "LogEvent"},
				{lexer.DOT, "."},
				{lexer.TIME_BATCH, "time_batch"},
				{lexer.LPAREN, "("},
				{lexer.INT, "10"},
				{lexer.SEC, "sec"},
				{lexer.RPAREN, ")"},
			},
		},
	}

	for _, c := range cases {
		l := lexer.New(strings.NewReader(c.in))
		for _, w := range c.want {
			token, literal := l.Tokenize()
			if token != w.token || literal != w.literal {
				t.Errorf("got=%v:%v, want=%v:%v", token, literal, w.token, w.literal)
			}
		}

		if len(l.Errors()) != 0 {
			t.Errorf("errors=%v", l.Errors())
		}
	}
}
