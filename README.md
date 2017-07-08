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
    + [ ] Having
    + [ ] GroupBy
 - [x] Function
    + [x] Max, Min, Median
    + [x] Count, Sum, Average
    + [x] Cast
    + [x] As
 - [x] View
    + [x] OrderBy, Limit
    + [x] First, Last
 - [x] Insert

## Install

```console
go get github.com/itsubaki/gocep
```

# Example

## Window

```go
type MyEvent struct {
  Name  string
  Value int
}
```

```go
// select * from MyEvent.time(10msec)
//  where Value > 97
//  orderby Value DESC
//  limit 10 offset 5

// 1024 is capacity of input/output queue
w := NewTimeWindow(10 * time.Millisecond, 1024)
defer w.Close()

w.Selector(EqualsType{MyEvent{}})
w.Selector(LargerThanInt{"Value", 97})
w.View(OrderByInt{"Value", true})
w.View(Limit{5, 10})

go func() {
  for {
    event := <-w.Output()
    fmt.Println(event)
  }
}()

for i := 0; i < 100; i++ {
  w.Input() <-MyEvent{"name", i}
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

## Stream

 - Stream is Event Dispatcher and Collector

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

s.Input() <-MyEvent{"name", 100}
s.Input() <-MapEvent{"name", map}
...
```

```go
// select sum(Value) from MyEvent.length(10)
s := NewStream(1024)
defer s.Close()
w := NewLengthWindow(10, 1024)
w.Selector(EqualsType{MyEvent{}})
w.Function(SumInt{"Value", "sum(Value)"})
s.Add(w)

// select * from RecordEvent.length(10) where sum(Value) > 10
is := NewStream(1024)
iw := NewLengthWindow(10, 1024)
iw.Selector(EqualsType{RecordEvent{}})
iw.Selector(LargerThanMapInt{"Record", "sum(Value)", 10})
is.Add(iw)

// insert into RecordEvent select sum(Value) from MyEvent.length(10)
// select * from RecordEvent.length(10) where sum(Value) > 10
s.Insert(is)

s.Input() <-MyEvent{"name", 100}
fmt.Println(<-is.Output())
```
