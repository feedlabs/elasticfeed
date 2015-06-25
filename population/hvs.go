package population

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
}

func (this *VHSController) GetPVP() *PVPController {
	return this.pvp
}

func NewVirtualHumanSurrogate() *VHSController {
	return &VHSController{nil}
}
