package plugin

type PluginManager struct {
	Indexers        []interface{}
	Crawlers        []interface{}
	Sensors         []interface{}
	Pipelines       []interface{}
	Scenarios       []interface{}
	Helpers         []interface{}
}

func NewPluginManager() *PluginManager {
	return &PluginManager{}
}
