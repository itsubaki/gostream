package _example

import (
	"fmt"
	"testing"
	"time"

	"github.com/itsubaki/gostream/pkg/window"
)

func TestTimeWindow(t *testing.T) {
	type LogEvent struct {
		Time    time.Time
		Level   int
		Message string
	}

	w := window.NewTime(LogEvent{}, 10*time.Second)
	defer w.Close()

	w.Where().LargerThan().Int("Level", 2)
	w.Function().Count()

	fmt.Printf("%#v\n", w)
}

func TestLengthWindow(t *testing.T) {
	type MyEvent struct {
		Name  string
		Value int
	}

	w := window.NewLength(MyEvent{}, 10)
	defer w.Close()

	w.Function().Average().Int("Value")
	w.Function().Sum().Int("Value")

	fmt.Printf("%#v\n", w)
}

func TestView(t *testing.T) {
	type MyEvent struct {
		Name  string
		Value int
	}

	w := window.NewTime(MyEvent{}, 10*time.Millisecond)
	defer w.Close()

	w.Where().LargerThan().Int("Value", 97)
	w.Function().Select().String("Name")
	w.Function().Select().Int("Value")
	w.OrderBy().Int("Value", true)
	w.Limit(10).Offset(5)

	fmt.Printf("%#v\n", w)
}
