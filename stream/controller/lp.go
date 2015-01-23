package controller

import (
	"github.com/feedlabs/feedify"
	"github.com/feedlabs/elasticfeed/stream/model"
)

type LongPollingController struct {
	feedify.Controller
}

func (this *LongPollingController) Join() {
	uname := this.GetString("uname")
	if len(uname) == 0 {
		return
	}

	Join(uname, nil)
}

func (this *LongPollingController) Post() {
	uname := this.GetString("uname")
	content := this.GetString("content")
	if len(uname) == 0 || len(content) == 0 {
		return
	}

	publish <- newEvent(model.EVENT_MESSAGE, uname, content)
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
	waitingList.PushBack(ch)
	<-ch

	this.Data["json"] = model.GetEvents(int(lastReceived))
	this.ServeJson()
}
