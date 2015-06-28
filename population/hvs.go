package population

import (
	"github.com/feedlabs/elasticfeed/population/vhs"
)

/*

	*******************************
	*   Virtual Human Surrogate   *
	*******************************


	- SHOULD DEFINE PRIVATE VIRTUAL PERSONALITY
	- SHOULD DEFINE INPUT LANGUAGE AND CHANNEL
	- SHOULD DEFINE OUTPUT LANGUAGE AND CHANNEL
	- SHOULD DEFINE MOBILITY FEATURES
	- ...

*/

type VHSController struct {
	pvp *PVPController

	communicator vhs.Communicator
}

func (this *VHSController) Query(query interface {}) interface {} {
	return this.communicator.Query(this.pvp, query)
}

func (this *VHSController) Train(train interface {}) interface {} {
	return this.communicator.Query(this.pvp, train)
}

func (this *VHSController) GetPVP() *PVPController {
	return this.pvp
}

func (this *VHSController) GetCommunicator() vhs.Communicator {
	return this.communicator
}

func NewVirtualHumanSurrogate(communicator vhs.Communicator) *VHSController {
	return &VHSController{nil, communicator}
}
