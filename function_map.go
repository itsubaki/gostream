package gocep

type SumMapInt struct {
	Name string
	Key  string
}

func (f SumMapInt) Apply(event []Event) []Event {
	sum := 0
	for _, e := range event {
		sum = sum + e.MapIntValue(f.Name, f.Key)
	}

	for _, e := range event {
		e.Record["sum("+f.Name+":"+f.Key+")"] = sum
	}

	return event
}

type SumMapFloat struct {
	Name string
	Key  string
}

func (f SumMapFloat) Apply(event []Event) []Event {
	var sum float64
	for _, e := range event {
		sum = sum + e.MapFloatValue(f.Name, f.Key)
	}

	for _, e := range event {
		e.Record["sum("+f.Name+":"+f.Key+")"] = sum
	}

	return event
}

type AverageMapInt struct {
	Name string
	Key  string
}

func (f AverageMapInt) Apply(event []Event) []Event {
	sum := 0
	for _, e := range event {
		sum = sum + e.MapIntValue(f.Name, f.Key)
	}
	length := len(event)
	avg := float64(sum) / float64(length)

	for _, e := range event {
		e.Record["avg("+f.Name+":"+f.Key+")"] = avg
	}

	return event
}

type AverageMapFloat struct {
	Name string
	Key  string
}

func (f AverageMapFloat) Apply(event []Event) []Event {
	var sum float64
	for _, e := range event {
		sum = sum + e.MapFloatValue(f.Name, f.Key)
	}
	length := len(event)
	avg := float64(sum) / float64(length)

	for _, e := range event {
		e.Record["avg("+f.Name+":"+f.Key+")"] = avg
	}

	return event
}
