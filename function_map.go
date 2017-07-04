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
