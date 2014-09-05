package adapter

import (
	"log"
)

type Neo4jAdapter struct {
	host string
	port string
}

func (m Neo4jAdapter) Connect() {
	log.Printf("%T connected", m)
}

func NewNeo4jAdapter() *Neo4jAdapter {
	return &Neo4jAdapter{}
}
