package stream

import (
	"github.com/feedlabs/feedify"
)

func init() {
	feedify.Router("/ws", &WebSocketController{})
	feedify.Router("/ws/join", &WebSocketController{}, "get:Join")
}
