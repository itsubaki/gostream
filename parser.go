package gocep

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Parser struct {
	query    string
	capacity int
	lexer    *Lexer
}

type Statement struct {
	window   Window
	selector []Selector
	function []Function
	view     []View
}

func NewParser(query string, capacity int) *Parser {
	return &Parser{query, capacity, NewLexer(strings.NewReader(query))}
}

func (p *Parser) Parse() (*Statement, error) {
	stmt := &Statement{nil, []Selector{}, []Function{}, []View{}}

	// Select or InsertInto
	token, literal := p.lexer.Tokenize()
	if token != SELECT {
		return nil, errors.New("invalid token. token: " + literal)
	}

	// Function
	for {
		token, literal := p.lexer.Tokenize()
		if token == EOF {
			break
		}
		if token == FROM {
			break
		}
		if token == LITERAL {
			stmt.function = append(stmt.function, SelectMapString{"Record", literal, literal})
			fmt.Println("add", token, literal)
		}
	}

	// Selector
	for {
		token, literal := p.lexer.Tokenize()
		if token == EOF {
			break
		}
		if token == DOT {
			break
		}
		if token == LITERAL {
			stmt.selector = append(stmt.selector, EqualsType{MapEvent{}})
			fmt.Println("add", token, literal)
		}
	}

	// Window
	for {
		token, literal := p.lexer.Tokenize()
		if token == EOF {
			break
		}
		if token == LENGTH {
			length := 0
			for {
				t, l := p.lexer.Tokenize()
				if t == LITERAL {
					length, _ = strconv.Atoi(l)
					break
				}
			}
			stmt.window = NewLengthWindow(length, p.capacity)
			fmt.Println("add", token, literal)
			break
		}
	}

	for {
		token, _ := p.lexer.Tokenize()
		if token == EOF {
			break
		}
		if token == WHERE {
			break
		}
	}

	// Selector
	for {
		token, literal := p.lexer.Tokenize()
		if token == EOF {
			break
		}
		if token == WHITESPACE {
			continue
		}
		fmt.Println(token, literal)
	}

	return stmt, nil
}
