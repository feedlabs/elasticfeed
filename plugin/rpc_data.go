package plugin

import (
	"net/rpc"
	"github.com/feedlabs/elasticfeed/plugin/model"
)

// An implementation of packer.Cache where the RpcData is actually executed
// over an RPC connection.
type RpcData struct {
	client *rpc.Client
	endpoint string
}

// CacheRpcServer wraps a packer.Cache implementation and makes it exportable
// as part of a Golang RPC server.
type DataRpcServer struct {
	data model.Data
}
