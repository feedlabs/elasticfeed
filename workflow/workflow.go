package workflow

import (
	"github.com/feedlabs/elasticfeed/plugin"
)

type Workflow struct {
	feed    interface{}
	manager *WorkflowManager

	profiler *plugin.Profiler
	data    map[string]interface{}
}

func (this *Workflow) GetManager() *WorkflowManager {
	return nil
}

func (this *Workflow) GetFeed() *interface{} {
	return nil
}

func (this *Workflow) GetProfiler() *plugin.Profiler {
	return this.profiler
}

func (this *Workflow) Init() {
	// verify Feed.Workflowfile stricture; does match WorkflowManager Templating
	// verify plugins availability: this.manager.findPlugin()
	// run Plugins if require specific Profiler
	// bind Feed to system Events: this.manager.BindToSystemEvents()
}

func (this *Workflow) DispatchIndexerHook(data interface{}) interface{} {
	return data
}

func (this *Workflow) DispatchPipelineHook(data interface{}) interface{} {
	return data
}

func NewWorkflow(data map[string]interface{}, f interface{}, wm *WorkflowManager) *Workflow {
	p := plugin.NewProfiler(data["profiler"].(map[string]string))
	w := &Workflow{f, wm, p, data}

	w.Init()

	return w
}
