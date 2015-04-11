package event

type Event struct {
	data interface{}

	eventGroup    string
	eventName     string
}

func NewEvent(data interface{}) *Event {
	return &Event{data, "", ""}
}
