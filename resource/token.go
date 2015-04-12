package resource

import (
	"errors"
	"strconv"

	"github.com/feedlabs/feedify/graph"
)

func GetTokenList(AdminId string, OrgId string) (tokenLinst []*Token, err error) {
	admin, err := GetAdmin(AdminId, OrgId)
	if err != nil {
		return nil, err
	}

	_id, _ := strconv.Atoi(admin.Id)
	_rels, _ := storage.RelationshipsNode(_id, "token")

	var tokens []*Token

	for _, rel := range _rels {
		data := rel.EndNode.Data["data"].(string)
		id := strconv.Itoa(rel.EndNode.Id)

		token := &Token{id, admin, data}
		tokens = append(tokens, token)
	}

	if tokens == nil {
		tokens = make([]*Token, 0)
	}

	return tokens, nil
}

func GetToken(id string, AdminId string, OrgId string) (token *Token, err error) {
	admin, err := GetAdmin(AdminId, OrgId)
	if err != nil {
		return nil, err
	}

	_id, err := strconv.Atoi(id)
	node, err := storage.Node(_id)

	if err != nil {
		return nil, err
	}

	if node != nil && Contains(node.Labels, RESOURCE_TOKEN_LABEL) {
		data := node.Data["data"].(string)
		return &Token{strconv.Itoa(node.Id), admin, data}, nil
	}

	return nil, errors.New("TokenId not exist")
}

func AddTokenForOrganisation(token Token, orgId string) (id string, err error) {
	return "0", nil
}

func AddToken(token Token, adminId string, orgId string) (id string, err error) {
	// get admin
	admin, err := GetAdmin(adminId, orgId)
	if err != nil {
		return "0", err
	}

	// add token
	properties := graph.Props{"data": token.Data}
	_token, err := storage.NewNode(properties, RESOURCE_TOKEN_LABEL)

	if err != nil {
		return "0", err
	}

	// create relation
	_adminId, _ := strconv.Atoi(admin.Id)
	rel, err := storage.RelateNodes(_adminId, _token.Id, "token", nil)

	if err != nil || rel.Type == "" {
		return "0", err
	}

	token.Id = strconv.Itoa(_token.Id)

	return token.Id, nil
}

func UpdateToken(id string, data string) (err error) {
	_id, _ := strconv.Atoi(id)
	return storage.SetPropertyNode(_id, "data", data)
}

func DeleteToken(id string) (error) {
	_id, _ := strconv.Atoi(id)
	return storage.DeleteNode(_id)
}
