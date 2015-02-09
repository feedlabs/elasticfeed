package channel

import (
	"github.com/feedlabs/feedify"

//	"github.com/mroth/sseserver"
)

type EventSourceController struct {
	feedify.Controller
}

func (this *EventSourceController) Join() {
}

func (this *EventSourceController) Post() {
}
