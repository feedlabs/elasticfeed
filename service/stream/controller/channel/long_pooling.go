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
	uname := this.GetString("uname")
	if len(uname) == 0 {
		return
	}

	room.Join(uname, nil)
}

func (this *LongPollingController) Post() {
	uname := this.GetString("uname")
	content := this.GetString("content")
	if len(uname) == 0 || len(content) == 0 {
		return
	}

	// Feature:
	// or specific request for this client;
	// should be executed and returned directly to user
	// lastReceived time should not be changed in that case

	room.Publish <- room.NewEvent(model.EVENT_MESSAGE, uname, content)
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
