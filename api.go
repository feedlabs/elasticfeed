package main

import (
	"github.com/feedlabs/elasticfeed/service"
	"github.com/feedlabs/elasticfeed/plugin"
	"github.com/feedlabs/elasticfeed/workflow"
	"github.com/feedlabs/elasticfeed/event"
	"github.com/feedlabs/elasticfeed/resource"
	"github.com/feedlabs/elasticfeed/elasticfeed"
)

func main() {
	rm := resource.NewResourceManager()
	em := event.NewEventManager()
	pm := plugin.NewPluginManager(rm)
	wm := workflow.NewWorkflowManager(nil, pm, em)
	sm := service.NewServiceManager()

	ServerEngine := elasticfeed.NewElasticfeed(rm, em, sm, pm, wm)
	ServerEngine.Run()
}
