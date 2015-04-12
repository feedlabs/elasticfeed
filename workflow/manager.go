package workflow

import (
	"github.com/feedlabs/elasticfeed/plugin"
	"github.com/feedlabs/elasticfeed/event"
)

type WorkflowManager struct {
	pManager *plugin.PluginManager
	eManager *event.EventManager

	workflows []*Workflow
	template  interface{}
}

func (this *WorkflowManager) InitTemplate(t interface{}) {
	// verify event availability into EventsManger
	// verify hooks workflows
	this.template = t
}

func (this *WorkflowManager) CreateFeedWorkflow(f interface {}, data map[string]interface {}) *Workflow {
	w := NewWorkflow(data, f, this)
	w.Init()
	this.workflows = append(this.workflows, w)
	return w
}

func NewWorkflowManager(tpl interface{}, pm *plugin.PluginManager, em *event.EventManager) *WorkflowManager {
	// load template if not passed
	if tpl == nil {
		tpl = make(map[string]interface {})
	}

	wm := &WorkflowManager{pm, em, nil, nil}
	wm.InitTemplate(tpl)

	return wm
}
