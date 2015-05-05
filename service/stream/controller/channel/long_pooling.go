package channel

import (
	"github.com/feedlabs/elasticfeed/service/stream/model"
	"github.com/feedlabs/elasticfeed/service/stream/controller/room"
)

type LongPollingController struct {
	DefaultController
}

func (this *LongPollingController) Join() {
	chid := this.GetString("chid")
	if len(chid) == 0 {
		return
	}

	w := this.GetCtx().ResponseWriter
	r := this.GetCtx().Input.Request
	sess := room.GlobalSessions.SessionStart(w, r)
	defer sess.SessionRelease(w)

	room.FeedRoom.Join(chid, nil)

	list := make(map[string]interface {})
	list["response"] = room.NewChannelEvent(room.CHANNEL_JOIN, "system", "join")

	this.Data["json"] = list
	this.ServeJson()
}

func (this *LongPollingController) Post() {
	chid := this.GetString("chid")
	data := this.GetString("data")
	if len(chid) == 0 || len(data) == 0 {
		return
	}

	// Feature:
	// or specific request for this client;
	// should be executed and returned directly to user
	// lastReceived time should not be changed in that case

	//	room.Publish <- room.NewSystemEvent(room.CHANNEL_MESSAGE, chid, data)

	ch := make(chan []byte)

	room.FeedRoom.ResourceEvent <- room.NewSocketEvent([]byte(data), nil, ch)

	response := <-ch

	list := make(map[string]string)
	list["response"] = string(response)

	this.Data["json"] = list
	this.ServeJson()
}

func (this *LongPollingController) Fetch() {
	lastReceived, err := this.GetInt("lastReceived")
	if err != nil {
		return
	}

	events := model.GetEvents(int(lastReceived))
	if len(events) > 0 {
		this.Data["json"] = events
		this.ServeJson()
		return
	}

	// Wait for new message(s).
	ch := make(chan bool)
	room.FeedRoom.WaitingList.PushBack(ch)
	<-ch

	this.Data["json"] = model.GetEvents(int(lastReceived))
	this.ServeJson()
}
