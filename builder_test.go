package gocep

import "testing"

func TestBuilder(t *testing.T) {
	b := NewBuilder()
	w := b.Build()
	defer w.Close()
}
