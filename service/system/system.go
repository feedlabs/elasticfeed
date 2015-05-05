package system

import (
	"github.com/feedlabs/elasticfeed/service/system/v1/router"
)

type SystemService struct {}

func (this *SystemService) Init() {
	router.InitRouters()
	router.InitPluginRouters()
}

func NewSystemService() *SystemService {
	return &SystemService{}
}
