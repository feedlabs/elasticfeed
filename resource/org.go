package resource

import (
	"errors"
	"strconv"

	"github.com/feedlabs/feedify/graph"
)

func (this *Org) AddApplication(app Application) (id string, err error) {
	return AddApplication(app, this.Id)
}

func (this *Org) AssignAdmin(adminId int) bool {
	_orgId, _ := strconv.Atoi(this.Id)
	rel, err := storage.RelateNodes(_orgId, adminId, "admin", nil)

	if err != nil || rel.Type == "" {
		return false
	}

	return true
}

func GetOrgList() (orgList []*Org, err error) {
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

		if node.Data["name"] == nil {
			node.Data["name"] = ""
		}

		org := &Org{id , node.Data["name"].(string), data, len(token_rels), len(admin_rels), len(app_rels)}
		orgs = append(orgs, org)
	}

	if orgs == nil {
		orgs = make([]*Org, 0)
	}

	return orgs, nil
}

func GetOrg(id string) (org *Org, err error) {
	_id, err := strconv.Atoi(id)
	node, err := storage.Node(_id)

	if err != nil {
		return nil, err
	}

	if node != nil && Contains(node.Labels, RESOURCE_ORG_LABEL) {
		data := node.Data["data"].(string)
		admin_rels, _ := storage.RelationshipsNode(node.Id, "admin")
		app_rels, _ := storage.RelationshipsNode(node.Id, "app")
		token_rels, _ := storage.RelationshipsNode(node.Id, "token")

		if node.Data["name"] == nil {
			node.Data["name"] = ""
		}

		return &Org{strconv.Itoa(node.Id), node.Data["name"].(string), data, len(token_rels), len(admin_rels), len(app_rels)}, nil
	}

	return nil, errors.New("OrgId `"+id+"` not exist")
}

func AddOrg(org *Org) (err error) {
	properties := graph.Props{
		"name": org.Name,
		"data": org.Data,
	}

	_org, err := storage.NewNode(properties, RESOURCE_ORG_LABEL)

	if err != nil {
		return err
	}

	org.Id = strconv.Itoa(_org.Id)

	return nil
}

func UpdateOrg(org *Org) (err error) {
	_id, _ := strconv.Atoi(org.Id)
	return storage.SetPropertyNode(_id, "name", org.Name)
}

func DeleteOrg(id string) (error) {
	_id, _ := strconv.Atoi(id)
	return storage.DeleteNode(_id)
}

// TOKEN PART

func GetOrgTokenList(orgId string) (orgList []*Token, err error) {
	var tokens []*Token

	if tokens == nil {
		tokens = make([]*Token, 0)
	}

	return tokens, nil
}

func GetOrgToken(id string, orgId string) (org *Token, err error) {
	return nil, errors.New("TokenId not exist")
}

func AddOrgToken(token Token, orgId string) (id string, err error) {
	// get org
	org, err := GetOrg(orgId)
	if err != nil {
		return "0", err
	}

	// add token
	properties := graph.Props{
		"data": token.Data,
	}
	_token, err := storage.NewNode(properties, RESOURCE_TOKEN_LABEL)

	if err != nil {
		return "0", err
	}

	// create relation
	_adminId, _ := strconv.Atoi(org.Id)
	rel, err := storage.RelateNodes(_adminId, _token.Id, "token", nil)

	if err != nil || rel.Type == "" {
		return "0", err
	}

	token.Id = strconv.Itoa(_token.Id)

	return token.Id, nil
}
