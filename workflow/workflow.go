package workflow

import (
	"github.com/feedlabs/elasticfeed/plugin/model"
	"github.com/feedlabs/elasticfeed/resource"
)

type WorkflowController struct {
	feed    *resource.Feed
	manager *WorkflowManager

	profiler *model.Profiler

	IndexerTimeout     int
	CrawlerTimeout     int
	SensorTimeout      int
	PipelineTimeout    int
	ScenarioTimeout    int
}

func (this *WorkflowController) GetManager() *WorkflowManager {
	return nil
}

func (this *WorkflowController) GetFeed() *interface{} {
	return nil
}

func (this *WorkflowController) GetProfiler() *model.Profiler {
	return this.profiler
}

func (this *WorkflowController) Init() {
	// verify Feed.Workflowfile stricture; does match WorkflowManager Templating
	// verify plugins availability: this.manager.findPlugin()
	// run Plugins if require specific Profiler
	// bind Feed to system Events: this.manager.BindToSystemEvents()
}

func (this *WorkflowController) DispatchIndexerHook(data interface{}) interface{} {
	return data
}

func (this *WorkflowController) DispatchPipelineHook(data interface{}) interface{} {
	return data
}

func NewWorkflowController(feed *resource.Feed, wm *WorkflowManager) *WorkflowController {
	data := feed.GetWorkflow().GetProfilerRawData()
	p := model.NewProfiler(data)
	w := &WorkflowController{feed, wm, p, 100, 100, 100, 100, 100}

	w.Init()

	return w
}
