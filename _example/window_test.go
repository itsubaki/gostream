package _example

import (
	"fmt"
	"testing"
	"time"

	"github.com/itsubaki/gostream/pkg/stream"
)

func TestTimeWindow(t *testing.T) {
	type LogEvent struct {
		Time    time.Time
		Level   int
		Message string
	}

	w := stream.NewTime(LogEvent{}, 10*time.Second)
	defer w.Close()

	w.Where().LargerThan().Int("Level", 2)
	w.Function().Count("count")

	fmt.Printf("%#v\n", w)
}

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
	w.Limit(10).Offset(5)

	fmt.Printf("%#v\n", w)
}
