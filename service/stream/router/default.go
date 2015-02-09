package router

import (
	"github.com/feedlabs/feedify"
	"github.com/feedlabs/elasticfeed/service/stream/controller/channel"
)

func init() {
	feedify.SetStaticPath("/static", "service/stream/static")

	feedify.Router("/stream/lp/join", &channel.LongPollingController{}, "get:Join")
	feedify.Router("/stream/lp/post", &channel.LongPollingController{})
	feedify.Router("/stream/lp/fetch", &channel.LongPollingController{}, "get:Fetch")

	feedify.Router("/stream/ws/join", &channel.WebSocketController{}, "get:Join")

	feedify.Router("/stream/es/join", &channel.EventSourceController{}, "get:Join")
	feedify.Router("/stream/es/post", &channel.EventSourceController{})
}
