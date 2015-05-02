package channel

import (
	"net/http"
	"encoding/json"

	"github.com/feedlabs/feedify"
	"github.com/gorilla/websocket"

	"github.com/feedlabs/elasticfeed/service/stream/controller/room"
)

type WebSocketController struct {
	DefaultController
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

	room.FeedRoom.Join(chid, ws)
	defer room.FeedRoom.Leave(chid)

	data, _ := json.Marshal(room.NewChannelEvent(room.CHANNEL_JOIN, chid, "join"))
	ws.WriteMessage(websocket.TextMessage, data)

	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			return
		}

		room.FeedRoom.ResourceEvent <- room.NewSocketEvent(p, ws, nil)
	}
}
