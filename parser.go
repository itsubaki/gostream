package gocep

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"time"
)

type Statement struct {
	token    Token
	length   int
	unit     time.Duration
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

func (stmt *Statement) Selector(s Selector) {
	stmt.selector = append(stmt.selector, s)
}

func (stmt *Statement) Function(s Function) {
	stmt.function = append(stmt.function, s)
}

func (stmt *Statement) View(v View) {
	stmt.view = append(stmt.view, v)
}

func (stmt *Statement) New(capacity int) (w Window) {
	if stmt.token == LENGTH {
		w = NewLengthWindow(stmt.length, capacity)
	}

	for _, s := range stmt.selector {
		w.Selector(s)
	}

	for _, f := range stmt.function {
		w.Function(f)
	}

	for _, v := range stmt.view {
		w.View(v)
	}

	return w
}

type Parser struct {
	query string
	lexer *Lexer
}

func NewParser(query string) *Parser {
	return &Parser{
		query,
		NewLexer(strings.NewReader(query)),
	}
}

func (p *Parser) Parse() (*Statement, error) {
	stmt := NewStatement()

	// Select or InsertInto
	if token, literal := p.lexer.Tokenize(); token != SELECT {
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
			log.Println("add Function", token, literal)
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
			log.Println("add Selector", token, literal)
		}
	}

	// Window
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
		stmt.token = token
		stmt.length = length
		log.Println("add Window", token, literal, length)
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
