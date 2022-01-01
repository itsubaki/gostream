package gostream

import "github.com/itsubaki/gostream/lexer"

func Parse(l *lexer.Lexer) Stream {
	return NewIdentityStream()
}
