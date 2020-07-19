package _example

import (
	"log"
	"testing"

	"github.com/itsubaki/gostream/pkg/parser"
)

func TestQuery(t *testing.T) {
	type MyEvent struct {
		Name  string
		Value int
	}

	p := parser.New()
	p.Register("MyEvent", MyEvent{})

	query := "select * from MyEvent.length(10)"
	statement, err := p.Parse(query)
	if err != nil {
		log.Println("failed.")
		return
	}

	window := statement.New()
	defer window.Close()
}
