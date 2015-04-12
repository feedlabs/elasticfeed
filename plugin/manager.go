package plugin

type PluginManager struct {
	Indexers        map[string]interface{}
	Crawlers        map[string]interface{}
	Sensors         map[string]interface{}
	Pipelines       map[string]interface{}
	Scenarios       map[string]interface{}
	Helpers         map[string]interface{}

	api            *ResourceApi
}

func (this *PluginManager) GetResourceApi() interface{} {
	return this.api
}

func (this *PluginManager) InitIndexer(name string, profiler *Profiler) *Plugin {
	p := NewPlugin(this, this.api, profiler)

	p.Init()
	this.Indexers[name] = p

	return p
}

func (this *PluginManager) FindPlugin(name string, profiler *Profiler) *interface{} {
	return nil
}

func (this *PluginManager) ExecPlugin(p Plugin) {
	//	profiler := p.profiler
}

func NewPluginManager(resourceManager interface{}) *PluginManager {
	pm := &PluginManager{}

	pm.api = NewResourceApi(resourceManager)

	return pm
}
