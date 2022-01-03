package stream

import "time"

type Event struct {
	Time       time.Time
	Underlying interface{}
	ResultSet  []interface{}
}

func NewEvent(input interface{}) Event {
	return Event{
		Time:       time.Now(),
		Underlying: input,
		ResultSet:  make([]interface{}, 0),
	}
}

func Newest(e []Event) Event {
	return e[len(e)-1]
}
