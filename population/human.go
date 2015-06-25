package population

/*
	IT DESCRIBE PERSON AS ENTITY OF PERSONALITY

	IT HAS:
	- HABITS
	- ACTIONS
	- BEHAVIOURS
	- MOODS
	- PHASES
	- ...
 */

const (
	HUMAN_SEX_FEMALE = 1
	HUMAN_SEX_MALE = 2
)

type HumanController struct {
	UID	string

	pvps map[string]*PVPController

	sex int
}

func (this *HumanController) GetSex() int {
	return this.sex
}

func (this *HumanController) SetSex(sex int) {
	this.sex = sex
}

// experimental thinking...
func (this *HumanController) GetMoods() {}
func (this *HumanController) GetLActivityHours() {}

// SUPER SUPER GOAL: Private Virtual Personality
func (this *HumanController) GetNewPVP() *PVPController {
	return NewPrivateVirtualPersonality()
}

func NewHumanController() *HumanController {
	return &HumanController{"0", nil, 0}
}
