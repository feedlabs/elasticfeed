package store

import (
	"github.com/feedlabs/elasticfeed/service/store/v1/router"
	"github.com/feedlabs/elasticfeed/service/store/v1/controller"
)

type DbService struct {}

func (this *DbService) Init() {
	router.InitRouters()
	controller.InitService()
}

func NewDbService() *DbService {
	return &DbService{}
}
