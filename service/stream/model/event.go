package model

import (
	"container/list"
)

type EventType int

const (
	EVENT_JOIN = iota
	EVENT_LEAVE
	EVENT_MESSAGE
)

type Event struct {
	Type           EventType
	User           string
	Timestamp      int64
	Content        string
}

const archiveSize = 100

var Archive = list.New()

func NewArchive(event Event) {
	if Archive.Len() >= archiveSize {
		Archive.Remove(Archive.Front())
	}
	Archive.PushBack(event)
}

func GetEvents(lastReceived int) []Event {
	events := make([]Event, 0, Archive.Len())
	for event := Archive.Front(); event != nil; event = event.Next() {
		e := event.Value.(Event)
		if e.Timestamp > int64(lastReceived) {
			events = append(events, e)
		}
	}
	return events
}
