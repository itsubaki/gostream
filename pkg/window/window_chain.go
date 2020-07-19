package window

import "github.com/itsubaki/gostream/pkg/clause"

func (w *IdentityWindow) Where() *Where {
	return &Where{w}
}

func (w *IdentityWindow) Function() *Function {
	return &Function{w}
}

func (w *IdentityWindow) OrderBy() *OrderBy {
	return &OrderBy{w}
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
