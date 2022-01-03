package stream

import (
	"fmt"
	"strings"
)

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
	return nil
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
