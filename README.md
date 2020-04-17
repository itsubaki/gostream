# gostream

[![Build Status](https://travis-ci.org/itsubaki/gostream.svg?branch=develop)](https://travis-ci.org/itsubaki/gostream)
[![Go Report Card](https://goreportcard.com/badge/github.com/itsubaki/gostream?style=flat-square)](https://goreportcard.com/report/github.com/itsubaki/gostream)

A Stream Processing API for Go

## TODO

 - [x] Window
    + [x] LengthWindow
    + [x] LengthBatchWindow
    + [x] TimeWindow
    + [x] TimeBatchWindow
 - [x] Where
    + [x] EqualsType, NotEqualsType
    + [x] Equals, NotEquals
    + [x] LargerThan, LessThan
 - [ ] GroupBy
 - [x] Function
    + [x] Max, Min, Median
    + [x] Count, Sum, Average
    + [x] Cast
    + [x] As
 - [ ] Having
 - [x] Select
    + [ ] Distinct
 - [x] OrderBy
 - [x] Limit, First, Last
 - [ ] Tool
    + [x] Builder
    + [x] Lexer
    + [ ] Parser

## Install

```console
go get github.com/itsubaki/gostream
```

# Example

```go
type LogEvent struct {
  Time    time.Time
  Level   int
  Message string
}

// select count(*) from LogEvent.time(10sec) where Level > 2
w := window.NewTime(LogEvent{}, 10*time.Second)
defer w.Close()

w.SetWhere(
  expr.LargerThanInt{
    Name: "Level",
    Value: 2,
  },
)
w.SetFunction(
  expr.Count{
    As: "count",
  },
)

go func() {
  for {
    newest := event.Newest(<-w.Output())
    if newest.Int("count") > 10 {
      // notification
    }
  }
}()

w.Input() <- LogEvent{
  Time:    time.Now(),
  Level:   1,
  Message: "this is text log.",
}
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

w := window.NewTime(MyEvent{}, 10 * time.Millisecond)
defer w.Close()

w.SetWhere(
  expr.LargerThanInt{
    Name: "Value",
    Value: 97,
  },
)
w.SetFunction(
  expr.SelectString{
    Name: "Name",
    As: "n",
  },
  expr.SelectInt{
    Name: "Value",
    As: "v",
  },
)
w.SetOrderBy(
  expr.OrderByInt{
    Name: "Value",
    Reverse: true,
  },
)
w.SetLiimt(
  expr.Limit{
    Limit: 10,
    Offset: 5,
  },
)

go func() {
  for {
    fmt.Println(<-w.Output())
  }
}()

for i := 0; i < 100; i++ {
  w.Input() <-MyEvent{
    Name:  "name",
    Value: i,
  }
}
```


```go
// select avg(Value), sum(Value) from MyEvent.length(10)
w := window.NewLength(MyEvent{}, 10)
defer w.Close()

w.SetFunction(
  expr.AverageInt{
    Name: "Value",
    As:   "avg(Value)",
  },
  expr.SumInt{
    Name: "Value",
    As:   "sum(Value)",
  },
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
