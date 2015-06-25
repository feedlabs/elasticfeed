package population

/*
	***********************************
	*   Private Virtual Personality   *
	***********************************

	https://en.wikipedia.org/wiki/Talk%3APrivate_virtual_personality

	VI vs AI

	A VI is restricted to certain responses and actions, and can't become self-aware,
	while an AI is free to think how it likes and is self-aware.

 */

type PVPController struct {

	/*
		Basic:
		- SHOULD LEARNs FROM HUMAN
		- SHOULD SERVEs HUMAN
		- SHOULD ADVISEs HUMAN
		- SHOULD TEACHes HUMAN

		Advnaced:
		- SHOULD TALK TO OTHER PVPs
		- SHOULD ANSWER QUESTIONS

		Intelligence:
		- Intelligence Quotient (IQ)
		- Emotional Intelligence (EI)
	 */

}

func (this *PVPController) Train() {
	/*
		PVP LEARNS VIA HUMAN TEACHING
		- CARROT/STICK METHOD
		- MOOD DETECTION
		- ...
	 */
}

func (this *PVPController) Query() {
	/*
		QUESTION WITH LOGIC ANSWER: YES/NO

		QUESTION WITH PROBABILITY ANSWER: MAYBE YES, MAYBE NOT

		QUESTION ABOUT LIKENESS OF SUBJECT: LIKE/DISLIKE

		QUESTION ABOUT LIKENESS OF PRICE: LIKE/DISLIKE

		QUESTION ABOUT MOOD: GOOD, BAD, SAD
	 */
}

// SUPER GOAL: Artificial Intelligence
func (this *PVPController) GetAI() {
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

func NewPrivateVirtualPersonality() *PVPController {
	return &PVPController{}
}
