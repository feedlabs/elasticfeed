package elasticfeed

import (
	"github.com/feedlabs/elasticfeed/plugin"
	"github.com/feedlabs/elasticfeed/workflow"
	"github.com/feedlabs/elasticfeed/service"
	"github.com/feedlabs/elasticfeed/event"
	"github.com/feedlabs/elasticfeed/resource"

	"github.com/feedlabs/feedify"
)

type Elasticfeed struct {
	rm *resource.ResourceManager
	em *event.EventManager
	sm *service.ServiceManager
	pm *plugin.PluginManager
	wm *workflow.WorkflowManager
}

func (this *Elasticfeed) GetEventManager() *event.EventManager {
	return this.em
}

func (this *Elasticfeed) GetResourceManager() *resource.ResourceManager {
	return this.rm
}

func (this *Elasticfeed) GetServiceManager() *service.ServiceManager {
	return this.sm
}

func (this *Elasticfeed) GetPluginManager() *plugin.PluginManager {
	return this.pm
}

func (this *Elasticfeed) GetWorkflowManager() *workflow.WorkflowManager {
	return this.wm
}

func (this *Elasticfeed) Run() {
	this.GetResourceManager().Init()
	this.GetServiceManager().Init()

	feedify.SetStaticPath("/static", "public")
	feedify.Run()
}

func NewElasticfeed(rm *resource.ResourceManager, em *event.EventManager, sm *service.ServiceManager, pm *plugin.PluginManager, wm *workflow.WorkflowManager) *Elasticfeed {
	return &Elasticfeed{rm, em, sm, pm, wm}
}
