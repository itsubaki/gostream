# gostream

[![PkgGoDev](https://pkg.go.dev/badge/github.com/itsubaki/gostream)](https://pkg.go.dev/github.com/itsubaki/gostream)
[![tests](https://github.com/itsubaki/gostream/workflows/tests/badge.svg)](https://github.com/itsubaki/gostream/actions)

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
	panic(err)
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
