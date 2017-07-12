package gocep

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

type Token int

const (
	ILLEGAL Token = iota
	EOF
	LITERAL
	WHITESPACE
	ASTERISK
	DOT
	COMMA
	SELECT
	FROM
	TIME
	SEC
	WHERE
	LARGER
)

type Lexer struct {
	eof rune
	r   *bufio.Reader
}

func NewLexer(r io.Reader) *Lexer {
	return &Lexer{rune(0), bufio.NewReader(r)}
}

func (l *Lexer) Tokenize() (tok Token, lit string) {
	ch := l.read()
	if l.isWhitespace(ch) {
		l.unread()
		return l.whitespace()
	}

	if l.isLetter(ch) || l.isDigit(ch) {
		l.unread()
		return l.scan()
	}

	switch ch {
	case l.eof:
		return EOF, ""
	case '*':
		return ASTERISK, string(ch)
	case ',':
		return COMMA, string(ch)
	case '.':
		return DOT, string(ch)
	case '>':
		return LARGER, string(ch)
	}

	return ILLEGAL, string(ch)
}

func (l *Lexer) scan() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(l.read())

	for {
		ch := l.read()
		if ch == l.eof {
			break
		}
		if !l.isLetter(ch) && !l.isDigit(ch) && ch != '_' {
			l.unread()
			break
		}
		_, _ = buf.WriteRune(ch)
	}

	switch strings.ToUpper(buf.String()) {
	case "SELECT":
		return SELECT, buf.String()
	case "FROM":
		return FROM, buf.String()
	case "WHERE":
		return WHERE, buf.String()
	case "TIME":
		return TIME, buf.String()
	case "SEC":
		return SEC, buf.String()
	}

	return LITERAL, buf.String()
}

func (l *Lexer) whitespace() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(l.read())

	for {
		ch := l.read()
		if ch == l.eof {
			break
		}
		if l.isWhitespace(ch) {
			buf.WriteRune(ch)
			continue
		}
		l.unread()
		break
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

func (l *Lexer) isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func (l *Lexer) isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func (l *Lexer) isDigit(ch rune) bool {
	return (ch >= '0' && ch <= '9')
}
