package clause

import (
	"sort"

	"github.com/itsubaki/gostream/pkg/event"
)

type OrderBy interface {
	Apply(events []event.Event) []event.Event
}

type SortableInt struct {
	events []event.Event
	name   string
}

func (s SortableInt) Len() int {
	return len(s.events)
}

func (s SortableInt) Less(i, j int) bool {
	return s.events[i].Int(s.name) < s.events[j].Int(s.name)
}

func (s SortableInt) Swap(i, j int) {
	s.events[i], s.events[j] = s.events[j], s.events[i]
}

type OrderByInt struct {
	Name    string
	Reverse bool
}

func (f OrderByInt) Apply(events []event.Event) []event.Event {
	data := SortableInt{events, f.Name}
	if f.Reverse {
		sort.Sort(sort.Reverse(data))
		return data.events
	}
	sort.Sort(data)
	return data.events
}

type SortableFloat struct {
	events []event.Event
	name   string
}

func (s SortableFloat) Len() int {
	return len(s.events)
}

func (s SortableFloat) Less(i, j int) bool {
	return s.events[i].Float(s.name) < s.events[j].Float(s.name)
}

func (s SortableFloat) Swap(i, j int) {
	s.events[i], s.events[j] = s.events[j], s.events[i]
}

type OrderByFloat struct {
	Name    string
	Reverse bool
}

func (f OrderByFloat) Apply(events []event.Event) []event.Event {
	data := SortableFloat{events, f.Name}
	if f.Reverse {
		sort.Sort(sort.Reverse(data))
		return data.events
	}
	sort.Sort(data)
	return data.events
}
