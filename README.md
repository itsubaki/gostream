# gostream

[![Go Report Card](https://goreportcard.com/badge/github.com/itsubaki/gostream?style=flat-square)](https://goreportcard.com/report/github.com/itsubaki/gostream)
[![tests](https://github.com/itsubaki/gostream/workflows/tests/badge.svg?branch=develop)](https://github.com/itsubaki/gostream/actions)

A Stream Processing API for Go

## TODO

- [x] Window
  - [x] LengthWindow
  - [x] LengthBatchWindow
  - [x] TimeWindow
  - [x] TimeBatchWindow
- [x] Where
  - [x] EqualsType, NotEqualsType
  - [x] Equals, NotEquals
  - [x] LargerThan, LessThan
- [ ] GroupBy
- [x] Function
  - [x] Max, Min, Median
  - [x] Count, Sum, Average
  - [x] Cast
  - [x] As
- [ ] Having
- [x] Select
  - [ ] Distinct
- [x] OrderBy
- [x] Limit, First, Last
- [ ] Tool
  - [x] Lexer
  - [ ] Parser

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
w.Function().Count()
w.Where().LargerThan().Int("Level", 2)
defer w.Close()

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

// select Name, Value
// from MyEvent.time(10msec)
// where Value > 97
// orderby Value DESC
// limit 10 offset 5

w := window.NewTime(MyEvent{}, 10 * time.Millisecond)
w.Function().Select().String("Name")
w.Function().Select().Int("Value")
w.Where().LargerThan().Int("Value", 97)
w.OrderBy().Desc().Int("Value")
w.Limit(10).Offset(5)
defer w.Close()

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
w.Function().Average().Int("Value")
w.Function().Sum().Int("Value")
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
