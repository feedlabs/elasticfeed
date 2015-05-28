package workflow

import (
	"fmt"

	"encoding/json"
	"time"
	"math/rand"

	"github.com/feedlabs/elasticfeed/elasticfeed/model"
	"github.com/feedlabs/elasticfeed/resource"

	emodel "github.com/feedlabs/elasticfeed/elasticfeed/model"
	smodel "github.com/feedlabs/elasticfeed/service/stream/model"
	pmodel "github.com/feedlabs/elasticfeed/plugin/model"

	"github.com/feedlabs/elasticfeed/service/stream/controller/room"
	"github.com/feedlabs/elasticfeed/service/stream"

	jsonutil "github.com/mitchellh/packer/common/json"
	"github.com/mitchellh/mapstructure"
)


var (
	pluginManagerAnn pmodel.Pipeline
	entryListCache map[string][]*resource.Entry
)

type WorkflowManager struct {
	engine emodel.Elasticfeed

	pManager model.PluginManager
	eManager model.EventManager

	workflows []*WorkflowController
	template  interface{}
}

func (this *WorkflowManager) Init() {

	sss := `
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
	`
	//	var data []byte
	//	copy(data[:], sss)
	data := []byte(sss)

	var rawTplInterface interface{}
	err := jsonutil.Unmarshal(data, &rawTplInterface)
	if err != nil {
		fmt.Println(data)
		fmt.Println(err)
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
		fmt.Println("err2")
		fmt.Println(err)
		return
	}

	err = decoder.Decode(rawTplInterface)
	if err != nil {
		fmt.Println("err3")
		fmt.Println(err)
		return
	}

	fmt.Println(rawTpl)
	fmt.Println(rawTpl.Storing.NewEntryEvent.Indexers[0]["type"])
	fmt.Println(rawTpl.Storing.NewEntryEvent.Crawlers[0]["type"])


	this.BindServiceEvents()

	this.InstallSensorsSchedule()
	this.InstallFeedMaintenanceSchedule()
}

/**
	TODO:
	MAYBE COULD BIND TO "SYSTEM EVENT MANAGER"
	- COULD BIND TO RESOURCE EVENTS: NEW ENTRY, NEW METRIC, NEW VIEWER
	- COULD BIND TO CRON JOBS: FEED MAINTAINER, SENSORS UPDATE

	** OVERALL THE WORKFLOW MANAGER COULD BIND TO EVENTS AND ALSO CREATE OWN ALARMS/INTERRUPTS EVENTS
 */

/**
	TODO:

	 - IMPLEMENT EVENTS TRIGGERS

	 - IMPLEMENT STREAM SERVICE EVENT/HOOKS BINDING (LISTEN TO EVENTS AND HOOKS ON STREAM SERVICE)
	 - IMPLEMENT STORE SERVICE EVENT/HOOKS BINDING (SHOULD BE DONE BY "SYSTEM EVENTS MANAGER")

	 - IMPLEMENT LOCAL CRON JOB FOR ("SYSTEM EVENTS MANAGER")
	   - SENSOR REFRESH EVENT
	   - FEED MAINTAINER EVENT

	 - IMPLEMENT RESOURCE API WHICH
	   - CAN BE PASSED TO PLUGINS
	   - CAN PROVIDE/CREATE DATA
 */

func (this *WorkflowManager) GetStreamService() *stream.StreamService {
	return this.GetEngine().GetServiceManager().GetStreamService()
}

func (this *WorkflowManager) GetEngine() emodel.Elasticfeed {
	return this.engine
}

func (this *WorkflowManager) InitTemplate(t interface{}) {
	// verify event availability into EventsManager
	// verify hooks workflows
	this.template = t
}

func (this *WorkflowManager) CreateFeedWorkflow(feed *resource.Feed) *WorkflowController {
	w := NewWorkflowController(feed, this)
	w.Init()
	this.workflows = append(this.workflows, w)
	return w
}

func (this *WorkflowManager) BindServiceEvents() {
	// should bind service-stream-controllers to get handler to channel
	// should pass it down to listen for events on streaming controllers
	go this.BindStreamServiceEvents()
}

func (this *WorkflowManager) BindStreamServiceEvents() {

	for {
		select {
		case socketEvent := <-this.GetStreamService().GetFeedRoomManager().ResourceEvent:

			action := socketEvent.ActionId

			switch {
			case action == room.FEED_ENTRY_INIT || action == room.FEED_ENTRY_MORE:
				go this.ResourcePipelineRound(socketEvent)
			}

		}
	}
}

func (this *WorkflowManager) ResourcePipelineRound(socketEvent smodel.SocketEvent) {

	// TODO:

	// will run WorkflowManager with Pipeline plugins

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
		amt := time.Duration(100)
		time.Sleep(amt * time.Millisecond)
		timeout <- true
	}()

	// WORKFLOW PIPELINE
	go func(list []*resource.Entry, socketEvent smodel.SocketEvent) {

		if pluginManagerAnn == nil {
			pluginManagerAnn, _ = this.engine.GetPluginManager().LoadPipeline("ann")
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

func (this *WorkflowManager) InstallFeedMaintenanceSchedule() {
	e := this.GetEngine().GetEventManager()
	_ = e.InstallSchedule("feed", "0 12 * * * *", func() error {
		fmt.Println("hello feed schedule")
		return nil
	})
}

func (this *WorkflowManager) InstallSensorsSchedule() {
	e := this.GetEngine().GetEventManager()
	_ = e.InstallSchedule("sensor", "0 18 * * * *", func() error {
		fmt.Println("hello sensor schedule")
		return nil
	})
}

func NewWorkflowManager(engine emodel.Elasticfeed) *WorkflowManager {
	tpl := engine.GetConfig()

	wm := &WorkflowManager{engine, engine.GetPluginManager(), engine.GetEventManager(), nil, nil}
	wm.InitTemplate(tpl)

	return wm
}
