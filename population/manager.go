package population

import (
	emodel "github.com/feedlabs/elasticfeed/elasticfeed/model"
)

type PopulationManager struct {
	engine emodel.Elasticfeed

	people map[string]*HumanController
}

func (this *PopulationManager) Init() {
	this.people = make(map[string]*HumanController)
}

func NewPopulationManager(engine emodel.Elasticfeed) *PopulationManager {
	return &PopulationManager{engine, nil}
}
