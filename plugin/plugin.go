package plugin

type Plugin struct {
	pluginManager *PluginManager
	resourceApi   *ResourceApi

	profiler      *Profiler
	rpcAddress    interface{}
}

func (this *Plugin) Init() {}

func NewPlugin(pm *PluginManager, api *ResourceApi, profiler *Profiler) *Plugin {
	return &Plugin{pm, api, profiler, ""}
}
