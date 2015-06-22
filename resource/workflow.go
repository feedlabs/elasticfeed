package resource

import (
	"strconv"
	"errors"

	"github.com/feedlabs/feedify/graph"
)

func (this *Workflow) GetRawData() map[string]interface{} {
	return make(map[string]interface{})
}

func (this *Workflow) GetProfilerRawData() map[string]string {

	// workaround - first GetRawData() needs to return real data
	data :=  make(map[string]string)
	data["default-profiler"] = "true"
	data["mem"] = "256"
	data["cpus"] = "4"
	data["bandwidth"] = "1000"
	return data

	return this.GetRawData()["profiler"].(map[string]string)
}

func GetWorkflowList(FeedId string, ApplicationId string, OrgId string) (feedWorkflows []*Workflow, err error) {
	feed, err := GetFeed(FeedId, ApplicationId, OrgId)
	if err != nil {
		return nil, err
	}

	_id, _ := strconv.Atoi(feed.Id)
	_rels, _ := storage.RelationshipsNode(_id, "workflow")

	var workflows []*Workflow

	for _, rel := range _rels {
		data := rel.EndNode.Data["data"].(string)
		def := rel.EndNode.Data["default"]

		if def == nil {
			def = false
		}

		workflow := NewWorkflow(strconv.Itoa(rel.EndNode.Id), feed, def.(bool), data)
		if workflow != nil && Contains(rel.EndNode.Labels, RESOURCE_WORKFLOW_LABEL) && feed.Id == rel.EndNode.Data["feedId"].(string) {
			workflows = append(workflows, workflow)
		}
	}

	if workflows == nil {
		workflows = make([]*Workflow, 0)
	}

	return workflows, nil
}

func GetWorkflow(id string, FeedId string, ApplicationId string, OrgId string) (feedWorkflow *Workflow, err error) {
	feed, err := GetFeed(FeedId, ApplicationId, OrgId)
	if err != nil {
		return nil, err
	}

	_id, err := strconv.Atoi(id)
	workflow, err := storage.Node(_id)

	if err != nil {
		return nil, err
	}

	if workflow != nil && Contains(workflow.Labels, RESOURCE_WORKFLOW_LABEL) && feed.Id == workflow.Data["feedId"].(string) {
		data := workflow.Data["data"].(string)
		def := workflow.Data["default"]

		if def == nil {
			def = false
		}

		return NewWorkflow(strconv.Itoa(workflow.Id), feed, def.(bool), data), nil
	}

	return nil, errors.New("WorkflowId `"+id+"` not exist")
}

func AddWorkflow(feedWorkflow Workflow, FeedId string, ApplicationId string, OrgId string) (WorkflowId string, err error) {

	// get feed
	feed, err := GetFeed(FeedId, ApplicationId, OrgId)
	if err != nil {
		return "0", err
	}

	// add feed-workflow
	properties := graph.Props{
		"feedId": feed.Id,
		"data": feedWorkflow.Data,
	}

	workflow, err := storage.NewNode(properties, RESOURCE_WORKFLOW_LABEL)

	if err != nil {
		return "0", err
	}

	// create relation
	_feedId, _ := strconv.Atoi(feed.Id)
	rel, err := storage.RelateNodes(_feedId, workflow.Id, "workflow", nil)

	if err != nil || rel.Type == "" {
		return "0", err
	}

	feedWorkflow.Id = strconv.Itoa(workflow.Id)

	return feedWorkflow.Id, nil
}

func UpdateWorkflow(id string, FeedId string, ApplicationId string, OrgId string, data string) (workflow *Workflow, err error) {
	workflow, err = GetWorkflow(id, FeedId, ApplicationId, OrgId)

	if err != nil {
		return nil, err
	}

	// update workflow
	workflow.Data = data

	_id, _ := strconv.Atoi(workflow.Id)

	err = storage.SetPropertyNode(_id, "data", data)

	if err != nil {
		return nil, err
	}

	return workflow, nil
}

func DeleteWorkflow(id string, FeedId string, ApplicationId string, OrgId string) (error) {
	workflow, err := GetWorkflow(id, FeedId, ApplicationId, OrgId)

	if err != nil {
		return err
	}

	_id, _ := strconv.Atoi(workflow.Id)
	_rels, _ := storage.RelationshipsNode(_id, "workflow")

	for _, rel := range _rels {
		storage.DeleteRelation(rel.Id)
	}

	return storage.DeleteNode(_id)
}

func NewWorkflow(id string, feed *Feed, def bool, data string) *Workflow {
	return &Workflow{id, feed, def, data}
}
