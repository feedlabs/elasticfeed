package room

import (
	"strconv"
	"time"
	"container/list"
	"encoding/json"

	"github.com/feedlabs/feedify"
	"github.com/gorilla/websocket"

	"github.com/astaxie/beego/session"

	"github.com/feedlabs/elasticfeed/service/stream/model"
)

const (
	SYSTEM_FEED_MESSAGE = 1
	FEED_ENTRY_MESSAGE  = 2
	ENTRY_MESSAGE = 8
)

var (
	Subscribe   = make(chan Subscriber, 10)
	Unsubscribe = make(chan string, 10)
	Publish     = make(chan model.Event, 10)
	P2P         = make(chan *websocket.Conn, 10)

	WaitingList = list.New()
	Subscribers = list.New()

	GlobalSessions *session.Manager
)

type Subscription struct {
	Archive []model.Event
	New     <-chan model.Event
}

type Subscriber struct {
	Name    string
	Conn    *websocket.Conn
}

func NewEvent(ep model.EventType, user, msg string) model.Event {
	ts := time.Now().UnixNano()
	return model.Event{ep, user, ts, strconv.Itoa(int(ts)), msg}
}

func NewChannelEvent(ep model.EventType, user, msg string) model.Event {
	return NewEvent(ep, user, msg)
}

func NewSystemEvent(ep model.EventType, user, msg string) model.Event {
	event := NewEvent(ep, user, msg)
	data, _ := json.Marshal(event)

	return NewChannelEvent(model.EVENT_MESSAGE, user, string(data))
}

func NewFeedEvent(ep model.EventType, user, msg string) model.Event {
	// "msg" is a feed action; can contain entry specific event
	event := NewEvent(ep, user, msg)
	data, _ := json.Marshal(event)

	// "user" is and feed-id; "*" means all feeds on the client site
	return NewSystemEvent(SYSTEM_FEED_MESSAGE, user, string(data))
}

func NewEntryEvent(ep model.EventType, user, msg string) model.Event {
	// "msg" is a feed entry data as a string
	event := NewEvent(ep, user, msg)
	data, _ := json.Marshal(event)

	// "*" all feeds on client site will receive this message
	return NewFeedEvent(ENTRY_MESSAGE, "*", string(data))
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
		Publish <- NewChannelEvent(model.EVENT_JOIN, sub.Name, "")

		case client := <-P2P:
			data, _ := json.Marshal(NewSystemEvent(model.EVENT_MESSAGE, "system", "ok"))
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
					Publish <- NewChannelEvent(model.EVENT_LEAVE, unsub, "")
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
	GlobalSessions, _ = session.NewManager("memory", `{"cookieName":"elasticfeedsessid","gclifetime":3600}`)

	go FeedManager()
	go GlobalSessions.GC()
}
