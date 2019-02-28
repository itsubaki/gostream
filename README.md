# gocep

[![Build Status](https://travis-ci.org/itsubaki/gocep.svg?branch=develop)](https://travis-ci.org/itsubaki/gocep)
[![Go Report Card](https://goreportcard.com/badge/github.com/itsubaki/gocep?style=flat-square)](https://goreportcard.com/report/github.com/itsubaki/gocep)


The Stream Processing API for Go

## TODO

 - [x] Window
    + [x] LengthWindow
    + [x] LengthBatchWindow
    + [x] TimeWindow
    + [x] TimeBatchWindow
 - [x] Selector
    + [x] EqualsType, NotEqualsType
    + [x] Equals, NotEquals
    + [x] LargerThan, LessThan
 - [x] Function
    + [x] Max, Min, Median
    + [x] Count, Sum, Average
    + [x] Cast
    + [x] As
    + [x] SelectAll, Select
    + [ ] GroupBy
    + [ ] Having
 - [x] View
    + [x] OrderBy, Limit
    + [x] First, Last
 - [ ] Tool
    + [x] Builder
    + [x] Lexer
    + [ ] Parser

## Install

```console
go get github.com/itsubaki/gocep
```

# Example

```go
type LogEvent struct {
  Time    time.Time
  Level   int
  Message string
}

// select count(*) from LogEvent.time(10sec) where Level > 2
w := window.NewTime(10*time.Second)
defer w.Close()

w.SetSelector(
  selector.EqualsType{Accept: LogEvent{}},
  selector.LargerThanInt{Name: "Level", Value: 2},
)
w.SetFunction(function.Count{As: "count"})

go func() {
  for {
    newest := event.Newest(<-w.Output())
    if newest.Int("count") > 10 {
      // notification
    }
  }
}()

w.Input() <- LogEvent{time.Now(), 1, "this is text log."}
```

```go
type MyEvent struct {
  Name  string
  Value int
}

// select Name as n, Value as v
//  from MyEvent.time(10msec)
//  where Value > 97
//  orderby Value DESC
//  limit 10 offset 5

w := window.NewTime(10 * time.Millisecond)
defer w.Close()

w.SetSelector(
  selector.EqualsType{Accept: MyEvent{}},
  selector.LargerThanInt{Name: "Value", Value: 97},
)
w.SetFunction(
  function.SelectString{Name: "Name", As: "n"},
  function.SelectInt{Name: "Value", As: "v"},
)
w.SetView(
  view.OrderByInt{Name: "Value", Reverse: true},
  view.Limit{Limit: 10, Offset: 5},
)

go func() {
  for {
    fmt.Println(<-w.Output())
  }
}()

for i := 0; i < 100; i++ {
  w.Input() <-MyEvent{"name", i}
}
```


```go
// select avg(Value), sum(Value) from MyEvent.length(10)
w := window.NewLength(10)
defer w.Close()

w.SetSelector(
  selector.EqualsType{Accept: MyEvent{}},
)
w.SetFunction(
  function.AverageInt{Name: "Value", As: "avg(Value)"},
  function.SumInt{Name: "Value", As: "sum(Value)"},
)
```

# RuntimeEventBuilder

```go
// type RuntimeEvent struct {
//  Name string
//  Value int
// }
b := builder.New()
b.SetField("Name", reflect.TypeOf(""))
b.SetField("Value", reflect.TypeOf(0))
s := b.Build()


// i.Value()
// -> RuntimeEvent{Name: "foobar", Value: 123}
// i.Pointer()
// -> &RuntimeEvent{Name: "foobar", Value: 123}
i := s.NewInstance()
i.SetString("Name", "foobar")
i.SetInt("Value", 123)

w.Input() <-i.Value()
```

# (WIP) Query

```go
p := parser.New()
p.Register("MapEvent", MapEvent{})

query := "select * from MapEvent.length(10)"
statement, err := p.Parse(query)
if err != nil {
  log.Println("failed.")
  return
}

window := statement.New()
defer window.Close()

window.Input() <-MapEvent{map}
fmt.Println(<-window.Output())
```
