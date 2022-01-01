package gostream

import (
	"fmt"
	"time"

	"github.com/itsubaki/gostream/lexer"
)

type Cursor struct {
	Token   lexer.Token
	Literal string
}

type Parser struct {
	l      *lexer.Lexer
	errors []error
	cursor *Cursor
	peek   *Cursor
}

func NewParser(l *lexer.Lexer) *Parser {
	return &Parser{
		l:      l,
		errors: make([]error, 0),
	}
}

func (p *Parser) Errors() []error {
	return p.errors
}

func (p *Parser) error(e error) {
	p.errors = append(p.errors, e)
}

func (p *Parser) expect(t lexer.Token) {
	if p.cursor.Token == t {
		return
	}

	p.error(fmt.Errorf(
		"got={Token:%v, Literal: %v}, want={Token:%v, Literal: %v}",
		p.cursor.Token, p.cursor.Literal,
		t, lexer.Tokens[t],
	))
}

func (p *Parser) next() *Cursor {
	token, literal := p.l.Tokenize()

	p.cursor = p.peek
	p.peek = &Cursor{
		Token:   token,
		Literal: literal,
	}

	return p.cursor
}

func (p *Parser) parseAcceptType() interface{} {
	return nil
}

func (p *Parser) parseLength() int {
	return 10
}

func (p *Parser) parseTimeDuration() time.Duration {
	return 10 * time.Minute
}

func Parse(l *lexer.Lexer) Stream {
	s := NewStream()

	p := NewParser(l)
	p.next() // preload
	for p.next().Token != lexer.EOF {
		switch p.cursor.Token {
		case lexer.SELECT:
		case lexer.FROM:
			s.Accept(p.parseAcceptType())
		case lexer.LENGTH:
			s.Length(p.parseLength())
		case lexer.LENGTH_BATCH:
			s.LengthBatch(p.parseLength())
		case lexer.TIME:
			s.Time(p.parseTimeDuration())
		case lexer.TIME_BATCH:
			s.TimeBatch(p.parseTimeDuration())
		case lexer.WHERE:
		}
	}

	return s
}
