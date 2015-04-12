package resource

type ResourceManager struct {}

func (this * ResourceManager) Init() {
	InitStorage()
	InitResources()
	InitStreamCommunicator()
}

func NewResourceManager() *ResourceManager {
	return &ResourceManager{}
}
