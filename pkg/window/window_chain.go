package window

import "github.com/itsubaki/gostream/pkg/clause"

func (w *IdentityWindow) Select() *Select {
	return w.Function().Select()
}

func (w *IdentityWindow) Sum() *Sum {
	return w.Function().Sum()
}

func (w *IdentityWindow) Average() *Average {
	return w.Function().Average()
}

func (w *IdentityWindow) Count() {
	w.Function().Count()
}

func (w *IdentityWindow) Where() *Where {
	return &Where{w}
}

func (w *IdentityWindow) Function() *Function {
	return &Function{w}
}

func (w *IdentityWindow) OrderBy() *OrderBy {
	return &OrderBy{w, false}
}

func (w *IdentityWindow) Limit(limit int) *Limit {
	w.SetLimit(clause.Limit{Limit: limit, Offset: 0})
	return &Limit{w, limit}
}

func (w *IdentityWindow) First() {
	w.SetLimit(clause.First{})
}

func (w *IdentityWindow) Last() {
	w.SetLimit(clause.Last{})
}
