package event

type Event struct {
	data interface{}

	eventGroup    string
	eventName     string

	parent        string
	target        string
}

func NewEvent(data interface{}) *Event {
	return &Event{data, "", "", "", ""}
}
