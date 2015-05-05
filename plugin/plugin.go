package plugin

import (
	"github.com/feedlabs/elasticfeed/plugin/model"
)

type Plugin struct {
	plugin             interface {}

	pluginManager      *PluginManager
	resourceApi        *model.ResourceApi

	profiler           *model.Profiler
	rpcAddress         interface{}

	pid                int
}

func (this *Plugin) Init() {}

func (this *Plugin) Run() (err error) {
	return nil
}

func (this *Plugin) GetPid() int {
	return this.pid
}

func NewPlugin(p *interface {}, pm *PluginManager, api *model.ResourceApi, profiler *model.Profiler) *Plugin {
	return &Plugin{p, pm, api, profiler, "", -1}
}
