package db

import (
	"log"
)

type Memcache struct {
	host	string
	port	string
}

func (m Memcache) Connect() {
	log.Printf("%T connected", m)
}

func NewMemcache() *Memcache {
	return &Memcache{}
}
