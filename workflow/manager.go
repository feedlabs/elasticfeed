package workflow

import (
	"fmt"

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
	entryListCache map[string][]*resource.Entry
)

type WorkflowManager struct {
	engine emodel.Elasticfeed

	pManager model.PluginManager
	eManager model.EventManager

	workflows map[string]*WorkflowController
	template  interface{}
}

func (this *WorkflowManager) Init() {

	this.workflows = make(map[string]*WorkflowController)

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

func (this *WorkflowManager) CreateWorkflowFeed(feed *resource.Feed) *WorkflowController {
	this.workflows[feed.Id] = NewWorkflowController(feed, this)
	return this.workflows[feed.Id]
}

func (this *WorkflowManager) BindServiceEvents() {
	// should bind service-stream-controllers to get handler to channel
	// should pass it down to listen for events on streaming controllers
	go this.BindStreamServiceEvents()
}

func (this *WorkflowManager) BindStreamServiceEvents() {

	/*
		- HANDLE API EVENTS LIKE NEW-ENTRY... OR..
		- BIND TO this.GetStreamService().GetFeedRoomManager().ResourceEvent FOR ANY POSSIBLE EVENTS?

		- MAYBE BIND TO SYSTEM EVENTS?
	 */

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

func (this *WorkflowManager) ResourceIndexerRound(socketEvent smodel.SocketEvent) {
}

func (this *WorkflowManager) ResourceCrawlerRound(socketEvent smodel.SocketEvent) {
}

func (this *WorkflowManager) ResourceSensorRound(socketEvent smodel.SocketEvent) {
}

func (this *WorkflowManager) ResourceScenarioRound(socketEvent smodel.SocketEvent) {
}

func (this *WorkflowManager) ResourcePipelineRound(socketEvent smodel.SocketEvent) {

	workflow, err := this.FindWorkflowByFeedId(socketEvent.FeedId)
	// create workflow for feed if not existing yet
	if err != nil {
		feed, _ := resource.GetFeed(socketEvent.FeedId, socketEvent.AppId, socketEvent.OrgId)
		this.CreateWorkflowFeed(feed)
	}

	workflow, err = this.FindWorkflowByFeedId(socketEvent.FeedId)
	if err != nil {
		fmt.Println("Cannot find Workflow for Feed ID: " + socketEvent.FeedId)
		return
	}

	workflow.ExecutePipelineChain(socketEvent)
}

func (this *WorkflowManager) FindWorkflowByFeedId(id string) (workflow *WorkflowController, err error) {
	if this.workflows[id] == nil {
		return nil, fmt.Errorf("Workflow for feedID %s does not exist!", id)
	}
	return this.workflows[id], nil
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
