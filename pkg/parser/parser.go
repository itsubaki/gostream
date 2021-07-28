package parser

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/itsubaki/gostream/pkg/clause"
	"github.com/itsubaki/gostream/pkg/lexer"
	"github.com/itsubaki/gostream/pkg/statement"
)

type Registry map[string]interface{}

type Parser struct {
	Registry Registry
}

func New() *Parser {
	return &Parser{make(map[string]interface{})}
}

func (p *Parser) Register(name string, t interface{}) {
	p.Registry[name] = t
}

func (p *Parser) SetFunction(s *statement.Statement, fieldname string, intFunc, floatFunc clause.Function) error {
	if IsIntField(s.EventType, fieldname) {
		s.SetFunction(intFunc)
		return nil
	}

	if IsFloatField(s.EventType, fieldname) {
		s.SetFunction(floatFunc)
		return nil
	}

	return fmt.Errorf("invalid parameter event type=%v fieldname=%v", s.EventType, fieldname)
}

func (p *Parser) ParseFunction(s *statement.Statement, l *lexer.Lexer) error {
	for {
		token, literal := l.Tokenize()
		switch token {
		case lexer.EOF:
			return fmt.Errorf("invalid token=%s", literal)
		case lexer.FROM:
			return nil
		case lexer.ASTERISK:
			s.SetFunction(clause.SelectAll{})
		case lexer.COUNT:
			lp, lpl := l.Tokenize()
			ast, astl := l.Tokenize()
			rp, rpl := l.Tokenize()
			if lp != lexer.LPAREN || ast != lexer.ASTERISK || rp != lexer.RPAREN {
				return fmt.Errorf("invalid token=%s%s%s", lpl, astl, rpl)
			}

			s.SetFunction(clause.Count{As: "count(*)"})
		case lexer.MAX:
			_, name := l.TokenizeIdentifier()
			if err := p.SetFunction(s, name,
				clause.MaxInt{Name: name, As: fmt.Sprintf("max(%s)", name)},
				clause.MaxFloat{Name: name, As: fmt.Sprintf("max(%s)", name)},
			); err != nil {
				return fmt.Errorf("set function: %v", err)
			}
		case lexer.MIN:
			_, name := l.TokenizeIdentifier()
			if err := p.SetFunction(s, name,
				clause.MinInt{Name: name, As: fmt.Sprintf("min(%s)", name)},
				clause.MinFloat{Name: name, As: fmt.Sprintf("min(%s)", name)},
			); err != nil {
				return fmt.Errorf("set function: %v", err)
			}
		case lexer.MED:
			_, name := l.TokenizeIdentifier()
			if err := p.SetFunction(s, name,
				clause.MedianInt{Name: name, As: fmt.Sprintf("med(%s)", name)},
				clause.MedianFloat{Name: name, As: fmt.Sprintf("med(%s)", name)},
			); err != nil {
				return fmt.Errorf("set function: %v", err)
			}
		case lexer.SUM:
			_, name := l.TokenizeIdentifier()
			if err := p.SetFunction(s, name,
				clause.SumInt{Name: name, As: fmt.Sprintf("sum(%s)", name)},
				clause.SumFloat{Name: name, As: fmt.Sprintf("sum(%s)", name)},
			); err != nil {
				return fmt.Errorf("set function: %v", err)
			}
		case lexer.AVG:
			_, name := l.TokenizeIdentifier()
			if err := p.SetFunction(s, name,
				clause.AverageInt{Name: name, As: fmt.Sprintf("avg(%s)", name)},
				clause.AverageFloat{Name: name, As: fmt.Sprintf("avg(%s)", name)},
			); err != nil {
				return fmt.Errorf("set function: %v", err)
			}
		}
	}
}

func (p *Parser) ParseEventType(s *statement.Statement, l *lexer.Lexer) error {
	for {
		if token, _ := l.Tokenize(); token == lexer.FROM {
			break
		}
	}

	for {
		token, literal := l.Tokenize()
		switch token {
		case lexer.EOF:
			return fmt.Errorf("invalid token=%s", literal)
		case lexer.DOT:
			return nil
		case lexer.IDENTIFIER:
			if v, ok := p.Registry[literal]; ok {
				s.SetEventType(v)
				continue
			}

			return fmt.Errorf("EventType [%v] is not registered", literal)
		}
	}
}

func (p *Parser) ParseWindow(s *statement.Statement, l *lexer.Lexer) error {
	for {
		if token, _ := l.Tokenize(); token == lexer.DOT {
			break
		}
	}

	token, literal := l.Tokenize()
	if token == lexer.EOF {
		return fmt.Errorf("invalid token=%s", literal)
	}
	s.SetWindow(token)

	if token == lexer.LENGTH {
		_, lex := l.TokenizeIdentifier()
		length, err := strconv.Atoi(lex)
		if err != nil {
			return fmt.Errorf("atoi=%s: %v", lex, err)
		}

		s.SetLength(length)
		return nil
	}

	if token == lexer.TIME {
		_, lex := l.TokenizeIdentifier()
		d, err := strconv.Atoi(lex)
		if err != nil {
			return fmt.Errorf("atoi=%s: %v", lex, err)
		}

		t, _ := l.TokenizeIgnore(lexer.WHITESPACE)
		switch t {
		case lexer.SEC:
			s.SetTime(time.Duration(d) * time.Second)
		case lexer.MIN:
			s.SetTime(time.Duration(d) * time.Minute)
		}

		return nil
	}

	return fmt.Errorf("invalid token=%s", literal)
}

func (p *Parser) GetWhere(eventType interface{}, name, value string, t lexer.Token, l *lexer.Lexer) (clause.Where, error) {
	if IsIntField(eventType, name) {
		v, err := strconv.Atoi(value)
		if err != nil {
			return nil, fmt.Errorf("atoi=%v", v)
		}

		switch t {
		case lexer.LARGER:
			return clause.LargerThanInt{Name: name, Value: v}, nil
		case lexer.LESS:
			return clause.LessThanInt{Name: name, Value: v}, nil
		}
	}

	if IsFloatField(eventType, name) {
		_, val := l.TokenizeIdentifier()
		v, err := strconv.ParseFloat(fmt.Sprintf("%s.%s", value, val), 64)
		if err != nil {
			return nil, fmt.Errorf("parse float=%s.%s", value, val)
		}

		switch t {
		case lexer.LARGER:
			return clause.LargerThanFloat{Name: name, Value: v}, nil
		case lexer.LESS:
			return clause.LessThanFloat{Name: name, Value: v}, nil
		}
	}

	return nil, fmt.Errorf("invalid parameter event type=%v fieldname=%v", eventType, name)
}

func (p *Parser) ParseWhere(s *statement.Statement, l *lexer.Lexer) error {
	for {
		if token, _ := l.Tokenize(); token == lexer.DOT {
			break
		}
	}

	list := make([]clause.Where, 0)
	for {
		token, _ := l.Tokenize()
		if token == lexer.EOF {
			break
		}

		if token != lexer.WHERE && token != lexer.AND && token != lexer.OR {
			continue
		}

		_, name := l.TokenizeIdentifier()
		t, _ := l.TokenizeIgnore(lexer.WHITESPACE, lexer.IDENTIFIER)
		_, value := l.TokenizeIdentifier()

		w, err := p.GetWhere(s.EventType, name, value, t, l)
		if err != nil {
			return fmt.Errorf("get where: %v", err)
		}

		list = append(list, w)
	}

	s.SetWhere(list...)
	return nil
}

func (p *Parser) Parse(query string) (*statement.Statement, error) {
	s := statement.New()

	if token, literal := lexer.New(strings.NewReader(query)).Tokenize(); token != lexer.SELECT {
		return nil, fmt.Errorf("invalid token=%s", literal)
	}

	if err := p.ParseEventType(s, lexer.New(strings.NewReader(query))); err != nil {
		return nil, fmt.Errorf("parse event type: %v", err)
	}

	if err := p.ParseFunction(s, lexer.New(strings.NewReader(query))); err != nil {
		return nil, fmt.Errorf("parse function: %v", err)
	}

	if err := p.ParseWindow(s, lexer.New(strings.NewReader(query))); err != nil {
		return nil, fmt.Errorf("parse window: %v", err)
	}

	if err := p.ParseWhere(s, lexer.New(strings.NewReader(query))); err != nil {
		return nil, fmt.Errorf("parse selector: %v", err)
	}

	return s, nil
}
