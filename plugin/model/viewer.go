package model

/*
	TODO:

	**********************************************************

	SHOULD BE BASICALLY A BRIDGE TO PUPULATION/PERSON CLASS

	**********************************************************

	- VIEWER IS USED IN DISTRIBUTION/PIPELINE WORKFLOW
	- VIEWER IS USED IN LEARNING/EVOLUTION/SCENARIO WORKFLOW
	- VIEWER IS UPDATED/USED IN SENSOR UPDATE WORKFLOW

	*****************************************************************************************
	VIEWER SHOULD BE GLOBALY REGISTERED AND UPDATED WITH SENSOR STATE BY SCHEDULE
	*****************************************************************************************
 */

type IViewer interface {
}

type Viewer struct {
}

func (this *Viewer) GetMetrics() {

	/*
		- GET HEAT MAPS
		- GET ACTIONS
		- GET HABITS
		- GET BEHAVIOURS
		- ...
	 */

}

func (this *Viewer) GetMetricsStats() {}

func (this *Viewer) GetMetricsForEntry() {

	/*
		- ENTRY METRICS/INDICES FOR SPECIFIC VIEWER
	 */

}

func (this *Viewer) GetMetricsForFeed() {

	/*
		- FEED METRICS/INDICES FOR SPECIFIC VIEWER
	 */

}

func (this *Viewer) GetGlobalIndicesForEntry() {

	/*
		- IT SHOULD BE A PROXY TO RESOURCE-API
		- IT SHOUD GET ALL MINDICES FOR ENTRY
	 */

}

func (this *Viewer) GetGlobalIndicesForFeed() {}

func (this *Viewer) GetStorages() {

	/*
		- GET ANNs BRAINS
		- GET SVMs
		- ...
	 */

}

func (this *Viewer) GetCurrentContext() {

	/*
		- IP
		- LOCATION
		- BROWSER
		- OS
		- SENSORS
		 - WEATHER
		 - DAY PART
		 - BIORYTHM
		 - ...
	 */

}

func (this *Viewer) GetEnvironment() {

	/*
		- SHOULD COMBINE
		 - CURRENT CONTEXT WITH...
		 - SENSORS WITH...
		 - METRICS...
	 */

}
