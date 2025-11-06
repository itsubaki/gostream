package stream

import "time"

type Event struct {
	Time       time.Time `json:"time"`
	Underlying any       `json:"underlying"`
	ResultSet  []any     `json:"result_set"`
}

func NewEvent(input any) Event {
	return Event{
		Time:       time.Now(),
		Underlying: input,
		ResultSet:  make([]any, 0),
	}
}
