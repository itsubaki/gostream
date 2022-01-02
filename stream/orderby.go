package stream

type OrderBy struct {
	Name string
	Desc bool
}

func (o *OrderBy) Apply(ev []Event) []Event {

	return nil
}

type Limit struct {
	Offset int
	Limit  int
}

func (l *Limit) Apply(ev []Event) []Event {
	if len(ev) < l.Offset+l.Limit {
		return ev
	}

	return ev[l.Offset : l.Offset+l.Limit]
}
