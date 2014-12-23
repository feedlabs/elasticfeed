package resources

import (
	"errors"
	"strconv"

	"github.com/feedlabs/feedify/graph"
)

func init() {
	Tokens = make(map[string]*Token)
}

func GetTokenList() []*Token {
	nodes, err := storage.FindNodesByLabel(RESOURCE_TOKEN_LABEL)
	if err != nil {
		nodes = nil
	}

	var tokens []*Token

	for _, node := range nodes {
		data := node.Data["data"].(string)
		id := strconv.Itoa(node.Id)

		token := &Token{id , data}
		tokens = append(tokens, token)
	}

	return tokens
}

func GetToken(id string) (token *Token, err error) {
	_id, err := strconv.Atoi(id)
	node, err := storage.Node(_id)

	if err != nil {
		return nil, err
	}

	if node != nil && contains(node.Labels, RESOURCE_TOKEN_LABEL) {
		data := node.Data["data"].(string)
		return &Token{strconv.Itoa(node.Id), data}, nil
	}

	return nil, errors.New("TokenId not exist")
}

func AddToken(token Token) (id string, err error) {
	properties := graph.Props{"data": token.Data}
	_token, err := storage.NewNode(properties, RESOURCE_TOKEN_LABEL)

	if err != nil {
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
