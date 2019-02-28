package window

import "context"

type Canceller struct {
	Context context.Context
	Cancel  func()
}

func NewCanceller() Canceller {
	ctx, cancel := context.WithCancel(context.Background())
	return Canceller{ctx, cancel}
}
