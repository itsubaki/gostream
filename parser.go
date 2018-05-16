package gocep

import (
	"errors"
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

func (p *Parser) Parse(query string) (*Statement, error) {
	stmt := NewStatement()
	lexer := NewLexer(strings.NewReader(query))

	// Select or InsertInto
	if token, literal := lexer.Tokenize(); token != SELECT {
		return nil, errors.New("invalid token. literal: " + literal)
	}

	// Function
	for {
		token, literal := lexer.Tokenize()
		if token == EOF {
			return nil, errors.New("invalid token. literal: " + literal)
		}
		if token == FROM {
			break
		}
		if token == ASTERISK {
			stmt.SetFunction(SelectAll{})
		}

		if token == COUNT {
			stmt.SetFunction(Count{"count(*)"})
		}
	}

	// Selector
	for {
		token, literal := lexer.Tokenize()
		if token == EOF {
			return nil, errors.New("invalid token. literal: " + literal)
		}
		if token == DOT {
			break
		}
		if token == IDENTIFIER {
			v, ok := p.registry[literal]
			if !ok {
				return nil, fmt.Errorf("EventType [%s] is not registered", literal)
			}
			stmt.SetSelector(EqualsType{v})
		}
	}

	// Window
	token, literal := lexer.Tokenize()
	if token == EOF {
		return nil, errors.New("invalid token. literal: " + literal)
	}

	if token == LENGTH {
		stmt.window = token

		length := 0
		for {
			t, l := lexer.Tokenize()
			if t == IDENTIFIER {
				length, _ = strconv.Atoi(l)
				break
			}
		}
		stmt.length = length
	}

	if token == TIME {
		stmt.window = token

		var dt time.Duration
		for {
			t, l := lexer.Tokenize()
			if t == IDENTIFIER {
				ct, err := strconv.Atoi(l)
				if err != nil {
					return nil, errors.New("invalid token. literal: " + l)
				}
				dt = time.Duration(ct)
				break
			}
		}

		for {
			t, l := lexer.Tokenize()
			if t == SEC {
				if l == "sec" {
					stmt.time = dt * time.Second
					break
				}
			}
		}
	}

	// Where
	for {
		token, _ := lexer.Tokenize()
		if token == EOF {
			return stmt, nil
		}

		if token == WHERE {
			for {
				t, _ := lexer.Tokenize()
				if t == IDENTIFIER {
					break
				}
			}
			break
		}
	}

	return stmt, nil
}
