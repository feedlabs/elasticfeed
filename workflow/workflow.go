package workflow

import (
	"fmt"

	"github.com/feedlabs/elasticfeed/plugin/model"
	"github.com/feedlabs/elasticfeed/resource"

	jsonutil "github.com/mitchellh/packer/common/json"
	"github.com/mitchellh/mapstructure"
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

	// TODO:
	// verify Feed.Workflowfile stricture; does match WorkflowManager Templating
	// verify plugins availability: this.manager.findPlugin()
	// run Plugins if require specific Profiler
	// bind Feed to system Events: this.manager.BindToSystemEvents()

	data := []byte(`
		{
			"description":"aaaa-aaaa-aaaa",
			"storing" : {
				"new-entry": {
					"indexers": [
						{
							"type": "ann"
						}
					],
					"crawlers": [
						{
							"type": "crawler-google"
						}
					]
				}
			}
		}
	`)

	var rawTplInterface interface{}
	err := jsonutil.Unmarshal(data, &rawTplInterface)
	if err != nil {
		fmt.Println(data)
		fmt.Println(err)
		return
	}

	// Decode the raw template interface into the actual rawTemplate
	// structure, checking for any extranneous keys along the way.
	var md mapstructure.Metadata
	var rawTpl rawTemplate
	decoderConfig := &mapstructure.DecoderConfig{
		Metadata: &md,
		Result:   &rawTpl,
	}

	decoder, err := mapstructure.NewDecoder(decoderConfig)
	if err != nil {
		fmt.Println("err2")
		fmt.Println(err)
		return
	}

	err = decoder.Decode(rawTplInterface)
	if err != nil {
		fmt.Println("err3")
		fmt.Println(err)
		return
	}


	/*
		- WE NEED PLUGINS UPLOADED AND INSTALLED
		- PLUGINS SHOULD BE ABLE TO RUN FOR SPECIFIC WORKFLOW
	 */

	fmt.Println(rawTpl)
	fmt.Println(rawTpl.Storing.NewEntryEvent.Indexers[0]["type"])
	fmt.Println(rawTpl.Storing.NewEntryEvent.Crawlers[0]["type"])

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
