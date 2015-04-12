package system

import (
	"github.com/feedlabs/elasticfeed/service/system/router"
)

type SystemService struct {}

func (this *SystemService) Init() {
	router.InitRouters()
}

func NewMetricService() *SystemService {
	return &SystemService{}
}
