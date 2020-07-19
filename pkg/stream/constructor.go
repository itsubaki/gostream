package stream

import "github.com/itsubaki/gostream/pkg/clause"

type Where struct {
	w *IdentityWindow
}

type LargerThan struct {
	w *Where
}

func (w *Where) LargerThan() *LargerThan {
	return &LargerThan{w}
}

func (l *LargerThan) Int(name string, value int) {
	l.w.w.SetWhere(
		clause.LargerThanInt{
			Name:  name,
			Value: value,
		},
	)
}

type Function struct {
	w *IdentityWindow
}

type Select struct {
	f *Function
}

func (f *Function) Select() *Select {
	return &Select{f}
}

func (s *Select) String(name, as string) {
	s.f.w.SetFunction(
		clause.SelectString{
			Name: name,
			As:   as,
		},
	)
}

func (s *Select) Int(name, as string) {
	s.f.w.SetFunction(
		clause.SelectInt{
			Name: name,
			As:   as,
		},
	)
}

type Average struct {
	f *Function
}

func (f *Function) Average() *Average {
	return &Average{f}
}

func (a *Average) Int(name, as string) {
	a.f.w.SetFunction(
		clause.AverageInt{
			Name: name,
			As:   as,
		},
	)
}

type Sum struct {
	f *Function
}

func (f *Function) Sum() *Sum {
	return &Sum{f}
}

func (a *Sum) Int(name, as string) {
	a.f.w.SetFunction(
		clause.SumInt{
			Name: name,
			As:   as,
		},
	)
}

type OrderBy struct {
	w *IdentityWindow
}

func (o *OrderBy) Int(name string, reverse bool) {
	o.w.SetOrderBy(
		clause.OrderByInt{
			Name:    name,
			Reverse: reverse,
		},
	)
}
