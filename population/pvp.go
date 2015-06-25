package population

/*

	***********************************
	*   Private Virtual Personality   *
	***********************************

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
		- ...

		TRAIN PACKAGE:
		- HUMAN ENVIRONMENT SENSORS STATUS
		- SUBJECT KNOWLEDGE
		- HUMAN DECISION
		- LIFETIME ACTIVITY&KNOWLEDGE
	 */
}

func (this *PVPController) Query() interface{} {
	/*
		QUESTION WITH LOGIC ANSWER: YES/NO

		QUESTION WITH PROBABILITY ANSWER: MOST LIKELY, MAYBE, PROBABLY YES/NO

		QUESTION ABOUT LIKENESS OF SUBJECT: LIKE/DISLIKE

		QUESTION ABOUT LIKENESS OF PRICE: CHEAP/LUXURY

		QUESTION ABOUT MOOD: GOOD, BAD, SAD

		QUESTION ABOUT BEHAVIOUR TODAY/TOMORROW: IDLE, SHOPPING, READING...

		QUESTION ABOUT ABILITY TO LEARN: LOW/MEDIUM/HIGH
	 */

	return map[string]interface{}{
		"query": "response",
	}
}

func (this *PVPController) AskMood(time int) interface{} {
	return this.Query()
}

func (this *PVPController) AskLike(subject interface{}) interface{} {
	return this.Query()
}

func (this *PVPController) AskPriceAcceptance(subject interface{}, price int) interface{} {
	return this.Query()
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
		- TRAINING USING CURRENT DECISIONS
		- CURRENT BRAIN STATE
		- CURRENT KNOWLEDGE/INDICES
		- CURRENT METRICS OF BEHAVIOURS, HABITS, ABILITIES, MOODS
		- ...
		********************************************
	 */
}

func NewPrivateVirtualPersonality() *PVPController {
	return &PVPController{}
}
