# gocep

[![Build Status](https://travis-ci.org/itsubaki/gocep.svg?branch=develop)](https://travis-ci.org/itsubaki/gocep)

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
    + [x] Having
    + [ ] GroupBy
 - [x] View
    + [x] OrderBy, Limit
    + [x] First, Last
 - [ ] Utility
    + [x] Builder
    + [x] Lexer
    + [ ] Parser

## Install

```console
go get github.com/itsubaki/gocep
```

# Usecase

```go
type LogEvent struct {
  Time    time.Time
  Level   int
  Message string
}

// select count(*) from LogEvent(10sec) where Level > 2
w := NewTimeWindow(10*time.Second, 1024)
defer w.Close()

w.SetSelector(EqualsType{LogEvent{}})
w.SetSelector(LargerThanInt{"Level", 2})
w.SetFunction(Count{As: "count"})

go func() {
  for {
    events := <-w.Output()
    last := events[len(events)-1]
    if last.Int("count") > 10 {
      // notification
    }
  }
}()

w.Input() <- LogEvent{time.Now(), 1, "this is text log."}
```

# Example

```go
type MyEvent struct {
  Name  string
  Value int
}
```

```go
// select Name as n, Value as v
//  from MyEvent.time(10msec)
//  where Value > 97
//  orderby Value DESC
//  limit 10 offset 5

// 1024 is capacity of input/output queue
w := NewTimeWindow(10 * time.Millisecond, 1024)
defer w.Close()

w.SetSelector(EqualsType{MyEvent{}})
w.SetSelector(LargerThanInt{"Value", 97})
w.SetFunction(SelectString{"Name", "n"})
w.SetFunction(SelectInt{"Value", "v"})
w.SetView(OrderByInt{"Value", true})
w.SetView(Limit{10, 5})

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
w := NewLengthWindow(10, 1024)
defer w.Close()

w.SetSelector(EqualsType{MyEvent{}})
w.SetFunction(AverageInt{"Value", "avg(Value)"})
w.SetFunction(SumInt{"Value", "sum(Value)"})
```

# EventBuilder

```go
// type RuntimeEvent struct {
//  Name string
//  Value int
// }
b := NewStructBuilder()
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
query := "select * from MapEvent.length(10)"
statement, err := NewParser(q).Parse()
if err != nil {
  log.Println("failed.")
  return
}

window := statement.New(1024)
window.Input() <-MapEvent{map}
fmt.Println(<-window.Output())
```
