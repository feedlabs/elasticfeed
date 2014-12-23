package resources

import (
	"errors"
	"strconv"

	"github.com/feedlabs/feedify/graph"
)

func init() {
	Applications = make(map[string]*Application)
}

func GetApplicationList() []*Application {
	nodes, err := storage.FindNodesByLabel(RESOURCE_APPLICATION_LABEL)
	if err != nil {
		nodes = nil
	}

	var applications []*Application

	for _, node := range nodes {
		data := node.Data["data"].(string)
		id := strconv.Itoa(node.Id)
		rels, _ := storage.RelationshipsNode(node.Id, "contains")

		application := &Application{id , data, len(rels)}
		applications = append(applications, application)
	}

	return applications
}

func GetApplication(id string) (application *Application, err error) {
	_id, err := strconv.Atoi(id)
	node, err := storage.Node(_id)

	if err != nil {
		return nil, err
	}

	if node != nil && contains(node.Labels, RESOURCE_APPLICATION_LABEL) {
		data := node.Data["data"].(string)
		rels, _ := storage.RelationshipsNode(node.Id, "contains")
		return &Application{strconv.Itoa(node.Id), data, len(rels)}, nil
	}

	return nil, errors.New("ApplicationId not exist")
}

func AddApplication(application Application) (id string, err error) {
	properties := graph.Props{"data": application.Data}
	_application, err := storage.NewNode(properties, RESOURCE_APPLICATION_LABEL)

	if err != nil {
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
