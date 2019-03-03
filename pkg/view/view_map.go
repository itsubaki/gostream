package view

import (
	"sort"

	"github.com/itsubaki/gostream-core/pkg/event"
)

type SortableMapInt struct {
	event []event.Event
	name  string
	key   string
}

func (s SortableMapInt) Len() int {
	return len(s.event)
}

func (s SortableMapInt) Less(i, j int) bool {
	return s.event[i].MapInt(s.name, s.key) < s.event[j].MapInt(s.name, s.key)
}

func (s SortableMapInt) Swap(i, j int) {
	s.event[i], s.event[j] = s.event[j], s.event[i]
}

type OrderByMapInt struct {
	Name    string
	Key     string
	Reverse bool
}

func (f OrderByMapInt) Apply(events []event.Event) []event.Event {
	data := SortableMapInt{events, f.Name, f.Key}
	if f.Reverse {
		sort.Sort(sort.Reverse(data))
		return data.event
	}
	sort.Sort(data)
	return data.event
}

type SortableMapFloat struct {
	event []event.Event
	name  string
	key   string
}

func (s SortableMapFloat) Len() int {
	return len(s.event)
}

func (s SortableMapFloat) Less(i, j int) bool {
	return s.event[i].MapFloat(s.name, s.key) < s.event[j].MapFloat(s.name, s.key)
}

func (s SortableMapFloat) Swap(i, j int) {
	s.event[i], s.event[j] = s.event[j], s.event[i]
}

type OrderByMapFloat struct {
	Name    string
	Key     string
	Reverse bool
}

func (f OrderByMapFloat) Apply(events []event.Event) []event.Event {
	data := SortableMapFloat{events, f.Name, f.Key}
	if f.Reverse {
		sort.Sort(sort.Reverse(data))
		return data.event
	}
	sort.Sort(data)
	return data.event
}
