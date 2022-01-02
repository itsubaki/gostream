package stream

import "time"

type Event struct {
	Time       time.Time
	Underlying interface{}
}

func NewEvent(input interface{}) Event {
	return Event{time.Now(), input}
}
