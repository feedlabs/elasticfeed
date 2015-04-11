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
	em := event.NewEventManager()
	pm := plugin.NewPluginManager()
	wm := workflow.NewWorkflowManager(nil, pm, em)

	ServerEngine := elasticfeed.NewElasticfeed(resource.NewResourceManager(), em, service.NewServiceManager(), pm, wm)
	ServerEngine.Run()
}
