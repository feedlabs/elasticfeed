package model

import (
	"container/list"

	"github.com/gorilla/websocket"
)

type EventType int

type Event struct {
	Type             EventType
	User             string
	Ts               int64
	Timestamp        string
	Content          string
}

type SocketEvent struct {
	Ws                    *websocket.Conn
	Ch                    chan []byte

	ActionId              int

	FeedId                string
	AppId                 string
	OrgId                 string
}

const archiveSize = 10

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
		if e.Ts > int64(lastReceived) {

			events = append(events, e)
		}
	}
	return events
}
