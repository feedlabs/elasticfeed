package workflow

type Workflow struct {
	feed interface {}
	manager *WorkflowManager
}

func (this *Workflow) GetManager() *WorkflowManager {
	return nil
}

func (this *Workflow) GetFeed() *interface {} {
	return nil
}

func (this *Workflow) Init() {
	// verify Feed.Workflowfile stricture; does match WorkflowManager Templating
	// verify plugins availability: this.manager.findPlugin()
	// bind Feed to system Events: this.manager.BindToSystemEvents()
}

func (this *Workflow) DispatchIndexerHook(data interface{}) interface{} {
	return data
}

func (this *Workflow) DispatchPipelineHook(data interface{}) interface{} {
	return data
}

func NewWorkflow(f interface {}, wm *WorkflowManager) *Workflow {
	w := &Workflow{f, wm}
	w.Init()
	return w
}
