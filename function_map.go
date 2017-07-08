package gocep

import "strconv"

type SelectMapAll struct {
}

func (f SelectMapAll) Apply(event []Event) []Event {
	return event
}

type SelectMapString struct {
	Name string
	Key  string
	As   string
}

func (f SelectMapString) Apply(event []Event) []Event {
	for _, e := range event {
		e.Record[f.As] = e.MapString(f.Name, f.Key)
	}
	return event
}

type SelectMapBool struct {
	Name string
	Key  string
	As   string
}

func (f SelectMapBool) Apply(event []Event) []Event {
	for _, e := range event {
		e.Record[f.As] = e.MapBool(f.Name, f.Key)
	}
	return event
}

type SelectMapInt struct {
	Name string
	Key  string
	As   string
}

func (f SelectMapInt) Apply(event []Event) []Event {
	for _, e := range event {
		e.Record[f.As] = e.MapInt(f.Name, f.Key)
	}
	return event
}

type SelectMapFloat struct {
	Name string
	Key  string
	As   string
}

func (f SelectMapFloat) Apply(event []Event) []Event {
	for _, e := range event {
		e.Record[f.As] = e.MapFloat(f.Name, f.Key)
	}
	return event
}

type SumMapInt struct {
	Name string
	Key  string
	As   string
}

func (f SumMapInt) Apply(event []Event) []Event {
	sum := 0
	for _, e := range event {
		sum = sum + e.MapInt(f.Name, f.Key)
	}

	for _, e := range event {
		e.Record[f.As] = sum
	}

	return event
}

type SumMapFloat struct {
	Name string
	Key  string
	As   string
}

func (f SumMapFloat) Apply(event []Event) []Event {
	var sum float64
	for _, e := range event {
		sum = sum + e.MapFloat(f.Name, f.Key)
	}

	for _, e := range event {
		e.Record[f.As] = sum
	}

	return event
}

type AverageMapInt struct {
	Name string
	Key  string
	As   string
}

func (f AverageMapInt) Apply(event []Event) []Event {
	sum := 0
	for _, e := range event {
		sum = sum + e.MapInt(f.Name, f.Key)
	}
	length := len(event)
	avg := float64(sum) / float64(length)

	for _, e := range event {
		e.Record[f.As] = avg
	}

	return event
}

type AverageMapFloat struct {
	Name string
	Key  string
	As   string
}

func (f AverageMapFloat) Apply(event []Event) []Event {
	var sum float64
	for _, e := range event {
		sum = sum + e.MapFloat(f.Name, f.Key)
	}
	length := len(event)
	avg := float64(sum) / float64(length)

	for _, e := range event {
		e.Record[f.As] = avg
	}

	return event
}

type CastMapStringToInt struct {
	Name string
	Key  string
	As   string
}

func (f CastMapStringToInt) Apply(event []Event) []Event {
	for _, e := range event {
		str := e.MapString(f.Name, f.Key)
		val, _ := strconv.Atoi(str)
		e.Record[f.As] = val
	}

	return event
}

type CastMapStringToFloat struct {
	Name string
	Key  string
	As   string
}

func (f CastMapStringToFloat) Apply(event []Event) []Event {
	for _, e := range event {
		str := e.MapString(f.Name, f.Key)
		val, _ := strconv.ParseFloat(str, 64)
		e.Record[f.As] = val
	}

	return event
}

type CastMapStringToBool struct {
	Name string
	Key  string
	As   string
}

func (f CastMapStringToBool) Apply(event []Event) []Event {
	for _, e := range event {
		str := e.MapString(f.Name, f.Key)
		val, _ := strconv.ParseBool(str)
		e.Record[f.As] = val
	}

	return event
}
