package channel

import (
	"net/http"

	"github.com/feedlabs/feedify"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"

	"github.com/feedlabs/elasticfeed/stream/model"

	"github.com/feedlabs/elasticfeed/stream/controller/room"
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

	room.Join(uname, ws)
	defer room.Leave(uname)

	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			return
		}
		room.Publish <- room.NewEvent(model.EVENT_MESSAGE, uname, string(p))

		room.System_rpc <- ws
	}
}
