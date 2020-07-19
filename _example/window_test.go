package _example

import (
	"fmt"
	"testing"
	"time"

	"github.com/itsubaki/gostream/pkg/stream"
)

func TestLengthWindow(t *testing.T) {
	type MyEvent struct {
		Name  string
		Value int
	}

	w := stream.NewLength(MyEvent{}, 10)
	defer w.Close()

	w.Function().Average().Int("Value", "avg(Value)")
	w.Function().Sum().Int("Value", "sum(Value)")

	fmt.Printf("%#v\n", w)
}

func TestView(t *testing.T) {
	type MyEvent struct {
		Name  string
		Value int
	}

	w := stream.NewTime(MyEvent{}, 10*time.Millisecond)
	defer w.Close()

	w.Where().LargerThan().Int("Value", 97)
	w.Function().Select().String("Name", "n")
	w.Function().Select().Int("Value", "v")
	w.OrderBy().Int("Value", true)
	w.Limit(10, 5)

	fmt.Printf("%#v\n", w)
}
