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
