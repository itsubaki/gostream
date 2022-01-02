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
		opt: &Option{
			Verbose: false,
		},
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
		var buf strings.Builder
		l := lexer.New(strings.NewReader(q))
		for {
			tok, lit := l.Tokenize()
			if tok == lexer.EOF {
				break
			}

			buf.WriteString(fmt.Sprintf("%v", lexer.Tokens[tok]))
			if lexer.IsBasicLit(tok) {
				buf.WriteString(fmt.Sprintf("(%v)", lit))
			}
			buf.WriteString(" ")
		}

		fmt.Println(strings.TrimRight(buf.String(), " "))
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
