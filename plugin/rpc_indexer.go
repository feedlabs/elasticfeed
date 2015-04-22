package plugin

import (
	"log"
	"net/rpc"

	"github.com/feedlabs/elasticfeed/plugin/model"
)

// An implementation of Indexer where the builder is actually executed
// over an RPC connection.
type RpcIndexer struct {
	client *rpc.Client
	mux    *muxBroker
}

// IndexerRpcServer wraps a Indexer implementation and makes it exportable
// as part of a Golang RPC server.
type IndexerRpcServer struct {
	indexer model.Indexer
	mux     *muxBroker
}

type IndexerPrepareArgs struct {
	Configs []interface{}
}

type IndexerPrepareResponse struct {
	Warnings []string
	Error    *BasicError
}

func (b *RpcIndexer) Prepare(config ...interface{}) ([]string, error) {
	var resp IndexerPrepareResponse
	cerr := b.client.Call("Indexer.Prepare", &IndexerPrepareArgs{config}, &resp)
	if cerr != nil {
		return nil, cerr
	}
	var err error = nil
	if resp.Error != nil {
		err = resp.Error
	}

	return resp.Warnings, err
}

func (b *RpcIndexer) Run(cache model.Cache) (model.Artifact, error) {
	nextId := b.mux.NextId()
	server := newRpcServerWithMux(b.mux, nextId)
	server.RegisterCache(cache)
	go server.Serve()

	var responseId uint32
	if err := b.client.Call("Indexer.Run", nextId, &responseId); err != nil {
		return nil, err
	}

	if responseId == 0 {
		return nil, nil
	}

	client, err := newRpcClientWithMux(b.mux, responseId)
	if err != nil {
		return nil, err
	}

	return client.Artifact(), nil
}

func (b *RpcIndexer) Cancel() {
	if err := b.client.Call("Indexer.Cancel", new(interface{}), new(interface{})); err != nil {
		log.Printf("Error cancelling indexer: %s", err)
	}
}

func (b *IndexerRpcServer) Prepare(args *IndexerPrepareArgs, reply *IndexerPrepareResponse) error {
	warnings, err := b.indexer.Prepare(args.Configs...)
	*reply = IndexerPrepareResponse{
		Warnings: warnings,
		Error:    NewBasicError(err),
	}
	return nil
}

func (b *IndexerRpcServer) Run(streamId uint32, reply *uint32) error {
	client, err := newRpcClientWithMux(b.mux, streamId)
	if err != nil {
		return NewBasicError(err)
	}
	defer client.Close()

	artifact, err := b.indexer.Run(client.Cache())
	if err != nil {
		return NewBasicError(err)
	}

	*reply = 0
	if artifact != nil {
		streamId = b.mux.NextId()
		server := newRpcServerWithMux(b.mux, streamId)
		server.RegisterArtifact(artifact)
		go server.Serve()
		*reply = streamId
	}

	return nil
}

func (b *IndexerRpcServer) Cancel(args *interface{}, reply *interface{}) error {
	b.indexer.Cancel()
	return nil
}
