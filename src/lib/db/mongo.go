package db

import (
	"log"
)

type Mongo struct {
	host	string
	port	string
}

func (m Mongo) Connect() {
	log.Printf("%T connected", m)
}

func NewMongo() *Mongo {
	return &Mongo{}
}
