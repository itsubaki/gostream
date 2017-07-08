package gocep

import "sort"

type View interface {
	Apply(event []Event) []Event
}

type First struct{}

func (f First) Apply(event []Event) (stream []Event) {
	if len(event) == 0 {
		return stream
	}
	return append(stream, event[0])
}

type Last struct{}

func (f Last) Apply(event []Event) (stream []Event) {
	if len(event) == 0 {
		return stream
	}
	return append(stream, event[len(event)-1])
}

type Limit struct {
	Offset int
	Limit  int
}

func (f Limit) Apply(event []Event) []Event {
	if len(event) < f.Offset+f.Limit {
		return event
	}
	return event[f.Offset : f.Offset+f.Limit]
}

type SortableInt struct {
	event []Event
	name  string
}

func (s SortableInt) Len() int {
	return len(s.event)
}

func (s SortableInt) Less(i, j int) bool {
	return s.event[i].Int(s.name) < s.event[j].Int(s.name)
}

func (s SortableInt) Swap(i, j int) {
	s.event[i], s.event[j] = s.event[j], s.event[i]
}

type SortInt struct {
	Name    string
	Reverse bool
}

func (f SortInt) Apply(event []Event) []Event {
	data := SortableInt{event, f.Name}
	if f.Reverse {
		sort.Sort(sort.Reverse(data))
		return data.event
	}
	sort.Sort(data)
	return data.event
}

type SortableFloat struct {
	event []Event
	name  string
}

func (s SortableFloat) Len() int {
	return len(s.event)
}

func (s SortableFloat) Less(i, j int) bool {
	return s.event[i].Float(s.name) < s.event[j].Float(s.name)
}

func (s SortableFloat) Swap(i, j int) {
	s.event[i], s.event[j] = s.event[j], s.event[i]
}

type SortFloat struct {
	Name    string
	Reverse bool
}

func (f SortFloat) Apply(event []Event) []Event {
	data := SortableFloat{event, f.Name}
	if f.Reverse {
		sort.Sort(sort.Reverse(data))
		return data.event
	}
	sort.Sort(data)
	return data.event
}
