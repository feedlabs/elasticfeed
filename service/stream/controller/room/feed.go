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
	CHANNEL_JOIN    = 0
	CHANNEL_LEAVE   = 1
	CHANNEL_MESSAGE = 2

	SYSTEM_FEED_MESSAGE = 1

	FEED_RELOAD        = 1
	FEED_EMPTY         = 2
	FEED_ENTRY_NEW     = 3
	FEED_ENTRY_INIT    = 4
	FEED_ENTRY_MORE    = 5
	FEED_HIDE          = 6
	FEED_SHOW          = 7
	FEED_ENTRY_MESSAGE = 8

	ENTRY_UPDATE = 1
	ENTRY_DELETE = 2
	ENTRY_SHOW   = 3
	ENTRY_HIDE   = 4
)

var (
	FeedRoom *FeedRoomManager
	GlobalSessions *session.Manager
)

type Subscription struct {
	Archive []model.Event
	New     <-chan model.Event
}

type Subscriber struct {
	Name      string
	Conn      *websocket.Conn
}

func NewEvent(ep model.EventType, user, msg string) model.Event {
	ts := time.Now().UnixNano()
	return model.Event{ep, user, ts, strconv.Itoa(int(ts)), msg}
}

func NewSocketEvent(msg []byte, ws *websocket.Conn, ch chan []byte) model.SocketEvent {
	data := make(map[string]interface{})

	json.Unmarshal(msg, &data)

	return model.SocketEvent{ws, ch, 4, data["feedId"].(string), data["appId"].(string), data["orgId"].(string)}
}

func NewChannelEvent(ep model.EventType, user, msg string) model.Event {
	return NewEvent(ep, user, msg)
}

func NewSystemEvent(ep model.EventType, user, msg string) model.Event {
	event := NewEvent(ep, user, msg)
	data, _ := json.Marshal(event)

	return NewChannelEvent(CHANNEL_MESSAGE, user, string(data))
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
	return NewFeedEvent(FEED_ENTRY_MESSAGE, "*", string(data))
}

type FeedRoomManager struct {
	Subscribe     chan Subscriber
	Unsubscribe   chan string
	Publish       chan model.Event

	ResourceEvent chan model.SocketEvent

	WaitingList        *list.List
	Subscribers        *list.List
}

func (this *FeedRoomManager) Join(user string, ws *websocket.Conn) {
	this.Subscribe <- Subscriber{Name: user, Conn: ws}
}

func (this *FeedRoomManager) Leave(user string) {
	this.Unsubscribe <- user
}

func (this *FeedRoomManager) Run() {
	go func() {
		for {
			select {

			case sub := <-this.Subscribe:
				this.Subscribers.PushBack(sub)

			case event := <-this.Publish:

				// here must be handled where to send notification
				// - or to all sockets
				// - or to specific client/feed (single socket)
				// - or to public feed (multiple sockets)
				//
				// could be setup by resource manager go routine
				// room.FeedSubscribers[socketEvent.FeedId][channelID] = socketEvent

				model.NewArchive(event)

				for ch := this.WaitingList.Back(); ch != nil; ch = ch.Prev() {
					ch.Value.(chan bool) <- true
					this.WaitingList.Remove(ch)
				}

				this.BroadcastWebSocket(event)

			case unsub := <-this.Unsubscribe:
				for sub := this.Subscribers.Front(); sub != nil; sub = sub.Next() {
					if sub.Value.(Subscriber).Name == unsub {
						this.Subscribers.Remove(sub)

						ws := sub.Value.(Subscriber).Conn
						if ws != nil {
							ws.Close()
							feedify.Error("WebSocket closed:", unsub)
						}
						this.Publish <- NewChannelEvent(CHANNEL_LEAVE, unsub, "")
						break
					}
				}
			}
		}
	}()
}

func (this *FeedRoomManager) BroadcastWebSocket(event model.Event) {
	data, err := json.Marshal(event)
	if err != nil {
		feedify.Error("Fail to marshal event:", err)
		return
	}

	for sub := this.Subscribers.Front(); sub != nil; sub = sub.Next() {
		ws := sub.Value.(Subscriber).Conn
		if ws != nil {
			if ws.WriteMessage(websocket.TextMessage, data) != nil {
				this.Unsubscribe <- sub.Value.(Subscriber).Name
			}
		}
	}
}

func NewFeedRoomManager() *FeedRoomManager {
	subscribe := make(chan Subscriber, 10)
	unsubscribe := make(chan string, 10)
	publish := make(chan model.Event, 10)
	resourceEvent := make(chan model.SocketEvent, 10)

	waitingList := list.New()
	subscribers := list.New()

	FeedRoom = &FeedRoomManager{subscribe, unsubscribe, publish, resourceEvent, waitingList, subscribers}

	return FeedRoom
}

func InitSessionManager() {
	GlobalSessions, _ = session.NewManager("memory", `{"cookieName":"elasticfeedsessid","gclifetime":3600}`)
	go GlobalSessions.GC()
}
