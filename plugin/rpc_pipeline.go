package plugin

import (
	"log"
	"net/rpc"

	"github.com/feedlabs/elasticfeed/plugin/model"
)

// An implementation of Pipeline where the builder is actually executed
// over an RPC connection.
type RpcPipeline struct {
	client *rpc.Client
	mux    *muxBroker
}

// PipelineRpcServer wraps a Pipeline implementation and makes it exportable
// as part of a Golang RPC server.
type PipelineRpcServer struct {
	pipeline model.Pipeline
	mux     *muxBroker
}

type PipelinePrepareArgs struct {
	Configs []interface{}
}

type PipelinePrepareResponse struct {
	Warnings []string
	Error    *BasicError
}

func (b *RpcPipeline) Prepare(config ...interface{}) ([]string, error) {
	var resp PipelinePrepareResponse
	cerr := b.client.Call("Pipeline.Prepare", &PipelinePrepareArgs{config}, &resp)
	if cerr != nil {
		return nil, cerr
	}
	var err error = nil
	if resp.Error != nil {
		err = resp.Error
	}

	return resp.Warnings, err
}

func (b *RpcPipeline) Run(cache model.Cache) (model.Artifact, error) {
	nextId := b.mux.NextId()
	server := newRpcServerWithMux(b.mux, nextId)
	server.RegisterCache(cache)
	go server.Serve()

	var responseId uint32
	if err := b.client.Call("Pipeline.Run", nextId, &responseId); err != nil {
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

func (b *RpcPipeline) Cancel() {
	if err := b.client.Call("Pipeline.Cancel", new(interface{}), new(interface{})); err != nil {
		log.Printf("Error cancelling pipeline: %s", err)
	}
}

func (b *PipelineRpcServer) Prepare(args *PipelinePrepareArgs, reply *PipelinePrepareResponse) error {
	warnings, err := b.pipeline.Prepare(args.Configs...)
	*reply = PipelinePrepareResponse{
		Warnings: warnings,
		Error:    NewBasicError(err),
	}
	return nil
}

func (b *PipelineRpcServer) Run(streamId uint32, reply *uint32) error {
	client, err := newRpcClientWithMux(b.mux, streamId)
	if err != nil {
		return NewBasicError(err)
	}
	defer client.Close()

	artifact, err := b.pipeline.Run(client.Cache())
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

func (b *PipelineRpcServer) Cancel(args *interface{}, reply *interface{}) error {
	b.pipeline.Cancel()
	return nil
}
