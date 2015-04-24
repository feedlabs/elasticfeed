package plugin

import (
	"fmt"
	"github.com/hashicorp/go-msgpack/codec"
	"io"
	"log"
	"net/rpc"
	"sync/atomic"

	"github.com/feedlabs/elasticfeed/plugin/model"
)

var endpointId uint64

const (
	DefaultIndexerEndpoint      string = "Indexer"
	DefaultPipelineEndpoint     string = "Pipeline"
	DefaultArtifactEndpoint     string = "Artifact"
)

// RpcServer represents an RPC server for Packer. This must be paired on
// the other side with a Client.
type RpcServer struct {
	mux      *muxBroker
	streamId uint32
	server   *rpc.Server
	closeMux bool
}

// NewRpcServer returns a new Packer RPC server.
func NewRpcServer(conn io.ReadWriteCloser) *RpcServer {
	mux, _ := newMuxBrokerServer(conn)
	result := newRpcServerWithMux(mux, 0)
	result.closeMux = true
	go mux.Run()
	return result
}

func newRpcServerWithMux(mux *muxBroker, streamId uint32) *RpcServer {
	return &RpcServer{
		mux:      mux,
		streamId: streamId,
		server:   rpc.NewServer(),
		closeMux: false,
	}
}

func (s *RpcServer) Close() error {
	if s.closeMux {
		log.Printf("[WARN] Shutting down mux conn in RpcServer")
		return s.mux.Close()
	}

	return nil
}

func (s *RpcServer) RegisterCache(c model.Cache) {
	//	s.server.RegisterName(DefaultCacheEndpoint, &CacheServer{
	//		cache: c,
	//	})
}

func (s *RpcServer) RegisterArtifact(a model.Artifact) {
	s.server.RegisterName(DefaultArtifactEndpoint, &ArtifactRpcServer{
		artifact: a,
	})
}

func (s *RpcServer) RegisterIndexer(i model.Indexer) {
	s.server.RegisterName(DefaultIndexerEndpoint, &IndexerRpcServer{
		indexer: i,
		mux:     s.mux,
	})
}

func (s *RpcServer) RegisterPipeline(p model.Pipeline) {
	s.server.RegisterName(DefaultPipelineEndpoint, &PipelineRpcServer{
		pipeline: p,
		mux:     s.mux,
	})
}

// ServeConn serves a single connection over the RPC server. It is up
// to the caller to obtain a proper io.ReadWriteCloser.
func (s *RpcServer) Serve() {
	// Accept a connection on stream ID 0, which is always used for
	// normal client to server connections.
	stream, err := s.mux.Accept(s.streamId)
	if err != nil {
		log.Printf("[ERR] Error retrieving stream for serving: %s", err)
		return
	}
	defer stream.Close()

	h := &codec.MsgpackHandle{
		RawToString: true,
		WriteExt:    true,
	}
	rpcCodec := codec.GoRpc.ServerCodec(stream, h)
	s.server.ServeCodec(rpcCodec)
}

// registerComponent registers a single Packer RPC component onto
// the RPC server. If id is true, then a unique ID number will be appended
// onto the end of the endpoint.
//
// The endpoint name is returned.
func registerComponent(server *rpc.Server, name string, rcvr interface{}, id bool) string {
	endpoint := name
	if id {
		fmt.Sprintf("%s.%d", endpoint, atomic.AddUint64(&endpointId, 1))
	}

	server.RegisterName(endpoint, rcvr)
	return endpoint
}
