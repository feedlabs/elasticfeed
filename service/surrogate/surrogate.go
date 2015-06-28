package surrogate

import (
	"github.com/feedlabs/elasticfeed/service/surrogate/v1/router"
)

/*
	TODO:
	- SHOULD BE A BRIDGE TO PVP
	- SHOULD BE A BRIDGE VIA VHS
	- ..
 */

type SurrogateService struct {}

func (this *SurrogateService) Init() {
	router.InitStatusRouters()
	router.InitQueryRouters()
	router.InitTrainRouters()
}

func NewSurrogateService() *SurrogateService {
	return &SurrogateService{}
}
