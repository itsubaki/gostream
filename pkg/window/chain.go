package window

import (
	"fmt"

	"github.com/itsubaki/gostream/pkg/clause"
)

type Chain interface {
	Where() *Where
	Function() *Function
	OrderBy() *OrderBy
	Limit(limit int) *Limit
}

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
	l.w.w.SetWhere(clause.LargerThanInt{Name: name, Value: value})
}

type Function struct {
	w *IdentityWindow
}

func (f *Function) Count() {
	f.w.SetFunction(clause.Count{As: "count(*)"})
}

type Select struct {
	f *Function
}

func (f *Function) Select() *Select {
	return &Select{f}
}

func (s *Select) String(name string) {
	s.f.w.SetFunction(clause.SelectString{Name: name, As: name})
}

func (s *Select) Int(name string) {
	s.f.w.SetFunction(clause.SelectInt{Name: name, As: name})
}

type Average struct {
	f *Function
}

func (f *Function) Average() *Average {
	return &Average{f}
}

func (a *Average) Int(name string) {
	as := fmt.Sprintf("avg(%s)", name)
	a.f.w.SetFunction(clause.AverageInt{Name: name, As: as})
}

type Sum struct {
	f *Function
}

func (f *Function) Sum() *Sum {
	return &Sum{f}
}

func (a *Sum) Int(name string) {
	as := fmt.Sprintf("sum(%s)", name)
	a.f.w.SetFunction(clause.SumInt{Name: name, As: as})
}

type OrderBy struct {
	w    *IdentityWindow
	desc bool
}

func (o *OrderBy) Desc() *OrderBy {
	o.desc = true
	return o
}

func (o *OrderBy) Int(name string) {
	o.w.SetOrderBy(clause.OrderByInt{Name: name, Desc: o.desc})
}

type Limit struct {
	w     *IdentityWindow
	limit int
}

func (l *Limit) Offset(offset int) {
	l.w.SetLimit(clause.Limit{Limit: l.limit, Offset: offset})
}
