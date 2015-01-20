package stream

import (
	"github.com/feedlabs/feedify"
	"github.com/feedlabs/elasticfeed/stream/controller"
)

func init() {
	feedify.Router("/ws", &controller.WebSocketController{})
	feedify.Router("/ws/join", &controller.WebSocketController{}, "get:Join")

	feedify.Router("/lp", &controller.LongPollingController{}, "get:Join")
	feedify.Router("/lp/post", &controller.LongPollingController{})
	feedify.Router("/lp/fetch", &controller.LongPollingController{}, "get:Fetch")
}
