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
	if p.cursor.Token == lexer.SEC {
		return time.Duration(v) * time.Second, lexer.SEC
	}
	if p.cursor.Token == lexer.MIN {
		return time.Duration(v) * time.Minute, lexer.MIN
	}
	if p.cursor.Token == lexer.HOUR {
		return time.Duration(v) * time.Hour, lexer.HOUR
	}

	return -1, lexer.EOF
}

func (p *Parser) limit() (int, int) {
	p.next()
	p.expect(lexer.INT)

	l, err := strconv.Atoi(p.cursor.Literal)
	if err != nil {
		p.errors = append(p.errors, err)
	}

	p.next()
	if p.cursor.Token == lexer.OFFSET {
		p.next()
		p.expect(lexer.INT)

		o, err := strconv.Atoi(p.cursor.Literal)
		if err != nil {
			p.errors = append(p.errors, err)
		}

		return l, o
	}

	return l, 0
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
			for p.next().Token != lexer.FROM {
				if p.cursor.Token == lexer.ASTERISK {
					s.SelectAll()
					continue
				}

				if p.cursor.Token == lexer.IDENT {
					s.Select(p.cursor.Literal)
					continue
				}

				if p.cursor.Token == lexer.AVG {
					p.next()
					p.expect(lexer.LPAREN)
					s.Average(p.next().Literal)
					p.next()
					p.expect(lexer.RPAREN)
					continue
				}

				if p.cursor.Token == lexer.SUM {
					p.next()
					p.expect(lexer.LPAREN)
					s.Sum(p.next().Literal)
					p.next()
					p.expect(lexer.RPAREN)
					continue
				}

				if p.cursor.Token == lexer.COUNT {
					p.next()
					p.expect(lexer.LPAREN)
					s.Count(p.next().Literal)
					p.next()
					p.expect(lexer.RPAREN)
					continue
				}

				if p.cursor.Token == lexer.MAX {
					p.next()
					p.expect(lexer.LPAREN)
					s.Max(p.next().Literal)
					p.next()
					p.expect(lexer.RPAREN)
					continue
				}

				if p.cursor.Token == lexer.MIN {
					p.next()
					p.expect(lexer.LPAREN)
					s.Min(p.next().Literal)
					p.next()
					p.expect(lexer.RPAREN)
					continue
				}

				if p.cursor.Token == lexer.DISTINCT {
					p.next()
					p.expect(lexer.LPAREN)
					s.Distinct(p.next().Literal)
					p.next()
					p.expect(lexer.RPAREN)
					continue
				}
			}

			p.next()
			p.expect(lexer.IDENT)

			s.From(p.r[p.cursor.Literal])
		case lexer.LENGTH:
			s.Length(int(p.length()))
		case lexer.LENGTH_BATCH:
			s.LengthBatch(int(p.length()))
		case lexer.TIME:
			s.Time(p.time())
		case lexer.TIME_BATCH:
			s.TimeBatch(p.time())
		case lexer.ORDER_BY:
			p.next()
			p.expect(lexer.IDENT)
			v := p.cursor.Literal

			p.next()
			s.OrderBy(v, p.cursor.Token == lexer.DESC)

			// limit
			if p.cursor.Token == lexer.DESC {
				p.next()
			}
			if p.cursor.Token != lexer.LIMIT {
				continue
			}

			s.Limit(p.limit())
		case lexer.LIMIT:
			s.Limit(p.limit())
		case lexer.WHERE:
			p.next()
			p.expect(lexer.IDENT)
			name := p.cursor.Literal

			// >, <, =
			op := p.next()

			p.next()
			var value interface{}
			value = p.cursor.Literal // string
			if p.cursor.Token == lexer.INT {
				v, err := strconv.Atoi(p.cursor.Literal)
				if err != nil {
					p.errors = append(p.errors, err)
				}
				value = v
			}

			if p.cursor.Token == lexer.FLOAT {
				v, err := strconv.ParseFloat(p.cursor.Literal, 64)
				if err != nil {
					p.errors = append(p.errors, err)
				}
				value = v
			}

			if op.Token == lexer.LARGER {
				s.LargerThan(name, value)
			}

			if op.Token == lexer.LESS {
				s.LessThan(name, value)
			}

			if op.Token == lexer.EQUALS {
				s.Equals(name, value)
			}
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
