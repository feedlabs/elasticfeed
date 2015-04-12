package router

import (
	"github.com/feedlabs/feedify"
	"github.com/feedlabs/elasticfeed/service/stream/controller/channel"
)

func InitRouters() {
	feedify.Router("/stream/lp/join", &channel.LongPollingController{}, "get:Join")
	feedify.Router("/stream/lp/post", &channel.LongPollingController{})
	feedify.Router("/stream/lp/fetch", &channel.LongPollingController{}, "get:Fetch")

	feedify.Router("/stream/ws/join", &channel.WebSocketController{}, "get:Join")

	feedify.Router("/stream/sse/join", &channel.SSEController{}, "get:Join")
	feedify.Router("/stream/sse/post", &channel.SSEController{})
}
