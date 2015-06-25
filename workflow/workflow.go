package workflow

import (
//	"fmt"

	"encoding/json"
	"time"
	"math/rand"

	"github.com/feedlabs/elasticfeed/plugin/model"
	"github.com/feedlabs/elasticfeed/resource"
	"github.com/feedlabs/elasticfeed/population"

	jsonutil "github.com/mitchellh/packer/common/json"
	"github.com/mitchellh/mapstructure"

	smodel "github.com/feedlabs/elasticfeed/service/stream/model"

	"github.com/feedlabs/elasticfeed/service/stream/controller/room"
)

/*
	WORKFLOW CONTROLLER SHOULD
	- REGISTER FEED WORKFLOW
	- SHOULD STORE SUBSCRIBED HUMANS/VIEWERS
	 - SHOULD BE ABLE TO READ/WRTIE METRICS/INDICES/STORAGES
	 - SHOULD BE ABLE TO ACCESS SENSORS STATUS FOR SPECIFIC WORKFLOW

 */

type WorkflowController struct {
	feed    *resource.Feed
	manager *WorkflowManager

	population map[string]*population.HumanController

	profiler *model.Profiler

	IndexerTimeout     int
	CrawlerTimeout     int
	SensorTimeout      int
	PipelineTimeout    int
	ScenarioTimeout    int
}

func (this *WorkflowController) GetManager() *WorkflowManager {
	return nil
}

func (this *WorkflowController) GetFeed() *interface{} {
	return nil
}

func (this *WorkflowController) GetProfiler() *model.Profiler {
	return this.profiler
}

func (this *WorkflowController) Init() {

	// TODO:
	// verify Feed.Workflowfile stricture; does match WorkflowManager Templating
	// verify plugins availability: this.manager.findPlugin()
	// run Plugins if require specific Profiler
	// bind Feed to system Events: this.manager.BindToSystemEvents()

	data := []byte(`
		{
			"description":"aaaa-aaaa-aaaa",
			"storing" : {
				"new-entry": {
					"indexers": [
						{
							"type": "ann"
						}
					],
					"crawlers": [
						{
							"type": "crawler-google"
						}
					]
				}
			}
		}
	`)

	var rawTplInterface interface{}
	err := jsonutil.Unmarshal(data, &rawTplInterface)
	if err != nil {
		//		fmt.Println(data)
		//		fmt.Println(err)
		return
	}

	// Decode the raw template interface into the actual rawTemplate
	// structure, checking for any extranneous keys along the way.
	var md mapstructure.Metadata
	var rawTpl rawTemplate
	decoderConfig := &mapstructure.DecoderConfig{
		Metadata: &md,
		Result:   &rawTpl,
	}

	decoder, err := mapstructure.NewDecoder(decoderConfig)
	if err != nil {
		//		fmt.Println("err2")
		//		fmt.Println(err)
		return
	}

	err = decoder.Decode(rawTplInterface)
	if err != nil {
		//		fmt.Println("err3")
		//		fmt.Println(err)
		return
	}


	/*
		- WE NEED PLUGINS UPLOADED AND INSTALLED
		- PLUGINS SHOULD BE ABLE TO RUN FOR SPECIFIC WORKFLOW
	 */

	//	fmt.Println(rawTpl)
	//	fmt.Println(rawTpl.Storing.NewEntryEvent.Indexers[0]["type"])
	//	fmt.Println(rawTpl.Storing.NewEntryEvent.Crawlers[0]["type"])

}

func (this *WorkflowController) ExecutePipelineChain(socketEvent smodel.SocketEvent) {

	// *******************************************************************
	// REAL CONTENT IMPROVEMENT
	// based on connected user (viewer) or users (audience)!: habits, behaviours, stats etc.
	// WORKFLOW PIPE: filtering, customization
	// WORKFLOW SCENARIO-ENGINE: scenarios SHOULD BE IMPLEMENTED ON METRIC SERVICE
	// *******************************************************************

	// *******************************************************************
	// SCENARIO AND RULES/METRICS
	// should use go routine with time limit to query filter rules
	// if in specific time there is no rules the results should be sent
	// client feed. After this the next package should be sent with
	// rules which entries should be remove/hidden from the view!
	// *******************************************************************

	timeout := make(chan bool, 1)
	results := make(chan []*resource.Entry, 1)

	// COLLECTING ENTRIES
	if entryListCache == nil {
		entryListCache = make(map[string][]*resource.Entry)
	}

	if entryListCache[socketEvent.FeedId] == nil {
		entryListCache[socketEvent.FeedId], _ = resource.GetEntryList(socketEvent.FeedId, socketEvent.AppId, socketEvent.OrgId)
	}

	// WORKFLOW TIMEOUT
	// !! SHOULD BE CONFIGURABLE OVER RUNTIME SETTING
	// !! DEFAULT VALUE SHOULD BE IN CONFIG FILE
	go func() {
		amt := time.Duration(this.PipelineTimeout)
		time.Sleep(amt * time.Millisecond)
		timeout <- true
	}()

	// WORKFLOW PIPELINE
	go func(list []*resource.Entry, socketEvent smodel.SocketEvent) {

		if pluginManagerAnn == nil {
			pluginManagerAnn, _ = this.manager.engine.GetPluginManager().LoadPipeline("ann")
			pluginManagerAnn.Prepare()
		}

		newList, _ := pluginManagerAnn.Run(list)

		var newEntryList []*resource.Entry

		for _, v := range newList.([]interface{}) {
			Id := ""
			Data := ""
			for k, vv := range v.(map[interface{}]interface{}) {
				if k == "Id" {
					Id = vv.(string)
				}
				if k == "Data" {
					Data = vv.(string)
				}
			}
			if Id != "" && Data != "" {
				newEntryList = append(newEntryList, &resource.Entry{Id, nil, Data})
			}
		}

		list = newEntryList

		results <- list
	}(entryListCache[socketEvent.FeedId], socketEvent)

	select {

	// IF PIPE TAKES TOO MUCH TIME, DATA DELAYED
	case <-timeout:

		event := room.NewFeedEvent(room.FEED_ENTRY_NEW, socketEvent.FeedId, "{Content:\"tiemout\"}")
		data, _ := json.Marshal(event)

		if socketEvent.Ws != nil {
			amt := time.Duration(rand.Intn(500)) * 1000
			time.Sleep(amt * time.Microsecond)
			socketEvent.Ws.WriteMessage(1, data)
		}

		if socketEvent.Ch != nil {
			socketEvent.Ch <- data
		}

	// IF DATA ARRIVES WITHOUT DELAY
	case list := <-results:

	// *********************************************************************
	// register socket handler
	// needs to send notiffication to long pooling + ws
	// join should generate uniqe ID and client should use it
	// maybe sessionID could be as uniqeID ?
	// room.FeedSubscribers[socketEvent.FeedId][channelID] = socketEvent
	// *********************************************************************

		d, _ := json.Marshal(list)
		event := room.NewFeedEvent(room.FEED_ENTRY_INIT, socketEvent.FeedId, string(d))
		data, _ := json.Marshal(event)

		if socketEvent.Ws != nil {
			socketEvent.Ws.WriteMessage(1, data)
		}

		if socketEvent.Ch != nil {
			socketEvent.Ch <- data
		}

	}
}

func (this * WorkflowController) GetHumanByUID(uid string) *population.HumanController {
	return this.population[uid]
}

func (this * WorkflowController) RegisterHuman(human *population.HumanController) {
	this.population[human.UID] = human
}

func (this * WorkflowController) UnregisterHuman(human *population.HumanController) {
	delete(this.population, human.UID)
}

func NewWorkflowController(feed *resource.Feed, wm *WorkflowManager) *WorkflowController {
	population := make(map[string]*population.HumanController)

	data := feed.GetWorkflow().GetProfilerRawData()
	p := model.NewProfiler(data)
	w := &WorkflowController{feed, wm, population, p, 100, 100, 100, 50, 100}

	w.Init()

	return w
}
