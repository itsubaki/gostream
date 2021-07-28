package lexer

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

type Lexer struct {
	eof     rune
	r       *bufio.Reader
	Symbol  map[rune]Token
	Literal map[string]Token
}

func New(r io.Reader) *Lexer {
	lex := &Lexer{
		eof:     rune(-1),
		r:       bufio.NewReader(r),
		Symbol:  make(map[rune]Token),
		Literal: make(map[string]Token),
	}

	lex.Symbol['*'] = ASTERISK
	lex.Symbol[','] = COMMA
	lex.Symbol['.'] = DOT
	lex.Symbol['>'] = LARGER
	lex.Symbol['<'] = LESS
	lex.Symbol['('] = LPAREN
	lex.Symbol[')'] = RPAREN
	lex.Symbol['{'] = LBRACE
	lex.Symbol['}'] = RBRACE

	lex.Literal["SELECT"] = SELECT
	lex.Literal["COUNT"] = COUNT
	lex.Literal["SUM"] = SUM
	lex.Literal["AVG"] = AVG
	lex.Literal["MAX"] = MAX
	lex.Literal["MED"] = MED
	lex.Literal["FROM"] = FROM
	lex.Literal["WHERE"] = WHERE
	lex.Literal["TIME"] = TIME
	lex.Literal["LENGTH"] = LENGTH
	lex.Literal["TIME_BATCH"] = TIME_BATCH
	lex.Literal["LENGTH_BATCH"] = LENGTH_BATCH
	lex.Literal["SEC"] = SEC
	lex.Literal["MIN"] = MIN
	lex.Literal["AND"] = AND
	lex.Literal["OR"] = OR

	return lex
}

func (l *Lexer) TokenizeIgnore(t ...Token) (Token, string) {
	ignore := make(map[Token]bool)
	for _, tt := range t {
		ignore[tt] = true
	}

	for {
		token, str := l.Tokenize()
		if _, ok := ignore[token]; ok {
			continue
		}

		return token, str
	}
}

func (l *Lexer) TokenizeIdentifier() (Token, string) {
	for {
		token, str := l.TokenizeIgnore(WHITESPACE)
		if token != IDENTIFIER {
			continue
		}

		return token, str
	}
}

func (l *Lexer) Tokenize() (Token, string) {
	ch := l.read()
	if isWhitespace(ch) {
		l.unread()
		return l.whitespace()
	}

	if isLetter(ch) || isDigit(ch) {
		l.unread()
		word := l.scan()
		return l.literal(word)
	}

	return l.symbol(ch)
}

func (l *Lexer) symbol(ch rune) (Token, string) {
	if ch == l.eof {
		return EOF, ""
	}

	v, ok := l.Symbol[ch]
	if ok {
		return v, string(ch)
	}

	return ILLEGAL, string(ch)
}

func (l *Lexer) literal(literal string) (Token, string) {
	v, ok := l.Literal[strings.ToUpper(literal)]
	if ok {
		return v, literal
	}

	return IDENTIFIER, literal
}

func (l *Lexer) scan() string {
	var buf bytes.Buffer
	buf.WriteRune(l.read())

	for {
		ch := l.read()
		if ch == l.eof {
			break
		}

		if !isLetter(ch) && !isDigit(ch) {
			l.unread()
			break
		}

		_, _ = buf.WriteRune(ch)
	}

	return buf.String()
}

func (l *Lexer) whitespace() (Token, string) {
	var buf bytes.Buffer
	buf.WriteRune(l.read())

	for {
		ch := l.read()
		if ch == l.eof {
			break
		}

		if !isWhitespace(ch) {
			l.unread()
			break
		}

		buf.WriteRune(ch)
	}

	return WHITESPACE, buf.String()
}

func (l *Lexer) read() rune {
	ch, _, err := l.r.ReadRune()
	if err != nil {
		return l.eof
	}

	return ch
}

func (l *Lexer) unread() {
	_ = l.r.UnreadRune()
}

func isWhitespace(ch rune) bool {
	if ch == ' ' {
		return true
	}
	if ch == '\t' {
		return true
	}
	if ch == '\n' {
		return true
	}

	return false
}

func isLetter(ch rune) bool {
	if ch >= 'a' && ch <= 'z' {
		return true
	}
	if ch >= 'A' && ch <= 'Z' {
		return true
	}

	return false
}

func isDigit(ch rune) bool {
	if ch >= '0' && ch <= '9' {
		return true
	}

	return false
}
