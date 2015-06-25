package population

import (
	emodel "github.com/feedlabs/elasticfeed/elasticfeed/model"
)

type PopulationManager struct {
	engine emodel.Elasticfeed

	people map[string]*HumanController
	societies map[string]*SocietyController
}

func (this *PopulationManager) Init() {
	this.people = make(map[string]*HumanController)
	this.societies = make(map[string]*SocietyController)
}

func NewPopulationManager(engine emodel.Elasticfeed) *PopulationManager {
	return &PopulationManager{engine, nil, nil}
}
