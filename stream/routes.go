package stream

import (
	"github.com/feedlabs/feedify"
)

func init() {
	feedify.Router("/ws", &WebSocketController{})
	feedify.Router("/ws/join", &WebSocketController{}, "get:Join")

	feedify.Router("/lp", &LongPollingController{}, "get:Join")
	feedify.Router("/lp/post", &LongPollingController{})
	feedify.Router("/lp/fetch", &LongPollingController{}, "get:Fetch")
}
