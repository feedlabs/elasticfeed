package channel

import (
	"net/http"

	"github.com/feedlabs/feedify"
	"github.com/gorilla/websocket"

	"github.com/feedlabs/elasticfeed/service/stream/model"
	"github.com/feedlabs/elasticfeed/service/stream/controller/room"
)

type WebSocketController struct {
	feedify.Controller
}

func (this *WebSocketController) Join() {
	chid := this.GetString("chid")
	if len(chid) == 0 {
		return
	}

	w := this.GetCtx().ResponseWriter
	r := this.GetCtx().Input.Request
	sess := room.GlobalSessions.SessionStart(w, r)
	defer sess.SessionRelease(w)

	ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		feedify.Error("Cannot setup WebSocket connection:", err)
		return
	}

	room.Join(chid, ws)
	defer room.Leave(chid)

	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			return
		}

		room.Publish <- room.NewSystemEvent(model.EVENT_MESSAGE, chid, string(p))
		room.P2P <- ws
	}
}
