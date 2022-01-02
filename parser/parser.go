package parser

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/itsubaki/gostream/lexer"
	"github.com/itsubaki/gostream/stream"
)

type Cursor struct {
	Token   lexer.Token
	Literal string
}

type Parser struct {
	l      *lexer.Lexer
	r      Registry
	opt    *Option
	cursor *Cursor
	peek   *Cursor
	errors []error
}

type Option struct {
	Verbose bool
}

type Registry map[string]interface{}

func (r Registry) Add(t interface{}) {
	r[reflect.TypeOf(t).Name()] = t
}

func New(opt ...*Option) *Parser {
	p := &Parser{
		r:      make(Registry),
		errors: make([]error, 0),
	}

	if len(opt) > 0 {
		p.opt = opt[0]
	}

	return p
}

func (p *Parser) Errors() []error {
	return p.errors
}

func (p *Parser) Add(t interface{}) *Parser {
	p.r.Add(t)
	return p
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

func (p *Parser) accept() interface{} {
	p.next()
	p.expect(lexer.IDENT)

	return p.r[p.cursor.Literal]
}

func (p *Parser) length() int64 {
	p.next()
	p.expect(lexer.LPAREN)
	defer func() {
		p.next()
		p.expect(lexer.RPAREN)
	}()

	p.next()
	p.expect(lexer.INT)

	v, err := strconv.ParseInt(p.cursor.Literal, 10, 64)
	if err != nil {
		p.errors = append(p.errors, err)
	}

	return v
}

func (p *Parser) time() (time.Duration, lexer.Token) {
	p.next()
	p.expect(lexer.LPAREN)
	defer func() {
		p.next()
		p.expect(lexer.RPAREN)
	}()

	p.next()
	p.expect(lexer.INT)

	v, err := strconv.ParseInt(p.cursor.Literal, 10, 64)
	if err != nil {
		p.errors = append(p.errors, err)
	}

	p.next()
	if p.cursor.Token == lexer.MIN {
		return time.Duration(v) * time.Minute, lexer.MIN
	}
	if p.cursor.Token == lexer.SEC {
		return time.Duration(v) * time.Second, lexer.SEC
	}

	return -1, lexer.EOF
}

func (p *Parser) Query(q string) *Parser {
	p.l = lexer.New(strings.NewReader(q))
	return p
}

func (p *Parser) Parse() *stream.Stream {
	s := stream.New()

	p.next() // preload
	for p.next().Token != lexer.EOF {
		switch p.cursor.Token {
		case lexer.SELECT:
		case lexer.FROM:
			s.Accept(p.accept())
		case lexer.LENGTH:
			s.Length(int(p.length()))
		case lexer.LENGTH_BATCH:
			s.LengthBatch(int(p.length()))
		case lexer.TIME:
			s.Time(p.time())
		case lexer.TIME_BATCH:
			s.TimeBatch(p.time())
		}
	}

	return s
}

func (p *Parser) String() string {
	var buf strings.Builder
	for k := range p.r {
		buf.WriteString(k)
	}

	return buf.String()
}
