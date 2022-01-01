package gostream

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/itsubaki/gostream/lexer"
)

type GoStream struct {
	Registry map[string]interface{}
	Option   *Option
	Stream   Stream
}

type Option struct {
	Verbose bool
}

func New(opt ...*Option) *GoStream {
	s := &GoStream{
		Registry: make(map[string]interface{}),
	}

	if len(opt) > 0 {
		s.Option = opt[0]
	}

	return s
}

func (s *GoStream) Add(t ...interface{}) *GoStream {
	for i := range t {
		name := reflect.TypeOf(t[i]).Name()
		s.Registry[name] = t[i]
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

	s.Stream = Parse(lexer.New(strings.NewReader(q)))
	go s.Stream.Run()

	return s, nil
}

func (s *GoStream) Close() error {
	return s.Stream.Close()
}
