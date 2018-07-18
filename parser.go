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

func (p *Parser) Parse(query string) (*Statement, error) {
	stmt := NewStatement()
	lexer := NewLexer(strings.NewReader(query))

	// Select or InsertInto
	if token, literal := lexer.Tokenize(); token != SELECT {
		return nil, fmt.Errorf("invalid token=%s", literal)
	}

	// Function
	for {
		token, literal := lexer.Tokenize()
		switch token {
		case EOF:
			return nil, fmt.Errorf("invalid token=%s", literal)
		case ASTERISK:
			stmt.SetFunction(SelectAll{})
		case COUNT:
			stmt.SetFunction(Count{"count(*)"})
		}

		if token == FROM {
			break
		}
	}

	// Selector
	for {
		token, literal := lexer.Tokenize()
		switch token {
		case EOF:
			return nil, fmt.Errorf("invalid token=%s", literal)
		case IDENTIFIER:
			if v, ok := p.registry[literal]; ok {
				stmt.SetSelector(EqualsType{v})
			} else {
				return nil, fmt.Errorf("EventType [%s] is not registered", literal)
			}
		}

		if token == DOT {
			break
		}

	}

	// Window
	token, literal := lexer.Tokenize()
	if token == EOF {
		return nil, fmt.Errorf("invalid token=%s", literal)
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
				return nil, fmt.Errorf("atoi=%s: %v", l, err)
			}

			stmt.length = length
			break
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
				return nil, fmt.Errorf("atoi=%s: %v", l, err)
			}

			for {
				t, _ := lexer.Tokenize()
				if t == SEC {
					stmt.time = time.Duration(ct) * time.Second
					break
				}
			}

			break
		}
	}

	// Where
	for {
		token, _ := lexer.Tokenize()
		if token == EOF {
			break
		}

		if token != WHERE {
			continue
		}

		var name, value string
		var selector Token

		for {
			t, l := lexer.Tokenize()
			if t == IDENTIFIER {
				name = l
				break
			}
		}

		for {
			t, _ := lexer.Tokenize()
			if t == LARGER || t == LESS {
				selector = t
				break
			}
		}

		for {
			t, l := lexer.Tokenize()
			if t == IDENTIFIER {
				value = l
				break
			}
		}

		val, err := strconv.Atoi(value)
		if err != nil {
			return nil, fmt.Errorf("atoi=%s", value)
		}

		switch selector {
		case LARGER:
			stmt.SetSelector(LargerThanInt{Name: name, Value: val})
		case LESS:
			stmt.SetSelector(LessThanInt{Name: name, Value: val})
		}
	}

	return stmt, nil
}
