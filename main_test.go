package main_test

import (
	"fmt"
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/itsubaki/gostream/pkg/builder"
	"github.com/itsubaki/gostream/pkg/parser"
	"github.com/itsubaki/gostream/pkg/window"
)

func TestBuilder(t *testing.T) {
	b := builder.New()
	b.SetField("Name", reflect.TypeOf(""))
	b.SetField("Value", reflect.TypeOf(0))
	s := b.Build()

	i := s.NewInstance()
	i.SetString("Name", "foobar")
	i.SetInt("Value", 123)

	fmt.Printf("%#v\n", i.Value())
	fmt.Printf("%#v\n", i.Pointer())
}

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
	w.OrderBy().Desc().Int("Value")
	w.Limit(10).Offset(5)

	fmt.Printf("%#v\n", w)
}
