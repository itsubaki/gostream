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
	r      Registry
	errors []error
	cursor *Cursor
	peek   *Cursor
}

func NewParser(l *lexer.Lexer, r Registry) *Parser {
	return &Parser{
		l:      l,
		r:      r,
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
		"want={Token:%v, Literal: %v}, got={Token:%v, Literal: %v}",
		t, lexer.Tokens[t],
		p.cursor.Token, p.cursor.Literal,
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

func (p *Parser) Parse() Stream {
	s := NewStream()

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
