package population

/*

	Private Virtual Personality

	https://en.wikipedia.org/wiki/Talk%3APrivate_virtual_personality

	VI vs AI

	A VI is restricted to certain responses and actions, and can't become self-aware,
	while an AI is free to think how it likes and is self-aware.

 */

type PVPControler struct {}

// SUPER GOAL: Artificial Intelligence
func (this *PVPControler) GetAI() {
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

func NewPrivateVirtualPersonality() *PVPControler {
	return &PVPControler{}
}
