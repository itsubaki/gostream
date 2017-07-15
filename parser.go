package gocep

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Statement struct {
	window   Window
	selector []Selector
	function []Function
	view     []View
}

func NewStatement() *Statement {
	return &Statement{nil, []Selector{}, []Function{}, []View{}}
}

func (stmt *Statement) Window(w Window) {
	stmt.window = w
}

func (stmt *Statement) Selector(s Selector) {
	stmt.selector = append(stmt.selector, s)
}

func (stmt *Statement) Function(s Function) {
	stmt.function = append(stmt.function, s)
}

func (stmt *Statement) View(v View) {
	stmt.view = append(stmt.view, v)
}

func (stmt *Statement) Build() Window {
	for _, s := range stmt.selector {
		stmt.window.Selector(s)
	}

	for _, f := range stmt.function {
		stmt.window.Function(f)
	}

	for _, v := range stmt.view {
		stmt.window.View(v)
	}
	return stmt.window
}

type Parser struct {
	query    string
	capacity int
	lexer    *Lexer
}

func NewParser(query string, capacity int) *Parser {
	return &Parser{
		query,
		capacity,
		NewLexer(strings.NewReader(query)),
	}
}

func (p *Parser) Parse() (*Statement, error) {
	stmt := NewStatement()

	// Select or InsertInto
	token, literal := p.lexer.Tokenize()
	if token != SELECT {
		return nil, errors.New("invalid token. literal: " + literal)
	}

	// Function
	for {
		token, literal := p.lexer.Tokenize()
		if token == EOF {
			return nil, errors.New("invalid token. literal: " + literal)
		}
		if token == FROM {
			break
		}
		if token == ASTERISK {
			stmt.Function(SelectMapAll{"Record"})
			fmt.Println("add Function", token, literal)
		}
	}

	// Selector
	for {
		token, literal := p.lexer.Tokenize()
		if token == EOF {
			return nil, errors.New("invalid token. literal: " + literal)
		}
		if token == DOT {
			break
		}
		if token == IDENTIFIER {
			stmt.Selector(EqualsType{MapEvent{}})
			fmt.Println("add Selector", token, literal)
		}
	}

	// Window
	for {
		token, literal := p.lexer.Tokenize()
		if token == EOF {
			return nil, errors.New("invalid token. literal: " + literal)
		}
		if token == LENGTH {
			length := 0
			for {
				t, l := p.lexer.Tokenize()
				if t == IDENTIFIER {
					length, _ = strconv.Atoi(l)
					break
				}
			}
			stmt.Window(NewLengthWindow(length, p.capacity))
			fmt.Println("add Window", token, literal)
			break
		}
	}

	for {
		token, _ := p.lexer.Tokenize()
		if token == EOF {
			return stmt, nil
		}
		if token == WHERE {
			break
		}
	}

	return stmt, nil
}
