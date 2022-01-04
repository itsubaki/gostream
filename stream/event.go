package stream

import "time"

type Event struct {
	Time       time.Time     `json:"time"`
	Underlying interface{}   `json:"underlying"`
	ResultSet  []interface{} `json:"result_set"`
}

func NewEvent(input interface{}) Event {
	return Event{
		Time:       time.Now(),
		Underlying: input,
		ResultSet:  make([]interface{}, 0),
	}
}
