package controller

import (
	"github.com/feedlabs/elasticfeed/service/stream/controller/room"
)

func InitSession() {
	room.InitSessionManager()
}
