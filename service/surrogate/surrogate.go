package surrogate

import (
	"github.com/feedlabs/elasticfeed/service/surrogate/v1/router"
)

/*
	TODO:
	- THIS SERVICE SHOULD ALLOW FOR WORKFLOW files
	- WORKFLOW SHOULD be:
	 - SCENARIO -> YES (learning)
	 - PIPELINE -> YES (prediction)
	 - SENSOR/INDEXER/CRAWLER/HELPER -> NO (no pre-learning, no environment)
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
