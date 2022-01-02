package stream_test

import (
	"testing"
	"time"

	"github.com/itsubaki/gostream/lexer"
	"github.com/itsubaki/gostream/stream"
)

func TestWindowString(t *testing.T) {
	cases := []struct {
		in   stream.Window
		want string
	}{
		{&stream.Length{Length: 10}, "LENGTH(10)"},
		{&stream.LengthBatch{Length: 10}, "LENGTH_BATCH(10)"},
		{&stream.Time{Expire: 10 * time.Minute, Unit: lexer.MIN}, "TIME(10 MIN)"},
		{&stream.TimeBatch{Expire: 10 * time.Minute, Unit: lexer.MIN}, "TIME_BATCH(10 MIN)"},
	}

	for _, c := range cases {
		if c.in.String() != c.want {
			t.Errorf("got=%v, want=%v", c.in.String(), c.want)
		}
	}
}
