package plugin

import (
	"log"
	"net/rpc"
	"github.com/feedlabs/elasticfeed/plugin/model"
)

// An implementation of packer.Cache where the RpcCache is actually executed
// over an RPC connection.
type RpcCache struct {
	client *rpc.Client
}

// CacheRpcServer wraps a packer.Cache implementation and makes it exportable
// as part of a Golang RPC server.
type CacheRpcServer struct {
	cache model.Cache
}

type CacheRLockResponse struct {
	Path   string
	Exists bool
}

func (c *RpcCache) Lock(key string) (result string) {
	if err := c.client.Call("Cache.Lock", key, &result); err != nil {
		log.Printf("[ERR] Cache.Lock error: %s", err)
		return
	}

	return
}

func (c *RpcCache) RLock(key string) (string, bool) {
	var result CacheRLockResponse
	if err := c.client.Call("Cache.RLock", key, &result); err != nil {
		log.Printf("[ERR] Cache.RLock error: %s", err)
		return "", false
	}

	return result.Path, result.Exists
}

func (c *RpcCache) Unlock(key string) {
	if err := c.client.Call("Cache.Unlock", key, new(interface{})); err != nil {
		log.Printf("[ERR] Cache.Unlock error: %s", err)
		return
	}
}

func (c *RpcCache) RUnlock(key string) {
	if err := c.client.Call("Cache.RUnlock", key, new(interface{})); err != nil {
		log.Printf("[ERR] Cache.RUnlock error: %s", err)
		return
	}
}

func (c *CacheRpcServer) Lock(key string, result *string) error {
	*result = c.cache.Lock(key)
	return nil
}

func (c *CacheRpcServer) Unlock(key string, result *interface{}) error {
	c.cache.Unlock(key)
	return nil
}

func (c *CacheRpcServer) RLock(key string, result *CacheRLockResponse) error {
	path, exists := c.cache.RLock(key)
	*result = CacheRLockResponse{path, exists}
	return nil
}

func (c *CacheRpcServer) RUnlock(key string, result *interface{}) error {
	c.cache.RUnlock(key)
	return nil
}
