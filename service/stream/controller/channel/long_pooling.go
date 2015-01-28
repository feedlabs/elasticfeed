package channel

import (
	"github.com/feedlabs/feedify"

	"github.com/feedlabs/elasticfeed/service/stream/model"
	"github.com/feedlabs/elasticfeed/service/stream/controller/room"
)

type LongPollingController struct {
	feedify.Controller
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

	room.Join(chid, nil)
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

	room.Publish <- room.NewSystemEvent(model.EVENT_MESSAGE, chid, data)
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
	room.WaitingList.PushBack(ch)
	<-ch

	this.Data["json"] = model.GetEvents(int(lastReceived))
	this.ServeJson()
}
