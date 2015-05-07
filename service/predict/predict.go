package predict

import (
	"github.com/feedlabs/elasticfeed/service/predict/v1/router"
)

type PredictService struct {}

func (this *PredictService) Init() {
	router.InitStatusRouters()
	router.InitPredictRouters()
	router.InitTrainRouters()
}

func NewPredictService() *PredictService {
	return &PredictService{}
}
