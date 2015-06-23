package model

/*
	TODO:

	- VIEWER IS USED IN DISTRIBUTION/PIPELINE WORKFLOW
	- VIEWER IS USED IN LEARNING/EVOLUTION/SCENARIO WORKFLOW
	- VIEWER IS UPDATED/USED IN SENSOR UPDATE WORKFLOW

	*****************************************************************************************
	VIEWER SHOULD BE GLOBALY REGISTERED AND UPDATED WITH SENSOR STATE BY SCHEDULE
	*****************************************************************************************
 */

type Viewer interface {
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
		- ENTRY METRICS FOR SPECIFIC VIEWER
	 */

}

func (this *Viewer) GetMetricsForFeed() {

	/*
		- FEED METRICS FOR SPECIFIC VIEWER
	 */

}

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
