package pipeline

import (
	plugin "github.com/feedlabs/elasticfeed/plugin/pipeline"
)

func Filter(data interface{}) interface{} {
	// should call plugins of type PIPELINE
	return plugin.RandomAnimator(data)
}

func init() {}
