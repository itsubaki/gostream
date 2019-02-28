package parser

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/itsubaki/gocep/pkg/function"
	"github.com/itsubaki/gocep/pkg/lexer"
	"github.com/itsubaki/gocep/pkg/selector"
	"github.com/itsubaki/gocep/pkg/statement"
)

type Registry map[string]interface{}

type Parser struct {
	registry Registry
}

func New() *Parser {
	return &Parser{make(map[string]interface{})}
}

func (p *Parser) Register(name string, t interface{}) {
	p.registry[name] = t
}

func (p *Parser) ParseFunction(st *statement.Statement, l *lexer.Lexer) error {
	for {
		token, literal := l.Tokenize()
		switch token {
		case lexer.EOF:
			return fmt.Errorf("invalid token=%s", literal)
		case lexer.FROM:
			return nil
		case lexer.ASTERISK:
			st.SetFunction(function.SelectAll{})
		case lexer.COUNT:
			st.SetFunction(function.Count{As: "count(*)"})
		}
	}
}

func (p *Parser) ParseEventType(st *statement.Statement, l *lexer.Lexer) error {
	for {
		token, literal := l.Tokenize()
		switch token {
		case lexer.EOF:
			return fmt.Errorf("invalid token=%s", literal)
		case lexer.DOT:
			return nil
		case lexer.IDENTIFIER:
			v, ok := p.registry[literal]
			if !ok {
				return fmt.Errorf("EventType [%s] is not registered", literal)
			}
			st.SetSelector(selector.EqualsType{Accept: v})
		}
	}
}

func (p *Parser) ParseWindow(st *statement.Statement, l *lexer.Lexer) error {
	token, literal := l.Tokenize()
	if token == lexer.EOF {
		return fmt.Errorf("invalid token=%s", literal)
	}

	if token == lexer.LENGTH {
		st.SetWindow(token)

		_, lex := l.TokenizeIdentifier()
		length, err := strconv.Atoi(lex)
		if err != nil {
			return fmt.Errorf("atoi=%s: %v", lex, err)
		}

		st.SetLength(length)
		return nil
	}

	if token == lexer.TIME {
		st.SetWindow(token)

		_, lex := l.TokenizeIdentifier()
		ct, err := strconv.Atoi(lex)
		if err != nil {
			return fmt.Errorf("atoi=%s: %v", lex, err)
		}

		t, _ := l.TokenizeIgnoreWhiteSpace()
		switch t {
		case lexer.SEC:
			st.SetTime(time.Duration(ct) * time.Second)
		case lexer.MIN:
			st.SetTime(time.Duration(ct) * time.Minute)
		}

		return nil
	}

	return fmt.Errorf("invalid token=%s", literal)
}

func (p *Parser) ParseSelector(st *statement.Statement, l *lexer.Lexer) error {
	for {
		token, _ := l.Tokenize()
		if token == lexer.EOF {
			return nil
		}

		if token != lexer.WHERE && token != lexer.AND && token != lexer.OR {
			continue
		}

		_, name := l.TokenizeIdentifier()
		s, _ := l.TokenizeIgnoreIdentifier()
		_, value := l.TokenizeIdentifier()

		val, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("atoi=%s", value)
		}

		switch s {
		case lexer.LARGER:
			st.SetSelector(selector.LargerThanInt{Name: name, Value: val})
		case lexer.LESS:
			st.SetSelector(selector.LessThanInt{Name: name, Value: val})
		}
	}
}

func (p *Parser) Parse(query string) (*statement.Statement, error) {
	st := statement.New()

	l := lexer.New(strings.NewReader(query))
	if token, literal := l.Tokenize(); token != lexer.SELECT {
		return nil, fmt.Errorf("invalid token=%s", literal)
	}

	if err := p.ParseFunction(st, l); err != nil {
		return nil, fmt.Errorf("parse function: %v", err)
	}

	if err := p.ParseEventType(st, l); err != nil {
		return nil, fmt.Errorf("parse event type: %v", err)
	}

	if err := p.ParseWindow(st, l); err != nil {
		return nil, fmt.Errorf("parse window: %v", err)
	}

	if err := p.ParseSelector(st, l); err != nil {
		return nil, fmt.Errorf("parse selector: %v", err)
	}

	return st, nil
}
