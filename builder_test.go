package gocep

import "testing"

func TestBuilder(t *testing.T) {
	w := NewBuilder().
		Window("LengthWindow(16, 64)").
		Selector("IntEvent").
		Selector("LargerThanInt{Value,97}").
		Function("Average{Value}").
		View("Sort{Value}").
		Build()
	defer w.Close()
}
