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
func (this *HumanController) GetPVP() *PVPControler {

	/*
		- SHOULD LEARN FROM HUMAN
		- SHOULD ADVISE HUMAN
		- SHOULD TEACH HUMAN
	 */

	return NewPrivateVirtualPersonality()
}

// SUPER GOAL: Artificial Intelligence
func (this *HumanController) GetAI() {
	/*
		SHOULD IMPLEMENT BRIDGE TO AI FRAMEWORK
		- SHOULD ALLOW TO REGISTER BRAINS
		- SHOULD ALLOW TO TRAIN/FORWARD OVER BRAINS
		- SHOULD ALLOW FOR EVOLVING
		- SHOULD ALLOW FOR GENETIC ALGORITHMS
		- SHOULD ALLOW FOR SIMPLE SVM
		- SHOULD ALLOW FOR ANY KIND OF AI

		********************************************
		LOOKS LIKE IMPORTANT PART
		IT WILL BE THE GOAL OF ELASTICFEED
		TO MAKE IT WORKING.

		THE COMBINATION OF:
		- TRAINING USING CURRENT SENSORS
		- CURRENT BRAIN STATE
		- CURRENT INDICES
		- CURRENT METRICS OF ACTIONS, BEHAVIOURS
		- CURRENT HABITS
		- ...
		********************************************
	 */
}

func NewHumanController() *HumanController {
	return &HumanController{"0", 0}
}
