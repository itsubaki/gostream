package gostream

import (
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

func (p *Parser) next() *Cursor {
	token, literal := p.l.Tokenize()

	p.cursor = p.peek
	p.peek = &Cursor{
		Token:   token,
		Literal: literal,
	}

	return p.cursor
}

func Parse(l *lexer.Lexer) Stream {
	s := NewStream()
	p := NewParser(l)

	p.next() // preload
	for p.next().Token != lexer.EOF {
		switch p.cursor.Token {
		case lexer.SELECT:
		case lexer.FROM:
			s.Accept(Cursor{})
		case lexer.LENGTH:
			s.Length(10)
		case lexer.LENGTH_BATCH:
			s.LengthBatch(10)
		case lexer.TIME:
			s.Time(10 * time.Minute)
		case lexer.TIME_BATCH:
			s.TimeBatch(10 * time.Minute)
		}
	}

	return s
}
