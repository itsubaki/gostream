package gostream

import (
	"fmt"
	"strings"

	"github.com/itsubaki/gostream/lexer"
	"github.com/itsubaki/gostream/parser"
	"github.com/itsubaki/gostream/stream"
)

type GoStream struct {
	opt *Option
	r   parser.Registry
}

type Option struct {
	Verbose bool
}

func New(opt ...*Option) *GoStream {
	s := &GoStream{
		r: make(parser.Registry),
	}

	if len(opt) > 0 {
		s.opt = opt[0]
	}

	return s
}

func (s *GoStream) Add(t interface{}) *GoStream {
	s.r.Add(t)
	return s
}

func (s *GoStream) Query(q string) (*stream.Stream, error) {
	if s.opt.Verbose {
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

	p := parser.New().Query(q)
	for k := range s.r {
		p.Add(s.r[k])
	}

	out := p.Parse()
	if len(p.Errors()) > 0 {
		return nil, fmt.Errorf("parse: %v", p.Errors())
	}
	go out.Run()

	return out, nil
}
