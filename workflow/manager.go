package workflow

import (
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
)


var (
	pluginManagerAnn pmodel.Pipeline
	entryListCache []*resource.Entry
)

type WorkflowManager struct {
	engine emodel.Elasticfeed

	pManager model.PluginManager
	eManager model.EventManager

	workflows []*WorkflowController
	template  interface{}
}

func (this *WorkflowManager) Init() {
	this.BindServiceEvents()
}

/**
	MAYBE COULD BIND TO "SYSTEM EVENT MANAGER"
	- COULD BIND TO RESOURCE EVENTS: NEW ENTRY, NEW METRIC, NEW VIEWER
	- COULD BIND TO CRON JOBS: FEED MAINTAINER, SENSORS UPDATE

	** OVERALL THE WORKFLOW MANAGER COULD BIND TO EVENTS AND ALSO CREATE OWN ALARMS/INTERRUPTS EVENTS
 */

/**
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
	// verify event availability into EventsManger
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
		entryListCache, _ = resource.GetEntryList(socketEvent.FeedId, socketEvent.AppId, socketEvent.OrgId)
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
	}(entryListCache, socketEvent)

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

func NewWorkflowManager(engine emodel.Elasticfeed, tpl interface{}, pm model.PluginManager, em model.EventManager) *WorkflowManager {
	// load template if not passed
	if tpl == nil {
		tpl = make(map[string]interface {})
	}

	wm := &WorkflowManager{engine, pm, em, nil, nil}
	wm.InitTemplate(tpl)

	return wm
}
