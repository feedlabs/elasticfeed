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
	Data interface{}
}

type PipelinePrepareResponse struct {
	Warnings []string
	Error    *BasicError
	Data     interface {}
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

func (h *RpcPipeline) Run(data interface{}) (interface{}, error) {

	args := PipelinePrepareArgs{
		Data:     data,
	}

	var response PipelinePrepareResponse

	_ = h.client.Call("Pipeline.Run", &args, &response)

	return response.Data, nil
}

func (b *RpcPipeline) Cancel() {
	if err := b.client.Call("Pipeline.Cancel", new(interface{}), new(interface{})); err != nil {
		log.Printf("Error cancelling pipeline: %s", err)
	}
}

func (b *PipelineRpcServer) Prepare(args *PipelinePrepareArgs, reply *PipelinePrepareResponse) error {
	warnings, err := b.pipeline.Prepare(nil)
	*reply = PipelinePrepareResponse{
		Warnings: warnings,
		Error:    NewBasicError(err),
	}
	return nil
}

func (b *PipelineRpcServer) Run(args *PipelinePrepareArgs, reply *PipelinePrepareResponse) (err error) {

	data, err := b.pipeline.Run(args.Data)

	*reply = PipelinePrepareResponse{
		Warnings: nil,
		Error: nil,
		Data: data,
	}

	return nil
}

func (b *PipelineRpcServer) Cancel(args *interface{}, reply *interface{}) error {
	b.pipeline.Cancel()
	return nil
}
