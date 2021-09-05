package parser

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/itsubaki/gostream/pkg/function"
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

func (p *Parser) SetFunction(s *statement.Statement, fieldname string, intFunc, floatFunc function.Function) error {
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
		token, str := l.Tokenize()
		switch token {
		case lexer.EOF:
			return fmt.Errorf("invalid token=%s", str)
		case lexer.FROM:
			return nil
		case lexer.ASTERISK:
			s.SetFunction(function.SelectAll{})
		case lexer.COUNT:
			lpar, str0 := l.Tokenize()
			asterisk, str1 := l.Tokenize()
			rpar, str2 := l.Tokenize()

			if lpar != lexer.LPAREN || asterisk != lexer.ASTERISK || rpar != lexer.RPAREN {
				return fmt.Errorf("invalid token=%s%s%s", str0, str1, str2)
			}

			s.SetFunction(function.Count{As: "count(*)"})
		case lexer.MAX:
			_, name := l.TokenizeIdent()
			if err := p.SetFunction(s, name,
				function.MaxInt{Name: name, As: fmt.Sprintf("max(%s)", name)},
				function.MaxFloat{Name: name, As: fmt.Sprintf("max(%s)", name)},
			); err != nil {
				return fmt.Errorf("set function: %v", err)
			}
		case lexer.MIN:
			_, name := l.TokenizeIdent()
			if err := p.SetFunction(s, name,
				function.MinInt{Name: name, As: fmt.Sprintf("min(%s)", name)},
				function.MinFloat{Name: name, As: fmt.Sprintf("min(%s)", name)},
			); err != nil {
				return fmt.Errorf("set function: %v", err)
			}
		case lexer.MED:
			_, name := l.TokenizeIdent()
			if err := p.SetFunction(s, name,
				function.MedianInt{Name: name, As: fmt.Sprintf("med(%s)", name)},
				function.MedianFloat{Name: name, As: fmt.Sprintf("med(%s)", name)},
			); err != nil {
				return fmt.Errorf("set function: %v", err)
			}
		case lexer.SUM:
			_, name := l.TokenizeIdent()
			if err := p.SetFunction(s, name,
				function.SumInt{Name: name, As: fmt.Sprintf("sum(%s)", name)},
				function.SumFloat{Name: name, As: fmt.Sprintf("sum(%s)", name)},
			); err != nil {
				return fmt.Errorf("set function: %v", err)
			}
		case lexer.AVG:
			_, name := l.TokenizeIdent()
			if err := p.SetFunction(s, name,
				function.AverageInt{Name: name, As: fmt.Sprintf("avg(%s)", name)},
				function.AverageFloat{Name: name, As: fmt.Sprintf("avg(%s)", name)},
			); err != nil {
				return fmt.Errorf("set function: %v", err)
			}
		}
	}
}

func (p *Parser) ParseEventType(s *statement.Statement, l *lexer.Lexer) error {
	for {
		if t, _ := l.Tokenize(); t == lexer.FROM {
			break
		}
	}

	for {
		token, str := l.Tokenize()
		switch token {
		case lexer.EOF:
			return fmt.Errorf("invalid token=%s", str)
		case lexer.DOT:
			return nil
		case lexer.IDENTIFIER:
			if v, ok := p.Registry[str]; ok {
				s.SetEventType(v)
				continue
			}

			return fmt.Errorf("EventType [%v] is not registered", str)
		}
	}
}

func (p *Parser) ParseWindow(s *statement.Statement, l *lexer.Lexer) error {
	for {
		if t, _ := l.Tokenize(); t == lexer.DOT {
			break
		}
	}

	token, str := l.Tokenize()
	if token == lexer.EOF {
		return fmt.Errorf("invalid token=%s", str)
	}
	s.SetWindow(token)

	if token == lexer.LENGTH {
		_, value := l.TokenizeNumber()
		length, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("atoi string=%s: %v", value, err)
		}

		s.SetLength(length)
		return nil
	}

	if token == lexer.TIME {
		_, value := l.TokenizeNumber()
		d, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("atoi string=%s: %v", value, err)
		}

		unit, _ := l.Tokenize()
		switch unit {
		case lexer.SEC:
			s.SetTime(time.Duration(d) * time.Second)
		case lexer.MIN:
			s.SetTime(time.Duration(d) * time.Minute)
		}

		return nil
	}

	return fmt.Errorf("invalid token=%s", str)
}

func (p *Parser) GetWhere(eventType interface{}, name, value string, ope lexer.Token, l *lexer.Lexer) (function.Where, error) {
	if IsIntField(eventType, name) {
		v, err := strconv.Atoi(value)
		if err != nil {
			return nil, fmt.Errorf("atoi string=%v", v)
		}

		switch ope {
		case lexer.LARGER:
			return function.LargerThanInt{Name: name, Value: v}, nil
		case lexer.LESS:
			return function.LessThanInt{Name: name, Value: v}, nil
		}
	}

	if IsFloatField(eventType, name) {
		v, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, fmt.Errorf("parse string=%s", value)
		}

		switch ope {
		case lexer.LARGER:
			return function.LargerThanFloat{Name: name, Value: v}, nil
		case lexer.LESS:
			return function.LessThanFloat{Name: name, Value: v}, nil
		}
	}

	return nil, fmt.Errorf("invalid parameter event type=%v fieldname=%v", eventType, name)
}

func (p *Parser) ParseWhere(s *statement.Statement, l *lexer.Lexer) error {
	for {
		if t, _ := l.Tokenize(); t == lexer.DOT {
			break
		}
	}

	list := make([]function.Where, 0)
	for {
		token, _ := l.Tokenize()
		if token == lexer.EOF {
			break
		}

		if token != lexer.WHERE && token != lexer.AND && token != lexer.OR {
			continue
		}

		_, name := l.TokenizeIdent()
		ope, _ := l.Tokenize()
		_, value := l.TokenizeNumber()

		w, err := p.GetWhere(s.EventType, name, value, ope, l)
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

	if token, str := lexer.New(strings.NewReader(query)).Tokenize(); token != lexer.SELECT {
		return nil, fmt.Errorf("invalid token=%s", str)
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
