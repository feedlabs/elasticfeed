package stream

import (
	"github.com/feedlabs/elasticfeed/service/stream/router"
	"github.com/feedlabs/elasticfeed/service/stream/controller"
)

type StreamService struct {}

func (this *StreamService) Init() {
	router.InitRouters()
	controller.InitRooms()
}

func NewStreamService() *StreamService {
	return &StreamService{}
}
