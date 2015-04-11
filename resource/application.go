package resource

import (
	"errors"
	"strconv"

	"github.com/feedlabs/feedify/graph"
)

func (this *Application) AddFeed(feed Feed) (id string, err error) {
	return AddFeed(feed, this.Id, this.Org.Id)
}

func (this *Application) GetFeed(id string) (feed *Feed, err error) {
	return GetFeed(id, this.Id, this.Org.Id)
}

func (this *Application) GetFeedList() ([]*Feed, error) {
	return GetFeedList(this.Id, this.Org.Id)
}

func GetApplicationList(orgId string) (appList []*Application, err error) {
	org, err := GetOrg(orgId)
	if err != nil {
		return nil, err
	}

	_id, _ := strconv.Atoi(org.Id)
	_rels, _ := storage.RelationshipsNode(_id, "app")

	var applications []*Application

	for _, rel := range _rels {
		data := rel.EndNode.Data["data"].(string)
		id := strconv.Itoa(rel.EndNode.Id)
		rels, _ := storage.RelationshipsNode(rel.EndNode.Id, "feed")

		application := &Application{id, org, data, len(rels)}
		applications = append(applications, application)
	}

	if applications == nil {
		applications = make([]*Application, 0)
	}

	return applications, nil
}

func GetApplication(id string, orgId string) (application *Application, err error) {
	org, err := GetOrg(orgId)
	if err != nil {
		return nil, err
	}

	_id, err := strconv.Atoi(id)
	node, err := storage.Node(_id)

	if err != nil {
		return nil, err
	}

	if node != nil && Contains(node.Labels, RESOURCE_APPLICATION_LABEL) {
		data := node.Data["data"].(string)
		rels, _ := storage.RelationshipsNode(node.Id, "feed")
		return &Application{strconv.Itoa(node.Id), org, data, len(rels)}, nil
	}

	return nil, errors.New("ApplicationId not exist")
}

func AddApplication(application Application, orgId string) (id string, err error) {
	// get org
	org, err := GetOrg(orgId)
	if err != nil {
		return "0", err
	}

	// add app
	properties := graph.Props{
		"data": application.Data,
	}
	_application, err := storage.NewNode(properties, RESOURCE_APPLICATION_LABEL)

	if err != nil {
		return "0", err
	}

	// create relation
	_orgId, _ := strconv.Atoi(org.Id)
	rel, err := storage.RelateNodes(_orgId, _application.Id, "app", nil)

	if err != nil || rel.Type == "" {
		return "0", err
	}

	application.Id = strconv.Itoa(_application.Id)

	return application.Id, nil
}

func UpdateApplication(id string, data string) (err error) {
	_id, _ := strconv.Atoi(id)
	return storage.SetPropertyNode(_id, "data", data)
}

func DeleteApplication(id string) (error) {
	_id, _ := strconv.Atoi(id)
	return storage.DeleteNode(_id)
}
