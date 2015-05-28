package predict

import (
	"github.com/feedlabs/elasticfeed/service/predict/v1/router"
)

/*
	TODO:
	- THIS SERVICE SHOULD ALLOW FOR WORKFLOW files
	- WORKFLOW SHOULD be:
	 - SCENARIO -> YES (learning)
	 - PIPELINE -> YES (prediction)
	 - SENSOR/INDEXER/CRAWLER/HELPER -> NO (no pre-learning, no environment)
 */

type PredictService struct {}

func (this *PredictService) Init() {
	router.InitStatusRouters()
	router.InitPredictRouters()
	router.InitTrainRouters()
}

func NewPredictService() *PredictService {
	return &PredictService{}
}
