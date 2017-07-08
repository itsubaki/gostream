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
 - [ ] View
    + [x] Sort, Limit
    + [x] First, Last
    + [ ] Having
    + [ ] GroupBy
 - [ ] Insert into

## Install

```console
go get github.com/itsubaki/gocep
```

## Example

### Window

```go
type MyEvent struct {
  Name  string
  Value int
}
```

```go
// select * from MyEvent.time(10msec).sort(Value) where Value > 97
// 1024 is capacity of input/output queue
w := NewTimeWindow(10 * time.Millisecond, 1024)
defer w.Close()

w.Selector(EqualsType{MyEvent{}})
w.Selector(LargerThanInt{"Value", 97})
w.View(SortInt{"Value"})

go func() {
  for {
    event := <-w.Output()
    fmt.Println(event)
  }
}()

for i := 0; i < 100; i++ {
  w.Input() <- NewEvent(MyEvent{"name", i})
}
```


```go
// select avg(Value) as avgval, sum(Value) as sumval from MyEvent.length(10)
w := NewLengthWindow(10, 1024)
defer w.Close()

w.Selector(EqualsType{MyEvent{}})
w.Function(AverageInt{"Value", "avgval"})
w.Function(SumInt{"Value", "sumval"})
```

### Stream

 - Stream is Event Dispatcher

```go
s := NewStream(1024)
defer s.Close()

s.Add(NewTimeWindow(10 * time.Millisecond, 1024))
s.Add(NewLengthWindow(10, 1024))
s.Add(...)

go func() {
  for {
    event := <-s.Output()
    fmt.Println(event)
  }
}()

s.Input() <- MyEvent{"name", 100}
s.Input() <- MapEvent{"name", map}
...
```
