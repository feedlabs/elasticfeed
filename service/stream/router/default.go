package router

import (
	"github.com/feedlabs/feedify"
	"github.com/feedlabs/elasticfeed/service/stream/controller/channel"
)

func init() {
	feedify.SetStaticPath("/static", "service/stream/static")

	feedify.Router("/lp/join", &channel.LongPollingController{}, "get:Join")
	feedify.Router("/lp/post", &channel.LongPollingController{})
	feedify.Router("/lp/fetch", &channel.LongPollingController{}, "get:Fetch")

	feedify.Router("/ws/join", &channel.WebSocketController{}, "get:Join")
}
