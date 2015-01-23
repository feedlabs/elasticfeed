package controller

import (
	"encoding/json"
	"net/http"

	"github.com/feedlabs/feedify"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"

	"github.com/feedlabs/elasticfeed/stream/model"
)

type WebSocketController struct {
	feedify.Controller
}

func (this *WebSocketController) Join() {
	uname := this.GetString("uname")
	if len(uname) == 0 {
		return
	}

	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}

	Join(uname, ws)
	defer Leave(uname)

	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			return
		}
		publish <- newEvent(model.EVENT_MESSAGE, uname, string(p))
	}
}

func broadcastWebSocket(event model.Event) {
	data, err := json.Marshal(event)
	if err != nil {
		beego.Error("Fail to marshal event:", err)
		return
	}

	for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
		ws := sub.Value.(Subscriber).Conn
		if ws != nil {
			if ws.WriteMessage(websocket.TextMessage, data) != nil {
				unsubscribe <- sub.Value.(Subscriber).Name
			}
		}
	}
}
