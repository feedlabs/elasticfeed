package main

import (
	"github.com/feedlabs/elasticfeed/service"
	"github.com/feedlabs/elasticfeed/plugin"
	"github.com/feedlabs/elasticfeed/workflow"
	"github.com/feedlabs/elasticfeed/event"
	"github.com/feedlabs/elasticfeed/resource"
	"github.com/feedlabs/elasticfeed/elasticfeed"
)

var (
	ServerEngine *elasticfeed.Elasticfeed
)

func main() {
	eManager := event.NewEventManager()
	pManager := plugin.NewPluginManager()
	wManager := workflow.NewWorkflowManager(nil, pManager, eManager)

	ServerEngine = elasticfeed.NewElasticfeed(
		resource.NewResourceManager(),
		eManager,
		service.NewServiceManager(),
		pManager,
		wManager,
	)

	ServerEngine.Run()
}
