package adapter

import (
	"log"
)

type Neo4j struct {
	host	string
	port	string
}

func (m Neo4j) Connect() {
	log.Printf("%T connected", m)
}

func NewNeo4j() *Neo4j {
	return &Neo4j{}
}
