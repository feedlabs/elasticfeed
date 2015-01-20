package stream

import (
	"fmt"
	"github.com/feedlabs/feedify"
)

// LongPollingController handles long polling requests.
type LongPollingController struct {
	feedify.Controller
}

// Join method handles GET requests for LongPollingController.
func (this *LongPollingController) Join() {
	// Safe check.
	uname := this.GetString("uname")
	if len(uname) == 0 {
		this.Redirect("/", 302)
		return
	}

	fmt.Println(uname)

	// Join chat room.
	Join(uname, nil)
}

// Post method handles receive messages requests for LongPollingController.
func (this *LongPollingController) Post() {
	uname := this.GetString("uname")
	content := this.GetString("content")
	if len(uname) == 0 || len(content) == 0 {
		return
	}

	publish <- newEvent(EVENT_MESSAGE, uname, content)
}

// Fetch method handles fetch archives requests for LongPollingController.
func (this *LongPollingController) Fetch() {
	lastReceived, err := this.GetInt("lastReceived")
	if err != nil {
		return
	}

	events := GetEvents(int(lastReceived))
	if len(events) > 0 {
		this.Data["json"] = events
		this.ServeJson()
		return
	}

	// Wait for new message(s).
	ch := make(chan bool)
	waitingList.PushBack(ch)
	<-ch

	this.Data["json"] = GetEvents(int(lastReceived))
	this.ServeJson()
}
