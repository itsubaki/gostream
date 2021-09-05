package function

import "github.com/itsubaki/gostream/pkg/event"

type LimitIF interface {
	Apply(events []event.Event) []event.Event
}

type First struct{}

func (f First) Apply(events []event.Event) []event.Event {
	out := event.List()
	if len(events) == 0 {
		return out
	}

	return append(out, events[0])
}

type Last struct{}

func (f Last) Apply(events []event.Event) []event.Event {
	out := event.List()
	if len(events) == 0 {
		return out
	}

	return append(out, events[len(events)-1])
}

type Limit struct {
	Limit  int
	Offset int
}

func (f Limit) Apply(events []event.Event) []event.Event {
	if len(events) < f.Offset+f.Limit {
		return events
	}

	return events[f.Offset : f.Offset+f.Limit]
}
