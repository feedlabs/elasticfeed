package resources

import (
	"errors"
	"strconv"

	"github.com/feedlabs/feedify/graph"
)

func (this *Org) AddApplication(app Application) (id string, err error) {
	return AddApplication(app, this.Id)
}

func GetOrgList() []*Org {
	nodes, err := storage.FindNodesByLabel(RESOURCE_ORG_LABEL)
	if err != nil {
		nodes = nil
	}

	var orgs []*Org

	for _, node := range nodes {
		data := node.Data["data"].(string)
		id := strconv.Itoa(node.Id)
		admin_rels, _ := storage.RelationshipsNode(node.Id, "admin")
		app_rels, _ := storage.RelationshipsNode(node.Id, "app")
		token_rels, _ := storage.RelationshipsNode(node.Id, "token")

		if node.Data["apiKey"] == nil {
			node.Data["apiKey"] = ""
		}

		org := &Org{id , node.Data["apiKey"].(string), data, len(token_rels), len(admin_rels), len(app_rels)}
		orgs = append(orgs, org)
	}

	return orgs
}

func GetOrg(id string) (org *Org, err error) {
	_id, err := strconv.Atoi(id)
	node, err := storage.Node(_id)

	if err != nil {
		return nil, err
	}

	if node != nil && contains(node.Labels, RESOURCE_ORG_LABEL) {
		data := node.Data["data"].(string)
		admin_rels, _ := storage.RelationshipsNode(node.Id, "admin")
		app_rels, _ := storage.RelationshipsNode(node.Id, "app")
		token_rels, _ := storage.RelationshipsNode(node.Id, "token")

		if node.Data["apiKey"] == nil {
			node.Data["apiKey"] = ""
		}

		return &Org{strconv.Itoa(node.Id), node.Data["apiKey"].(string), data, len(token_rels), len(admin_rels), len(app_rels)}, nil
	}

	return nil, errors.New("OrgId `"+id+"` not exist")
}

func AddOrg(org Org) (id string, err error) {
	properties := graph.Props{"data": org.Data}
	_org, err := storage.NewNode(properties, RESOURCE_ORG_LABEL)

	if err != nil {
		return "0", err
	}

	org.Id = strconv.Itoa(_org.Id)

	return org.Id, nil
}

func UpdateOrg(id string, data string) (err error) {
	_id, _ := strconv.Atoi(id)
	return storage.SetPropertyNode(_id, "data", data)
}

func DeleteOrg(id string) (error) {
	_id, _ := strconv.Atoi(id)
	return storage.DeleteNode(_id)
}

func init() {
	Orgs = make(map[string]*Org)
}
