package model

import (
	pmodel "github.com/feedlabs/elasticfeed/plugin/model"
)

type PluginManager interface {

	LoadPipeline(name string) (pmodel.Pipeline, error)
}
