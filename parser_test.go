package gocep

import (
	"fmt"
	"testing"
)

func TestParser(t *testing.T) {
	query := "select sum(Value) from MapEvent.time(10msec) where Value > 100"
	b := &Parser{}
	w := b.Parse(query)
	fmt.Println(w)
}
