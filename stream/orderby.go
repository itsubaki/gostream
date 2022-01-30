package stream

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

const DESC = true
const ASC = false

type OrderByIF interface {
	Apply(e []Event) []Event
	String() string
}

type NoOrder struct{}

func (o *NoOrder) Apply(e []Event) []Event {
	return e
}

func (o *NoOrder) String() string {
	return ""
}

type OrderBy struct {
	Name  string
	Index int
	Desc  bool
}

func (o *OrderBy) Apply(e []Event) []Event {
	out := append(make([]Event, 0), e...)

	sort.Slice(out, func(i, j int) bool {
		vi := reflect.ValueOf(out[i].Underlying).Field(o.Index).Interface()
		vj := reflect.ValueOf(out[j].Underlying).Field(o.Index).Interface()

		switch v := vi.(type) {
		case int:
			return v < vj.(int)
		case int32:
			return v < vj.(int32)
		case int64:
			return v < vj.(int64)
		case float32:
			return v < vj.(float32)
		case float64:
			return v < vj.(float64)
		case string:
			return v < vj.(string)
		}

		return false
	})

	if o.Desc {
		for i := 0; i < len(out)/2; i++ {
			out[i], out[len(out)-1-i] = out[len(out)-1-i], out[i]
		}
	}

	return out
}

func (o *OrderBy) String() string {
	var buf strings.Builder

	buf.WriteString(fmt.Sprintf("ORDER BY %v", o.Name))
	if o.Desc {
		buf.WriteString(" DESC")
	}

	return buf.String()
}

type LimitIF interface {
	Apply(e []Event) []Event
	String() string
}

type NoLimit struct{}

func (l *NoLimit) Apply(e []Event) []Event {
	return e
}

func (l *NoLimit) String() string {
	return ""
}

type Limit struct {
	Offset int
	Limit  int
}

func (l *Limit) Apply(e []Event) []Event {
	if len(e) < l.Offset+l.Limit {
		return e
	}

	return e[l.Offset : l.Offset+l.Limit]
}

func (l *Limit) String() string {
	var buf strings.Builder

	buf.WriteString(fmt.Sprintf("LIMIT %v", l.Limit))
	if l.Offset > 0 {
		buf.WriteString(fmt.Sprintf(" OFFSET %v", l.Offset))
	}

	return buf.String()
}
