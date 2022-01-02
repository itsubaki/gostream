# gostream

[![Go Report Card](https://goreportcard.com/badge/github.com/itsubaki/gostream?style=flat-square)](https://goreportcard.com/report/github.com/itsubaki/gostream)
[![tests](https://github.com/itsubaki/gostream/workflows/tests/badge.svg?branch=develop)](https://github.com/itsubaki/gostream/actions)

Stream Processing Library for Go

## TODO

- [x] Window
  - [x] LengthWindow
  - [x] LengthBatchWindow
  - [x] TimeWindow
  - [x] TimeBatchWindow
- [ ] Select
  - [ ] First, Last
  - [ ] Distinct
- [x] Where
  - [x] Equals, NotEquals
  - [x] Larger, Less
- [ ] OrderBy, Limit
- [ ] Aggregate Function
  - [ ] Max, Min, Median
  - [ ] Count, Sum, Average
  - [ ] GroupBy

## Example

```go
type LogEvent struct {
  Time    time.Time
  Level   int
  Message string
}

q := "select * from LogEvent.length(10)"
s, err := gostream.New().Add(LogEvent{}).Query(q)
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
