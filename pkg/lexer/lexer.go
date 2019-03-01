package lexer

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
	WHITESPACE

	IDENTIFIER

	DOT
	COMMA
	SEMICOLON
	LPAREN
	RPAREN
	LBRACE
	RBRACE

	SELECT
	ASTERISK
	COUNT
	SUM
	AVG
	MAX
	MED

	FROM
	TIME
	LENGTH
	TIME_BATCH
	LENGTH_BATCH
	SEC
	MIN
	WHERE
	LARGER
	LESS
	AND
	OR
)

type Lexer struct {
	eof rune
	r   *bufio.Reader
}

func New(r io.Reader) *Lexer {
	return &Lexer{rune(0), bufio.NewReader(r)}
}

func (l *Lexer) TokenizeIgnoreWhiteSpace() (Token, string) {
	for {
		token, literal := l.Tokenize()
		if token == WHITESPACE {
			continue
		}
		return token, literal
	}
}

func (l *Lexer) TokenizeIgnoreIdentifier() (Token, string) {
	for {
		token, literal := l.TokenizeIgnoreWhiteSpace()
		if token == IDENTIFIER {
			continue
		}
		return token, literal
	}
}

func (l *Lexer) TokenizeIdentifier() (Token, string) {
	for {
		token, literal := l.TokenizeIgnoreWhiteSpace()
		if token != IDENTIFIER {
			continue
		}
		return token, literal
	}
}

func (l *Lexer) Tokenize() (Token, string) {
	ch := l.read()
	if l.isWhitespace(ch) {
		l.unread()
		return l.whitespace()
	}

	if l.isLetter(ch) || l.isDigit(ch) {
		l.unread()
		word := l.scan()
		return l.literal(word)
	}

	return l.symbol(ch)
}

func (l *Lexer) symbol(ch rune) (Token, string) {
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
	case '<':
		return LESS, string(ch)
	case '(':
		return LPAREN, string(ch)
	case ')':
		return RPAREN, string(ch)
	case '{':
		return LBRACE, string(ch)
	case '}':
		return RBRACE, string(ch)
	}

	return ILLEGAL, string(ch)
}

func (l *Lexer) literal(literal string) (Token, string) {
	switch strings.ToUpper(literal) {
	case "SELECT":
		return SELECT, literal
	case "COUNT":
		return COUNT, literal
	case "SUM":
		return SUM, literal
	case "AVG":
		return AVG, literal
	case "MAX":
		return MAX, literal
	case "MED":
		return MED, literal
	case "FROM":
		return FROM, literal
	case "WHERE":
		return WHERE, literal
	case "TIME":
		return TIME, literal
	case "LENGTH":
		return LENGTH, literal
	case "TIME_BATCH":
		return TIME_BATCH, literal
	case "LENGTH_BATCH":
		return LENGTH_BATCH, literal
	case "SEC":
		return SEC, literal
	case "MIN":
		return MIN, literal
	case "AND":
		return AND, literal
	case "OR":
		return OR, literal
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
		if !l.isLetter(ch) && !l.isDigit(ch) && ch != '_' {
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
		if !l.isWhitespace(ch) {
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

func (l *Lexer) isWhitespace(ch rune) bool {
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

func (l *Lexer) isLetter(ch rune) bool {
	if ch >= 'a' && ch <= 'z' {
		return true
	}
	if ch >= 'A' && ch <= 'Z' {
		return true
	}
	return false
}

func (l *Lexer) isDigit(ch rune) bool {
	if ch >= '0' && ch <= '9' {
		return true
	}
	return false
}
