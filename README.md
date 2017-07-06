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
 - [ ] View
    + [x] Sort, Limit
    + [x] First, Last
    + [x] Having
    + [ ] GroupBy
 - [ ] Insert into

## Install

```console
go get github.com/itsubaki/gocep
```

## Example

### Window

```go
type IntEvent struct {
  Name  string
  Value int
}
```

```go
// select avg(Value) from IntEvent.time(10msec) where Value > 97
// 1024 is capacity of input/output queue
w := NewTimeWindow(10 * time.Millisecond, 1024)
defer w.Close()

w.Selector(EqualsType{IntEvent{}})
w.Selector(LargerThanInt{"Value", 97})
w.Function(AverageInt{"Value"})
w.View(SortInt{"Value"})

go func() {
  for {
    event := <-w.Output()
    fmt.Println(event)
  }
}()

for i := 0; i < 100; i++ {
  w.Input() <- NewEvent(IntEvent{"name", i})
}
```

### Stream

 - Stream is Event Dispatcher

```go
s := NewStream(1024)
defer s.Close()

s.Add(NewTimeWindow(10 * time.Millisecond, 1024))
s.Add(NewLengthWindow(10, 1024))
s.Add(...)

s.Push(IntEvent{"name", i})
s.Push(MapEvent{"name", m})
s.Push(...)
```
