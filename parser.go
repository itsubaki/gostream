package gocep

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Statement struct {
	window   Token
	length   int
	time     time.Duration
	selector []Selector
	function []Function
	view     []View
}

func NewStatement() *Statement {
	return &Statement{
		selector: []Selector{},
		function: []Function{},
		view:     []View{},
	}
}

func (stmt *Statement) SetSelector(s Selector) {
	stmt.selector = append(stmt.selector, s)
}

func (stmt *Statement) SetFunction(s Function) {
	stmt.function = append(stmt.function, s)
}

func (stmt *Statement) SetView(v View) {
	stmt.view = append(stmt.view, v)
}

func (stmt *Statement) New(capacity ...int) (w Window) {
	if stmt.window == LENGTH {
		w = NewLengthWindow(stmt.length, capacity...)
	}

	if stmt.window == TIME {
		w = NewTimeWindow(stmt.time, capacity...)
	}

	for _, s := range stmt.selector {
		w.SetSelector(s)
	}

	for _, f := range stmt.function {
		w.SetFunction(f)
	}

	for _, v := range stmt.view {
		w.SetView(v)
	}

	return w
}

type Registry map[string]interface{}

type Parser struct {
	registry Registry
}

func NewParser() *Parser {
	return &Parser{make(map[string]interface{})}
}

func (p *Parser) Register(name string, t interface{}) {
	p.registry[name] = t
}

func (p *Parser) ParseFunction(stmt *Statement, lexer *Lexer) error {
	for {
		token, literal := lexer.Tokenize()
		switch token {
		case EOF:
			return fmt.Errorf("invalid token=%s", literal)
		case FROM:
			return nil
		case ASTERISK:
			stmt.SetFunction(SelectAll{})
		case COUNT:
			stmt.SetFunction(Count{"count(*)"})
		}
	}
}

func (p *Parser) ParseSelector(stmt *Statement, lexer *Lexer) error {
	for {
		token, literal := lexer.Tokenize()
		switch token {
		case EOF:
			return fmt.Errorf("invalid token=%s", literal)
		case DOT:
			return nil
		case IDENTIFIER:
			v, ok := p.registry[literal]
			if !ok {
				return fmt.Errorf("EventType [%s] is not registered", literal)
			}
			stmt.SetSelector(EqualsType{v})
		}
	}
}

func (p *Parser) ParseWindow(stmt *Statement, lexer *Lexer) error {
	token, literal := lexer.Tokenize()
	if token == EOF {
		return fmt.Errorf("invalid token=%s", literal)
	}

	if token == LENGTH {
		stmt.window = token
		for {
			t, l := lexer.Tokenize()
			if t != IDENTIFIER {
				continue
			}

			length, err := strconv.Atoi(l)
			if err != nil {
				return fmt.Errorf("atoi=%s: %v", l, err)
			}

			stmt.length = length
			return nil
		}
	}

	if token == TIME {
		stmt.window = token
		for {
			t, l := lexer.Tokenize()
			if t != IDENTIFIER {
				continue
			}

			ct, err := strconv.Atoi(l)
			if err != nil {
				return fmt.Errorf("atoi=%s: %v", l, err)
			}

			for {
				t, _ := lexer.Tokenize()
				if t == SEC {
					stmt.time = time.Duration(ct) * time.Second
					return nil
				}
			}
		}
	}

	return fmt.Errorf("invalid token=%s", literal)
}

func (p *Parser) ParseWhere(stmt *Statement, lexer *Lexer) error {
	for {
		token, _ := lexer.Tokenize()
		if token == EOF {
			return nil
		}

		if token != WHERE && token != AND && token != OR {
			continue
		}

		var name, value string
		var selector Token

		for {
			t, l := lexer.Tokenize()
			if t != IDENTIFIER {
				continue
			}
			name = l
			break
		}

		for {
			t, _ := lexer.Tokenize()
			if t != LARGER && t != LESS {
				continue
			}
			selector = t
			break
		}

		for {
			t, l := lexer.Tokenize()
			if t != IDENTIFIER {
				continue
			}
			value = l
			break
		}

		val, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("atoi=%s", value)
		}

		switch selector {
		case LARGER:
			stmt.SetSelector(LargerThanInt{Name: name, Value: val})
		case LESS:
			stmt.SetSelector(LessThanInt{Name: name, Value: val})
		}
	}
}

func (p *Parser) Parse(query string) (*Statement, error) {
	stmt := NewStatement()
	lexer := NewLexer(strings.NewReader(query))

	if token, literal := lexer.Tokenize(); token != SELECT {
		return nil, fmt.Errorf("invalid token=%s", literal)
	}

	if err := p.ParseFunction(stmt, lexer); err != nil {
		return nil, fmt.Errorf("parse function: %v", err)
	}

	if err := p.ParseSelector(stmt, lexer); err != nil {
		return nil, fmt.Errorf("parse selector: %v", err)
	}

	if err := p.ParseWindow(stmt, lexer); err != nil {
		return nil, fmt.Errorf("parse window: %v", err)
	}

	if err := p.ParseWhere(stmt, lexer); err != nil {
		return nil, fmt.Errorf("parse where: %v", err)
	}

	return stmt, nil
}
