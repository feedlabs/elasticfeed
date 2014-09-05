package memcache

import (
	"log"
)

type MemcacheClient struct {
	host string
	port string
}

func (m MemcacheClient) Connect() {
	log.Printf("%T connected", m)
}

func NewMemcacheClient() *MemcacheClient {
	return &MemcacheClient{}
}
