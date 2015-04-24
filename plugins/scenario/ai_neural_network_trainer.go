package scenario

import (
	_ "github.com/feedlabs/elasticfeed/elasticfeed/ai"
)

func AINeuralNetworkTrainer(data interface{}, viewerBrain interface{}) interface{} {

	//-----------------------------------------------------
	// Train ViewerBrain using Viewer specific metrics from
	// front-end behaviours
	//-----------------------------------------------------

	return viewerBrain
}

func init() {}
