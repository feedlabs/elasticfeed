package workflow

import (
	"github.com/feedlabs/elasticfeed/elasticfeed/model"

	"github.com/feedlabs/elasticfeed/resource"
)

type WorkflowManager struct {
	pManager model.PluginManager
	eManager model.EventManager

	workflows []*WorkflowController
	template  interface{}
}

/**
 - MOVE FEED ROOM MANAGER EVENT LOGIC FROM RESOURCE MANAGER TO WORKFLOW MANAGER CLASS

 - IMPLEMENT EVENTS TRIGGERS

 - IMPLEMENT STREAM SERVICE EVENT BINDING
 - IMPLEMENT STORE SERVICE EVENT BINDING

 - IMPLEMENT LOCAL CRON JOB FOR
   - SENSOR REFRESH EVENT
   - FEED MAINTAINER EVENT

 - IMPLEMENT RESOURCE API WHICH
   - CAN BE PASSED TO PLUGINS
   - CAN PROVIDE/CREATE DATA

 */

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

func NewWorkflowManager(tpl interface{}, pm model.PluginManager, em model.EventManager) *WorkflowManager {
	// load template if not passed
	if tpl == nil {
		tpl = make(map[string]interface {})
	}

	wm := &WorkflowManager{pm, em, nil, nil}
	wm.InitTemplate(tpl)

	return wm
}
