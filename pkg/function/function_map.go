package function

import (
	"strconv"

	"github.com/itsubaki/gocep/pkg/event"
)

type SelectMapAll struct {
	Name string
}

func (f SelectMapAll) Apply(events []event.Event) []event.Event {
	for _, e := range events {
		ref := e.Value(f.Name)
		for _, k := range ref.MapKeys() {
			key := k.Interface().(string)
			val := ref.MapIndex(k)
			e.Record[key] = val.Interface()
		}
	}

	return events
}

type SelectMapString struct {
	Name string
	Key  string
	As   string
}

func (f SelectMapString) Apply(events []event.Event) []event.Event {
	for _, e := range events {
		e.Record[f.As] = e.MapString(f.Name, f.Key)
	}

	return events
}

type SelectMapBool struct {
	Name string
	Key  string
	As   string
}

func (f SelectMapBool) Apply(events []event.Event) []event.Event {
	for _, e := range events {
		e.Record[f.As] = e.MapBool(f.Name, f.Key)
	}

	return events
}

type SelectMapInt struct {
	Name string
	Key  string
	As   string
}

func (f SelectMapInt) Apply(events []event.Event) []event.Event {
	for _, e := range events {
		e.Record[f.As] = e.MapInt(f.Name, f.Key)
	}

	return events
}

type SelectMapFloat struct {
	Name string
	Key  string
	As   string
}

func (f SelectMapFloat) Apply(events []event.Event) []event.Event {
	for _, e := range events {
		e.Record[f.As] = e.MapFloat(f.Name, f.Key)
	}

	return events
}

type SumMapInt struct {
	Name string
	Key  string
	As   string
}

func (f SumMapInt) Apply(events []event.Event) []event.Event {
	sum := 0
	for _, e := range events {
		sum = sum + e.MapInt(f.Name, f.Key)
	}

	for _, e := range events {
		e.Record[f.As] = sum
	}

	return events
}

type SumMapFloat struct {
	Name string
	Key  string
	As   string
}

func (f SumMapFloat) Apply(events []event.Event) []event.Event {
	var sum float64
	for _, e := range events {
		sum = sum + e.MapFloat(f.Name, f.Key)
	}

	for _, e := range events {
		e.Record[f.As] = sum
	}

	return events
}

type AverageMapInt struct {
	Name string
	Key  string
	As   string
}

func (f AverageMapInt) Apply(events []event.Event) []event.Event {
	sum := 0
	for _, e := range events {
		sum = sum + e.MapInt(f.Name, f.Key)
	}
	length := len(events)
	avg := float64(sum) / float64(length)

	for _, e := range events {
		e.Record[f.As] = avg
	}

	return events
}

type AverageMapFloat struct {
	Name string
	Key  string
	As   string
}

func (f AverageMapFloat) Apply(events []event.Event) []event.Event {
	var sum float64
	for _, e := range events {
		sum = sum + e.MapFloat(f.Name, f.Key)
	}

	avg := sum / float64(len(events))

	for _, e := range events {
		e.Record[f.As] = avg
	}

	return events
}

type CastMapStringToInt struct {
	Name string
	Key  string
	As   string
}

func (f CastMapStringToInt) Apply(events []event.Event) []event.Event {
	for _, e := range events {
		str := e.MapString(f.Name, f.Key)
		val, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}

		e.Record[f.As] = val
	}

	return events
}

type CastMapStringToFloat struct {
	Name string
	Key  string
	As   string
}

func (f CastMapStringToFloat) Apply(events []event.Event) []event.Event {
	for _, e := range events {
		str := e.MapString(f.Name, f.Key)
		val, err := strconv.ParseFloat(str, 64)
		if err != nil {
			panic(err)
		}

		e.Record[f.As] = val
	}

	return events
}

type CastMapStringToBool struct {
	Name string
	Key  string
	As   string
}

func (f CastMapStringToBool) Apply(events []event.Event) []event.Event {
	for _, e := range events {
		str := e.MapString(f.Name, f.Key)
		val, err := strconv.ParseBool(str)
		if err != nil {
			panic(err)
		}

		e.Record[f.As] = val
	}

	return events
}
