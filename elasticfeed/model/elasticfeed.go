package model

type Elasticfeed interface {

	GetEventManager() EventManager

	GetResourceManager() ResourceManager

	GetServiceManager() ServiceManager

	GetPluginManager() PluginManager

	GetWorkflowManager() WorkflowManager

	GetConfig() map[string]interface{}

	Run()
}
