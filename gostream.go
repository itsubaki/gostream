package gostream

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/itsubaki/gostream/lexer"
)

type GoStream struct {
	Registry Registry
	Option   *Option
	Stream   Stream
}

type Registry map[string]interface{}

func (r Registry) Add(t interface{}) {
	r[reflect.TypeOf(t).Name()] = t
}

type Option struct {
	Verbose bool
}

func New(opt ...*Option) *GoStream {
	s := &GoStream{
		Registry: make(Registry),
	}

	if len(opt) > 0 {
		s.Option = opt[0]
	}

	return s
}

func (s *GoStream) Add(t ...interface{}) *GoStream {
	for i := range t {
		s.Registry.Add(t[i])
	}

	return s
}

func (s *GoStream) Query(q string) (*GoStream, error) {
	if s.Option.Verbose {
		l := lexer.New(strings.NewReader(q))
		for {
			tok, lit := l.Tokenize()
			if tok == lexer.EOF {
				break
			}

			fmt.Printf("%v", lexer.Tokens[tok])
			if lexer.IsBasicLit(tok) {
				fmt.Printf("(%v) ", lit)
			} else {
				fmt.Printf(" ")
			}
		}
	}

	p := NewParser(lexer.New(strings.NewReader(q)), s.Registry)
	s.Stream = p.Parse()
	if len(p.errors) > 0 {
		return nil, fmt.Errorf("parse: %v", p.errors)
	}
	go s.Stream.Run()

	return s, nil
}

func (s *GoStream) Close() error {
	return s.Stream.Close()
}
