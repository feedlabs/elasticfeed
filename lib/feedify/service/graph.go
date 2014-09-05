package service

import (
	"github.com/feedlabs/feedify/lib/feedify/graph/adapter"
)

func NewCayley() *adapter.CayleyAdapter {
	return adapter.NewCayleyAdapter()
}

func NewNeo4j() *adapter.Neo4jAdapter {
	return adapter.NewNeo4jAdapter()
}
