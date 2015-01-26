package room

import (
	"time"
	"container/list"
	"encoding/json"

	"github.com/feedlabs/feedify"
	"github.com/gorilla/websocket"

	"github.com/feedlabs/elasticfeed/service/stream/model"
)

var (
	Subscribe   = make(chan Subscriber, 10)
	Unsubscribe = make(chan string, 10)
	Publish     = make(chan model.Event, 10)
	P2P         = make(chan *websocket.Conn, 10)

	WaitingList = list.New()
	Subscribers = list.New()
)

type Subscription struct {
	Archive []model.Event
	New     <-chan model.Event
}

type Subscriber struct {
	Name string
	Conn *websocket.Conn
}

func NewEvent(ep model.EventType, user, msg string) model.Event {
	return model.Event{ep, user, time.Now().UnixNano(), msg}
}

func Join(user string, ws *websocket.Conn) {
	Subscribe <- Subscriber{Name: user, Conn: ws}
}

func Leave(user string) {
	Unsubscribe <- user
}

func FeedManager() {
	for {
		select {

		case sub := <-Subscribe:
			Subscribers.PushBack(sub)
			Publish <- NewEvent(model.EVENT_JOIN, sub.Name, "")

		case client := <-P2P:
			data, _ := json.Marshal(&model.Event{model.EVENT_MESSAGE, "system", time.Now().UnixNano(), "ok"})
			client.WriteMessage(websocket.TextMessage, data)

		case event := <-Publish:
			model.NewArchive(event)

			for ch := WaitingList.Back(); ch != nil; ch = ch.Prev() {
				ch.Value.(chan bool) <- true
				WaitingList.Remove(ch)
			}

			broadcastWebSocket(event)

		case unsub := <-Unsubscribe:
			for sub := Subscribers.Front(); sub != nil; sub = sub.Next() {
				if sub.Value.(Subscriber).Name == unsub {
					Subscribers.Remove(sub)

					ws := sub.Value.(Subscriber).Conn
					if ws != nil {
						ws.Close()
						feedify.Error("WebSocket closed:", unsub)
					}
					Publish <- NewEvent(model.EVENT_LEAVE, unsub, "")
					break
				}
			}
		}
	}
}

func broadcastWebSocket(event model.Event) {
	data, err := json.Marshal(event)
	if err != nil {
		feedify.Error("Fail to marshal event:", err)
		return
	}

	for sub := Subscribers.Front(); sub != nil; sub = sub.Next() {
		ws := sub.Value.(Subscriber).Conn
		if ws != nil {
			if ws.WriteMessage(websocket.TextMessage, data) != nil {
				Unsubscribe <- sub.Value.(Subscriber).Name
			}
		}
	}
}

func init() {
	go FeedManager()
}
