package service

import (
	"github.com/feedlabs/feedify/lib/feedify/memcache"
)

func NewMemcache() *memcache.MemcacheClient {
	return memcache.NewMemcacheClient()
}
