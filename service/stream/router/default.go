package router

import (
	"github.com/feedlabs/feedify"
	"github.com/feedlabs/elasticfeed/service/stream/controller/channel"
)

func init() {
	feedify.SetStaticPath("/static", "service/stream/static")

	feedify.Router("/service/stream/lp/join", &channel.LongPollingController{}, "get:Join")
	feedify.Router("/service/stream/lp/post", &channel.LongPollingController{})
	feedify.Router("/service/stream/lp/fetch", &channel.LongPollingController{}, "get:Fetch")

	feedify.Router("/service/stream/ws/join", &channel.WebSocketController{}, "get:Join")
}
