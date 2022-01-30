# gostream

[![PkgGoDev](https://pkg.go.dev/badge/github.com/itsubaki/gostream)](https://pkg.go.dev/github.com/itsubaki/gostream)
[![Go Report Card](https://goreportcard.com/badge/github.com/itsubaki/gostream?style=flat-square)](https://goreportcard.com/report/github.com/itsubaki/gostream)
[![tests](https://github.com/itsubaki/gostream/workflows/tests/badge.svg?branch=develop)](https://github.com/itsubaki/gostream/actions)

Stream Processing Library for Go

## TODO

- [x] Window
  - [x] LengthWindow
  - [x] LengthBatchWindow
  - [x] TimeWindow
  - [x] TimeBatchWindow
- [x] Select
- [ ] Where
  - [x] Equals, NotEquals
  - [x] Larger, Less
  - [ ] AND, OR
- [x] OrderBy
- [x] Limit, Offset
- [x] Aggregate Function
  - [x] Avg, Sum, Count
  - [x] Max, Min

## Example

```go
type LogEvent struct {
  Time    time.Time
  Level   int
  Message string
}

q := "select * from LogEvent.length(10)"
s, err := gostream.New().
    Add(LogEvent{}).
    Query(q)
if err != nil {
  fmt.Printf("new gostream: %v", err)
  return
}
defer s.Close()

go func() {
  for {
    fmt.Printf("%v\n", <-s.Output())
  }
}()

s.Input() <- LogEvent{
  Time: time.Now()
  Level: 1
  Message: "something happened"
}
```

```go
type LogEvent struct {
  Time    time.Time
  Level   int
  Message string
}

s := stream.New().
  SelectAll().
  From(LogEvent{}).
  Length(10).
  OrderBy("Level", stream.DESC).
  Limit(10, 5)
defer s.Close()
go s.Run()

go func() {
  for {
    fmt.Printf("%v\n", <-s.Output())
  }
}()

s.Input() <- LogEvent{
  Time: time.Now()
  Level: 1
  Message: "something happened"
}
```
