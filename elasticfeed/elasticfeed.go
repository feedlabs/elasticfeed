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
	rManager *resource.ResourceManager
	eManager *event.EventManager
	sManager *service.ServiceManager
	pManager *plugin.PluginManager
	wManager *workflow.WorkflowManager
}

func (this *Elasticfeed) GetEventsManager() *event.EventManager {
	return this.eManager
}

func (this *Elasticfeed) Run() {
	this.rManager.Init()
	this.sManager.Init()

	feedify.SetStaticPath("/static", "public")
	feedify.Run()
}

func NewElasticfeed(rm *resource.ResourceManager, em *event.EventManager, sm *service.ServiceManager, pm *plugin.PluginManager, wm *workflow.WorkflowManager) *Elasticfeed {
	return &Elasticfeed{rm, em, sm, pm, wm}
}
