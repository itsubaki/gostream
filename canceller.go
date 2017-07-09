package gocep

import "context"

type Canceller struct {
	ctx    context.Context
	cancel func()
}

func NewCanceller() Canceller {
	ctx, cancel := context.WithCancel(context.Background())
	return Canceller{ctx, cancel}
}
