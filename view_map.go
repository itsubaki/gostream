package gocep

import "sort"

type SortableMapInt struct {
	event []Event
	name  string
	key   string
}

func (s SortableMapInt) Len() int {
	return len(s.event)
}

func (s SortableMapInt) Less(i, j int) bool {
	return s.event[i].MapIntValue(s.name, s.key) < s.event[j].MapIntValue(s.name, s.key)
}

func (s SortableMapInt) Swap(i, j int) {
	s.event[i], s.event[j] = s.event[j], s.event[i]
}

type SortMapInt struct {
	Name    string
	Key     string
	Reverse bool
}

func (f SortMapInt) Apply(event []Event) []Event {
	data := SortableMapInt{event, f.Name, f.Key}
	if f.Reverse {
		sort.Sort(sort.Reverse(data))
		return data.event
	}
	sort.Sort(data)
	return data.event
}

type SortableMapFloat struct {
	event []Event
	name  string
	key   string
}

func (s SortableMapFloat) Len() int {
	return len(s.event)
}

func (s SortableMapFloat) Less(i, j int) bool {
	return s.event[i].MapFloatValue(s.name, s.key) < s.event[j].MapFloatValue(s.name, s.key)
}

func (s SortableMapFloat) Swap(i, j int) {
	s.event[i], s.event[j] = s.event[j], s.event[i]
}

type SortMapFloat struct {
	Name    string
	Key     string
	Reverse bool
}

func (f SortMapFloat) Apply(event []Event) []Event {
	data := SortableMapFloat{event, f.Name, f.Key}
	if f.Reverse {
		sort.Sort(sort.Reverse(data))
		return data.event
	}
	sort.Sort(data)
	return data.event
}
