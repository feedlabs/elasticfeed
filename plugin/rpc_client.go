package plugin

import (
	"github.com/hashicorp/go-msgpack/codec"
	"io"
	"log"
	"net/rpc"

	"github.com/feedlabs/elasticfeed/plugin/model"
)

// RpcClient is the client end that communicates with a Packer RPC server.
// Establishing a connection is up to the user, the RpcClient can just
// communicate over any ReadWriteCloser.
type RpcClient struct {
	mux      *muxBroker
	client   *rpc.Client
	closeMux bool
}

func NewRpcClient(rwc io.ReadWriteCloser) (*RpcClient, error) {
	mux, err := newMuxBrokerClient(rwc)
	if err != nil {
		return nil, err
	}
	go mux.Run()

	result, err := newRpcClientWithMux(mux, 0)
	if err != nil {
		mux.Close()
		return nil, err
	}

	result.closeMux = true
	return result, err
}

func newRpcClientWithMux(mux *muxBroker, streamId uint32) (*RpcClient, error) {
	clientConn, err := mux.Dial(streamId)
	if err != nil {
		return nil, err
	}

	h := &codec.MsgpackHandle{
		RawToString: true,
		WriteExt:    true,
	}
	clientCodec := codec.GoRpc.ClientCodec(clientConn, h)

	return &RpcClient{
		mux:      mux,
		client:   rpc.NewClientWithCodec(clientCodec),
		closeMux: false,
	}, nil
}

func (c *RpcClient) Close() error {
	if err := c.client.Close(); err != nil {
		return err
	}

	if c.closeMux {
		log.Printf("[WARN] RpcClient is closing mux")
		return c.mux.Close()
	}

	return nil
}

func (c *RpcClient) Indexer() model.Indexer {
	return &RpcIndexer{
		client: c.client,
		mux:    c.mux,
	}
}

func (c *RpcClient) Pipeline() model.Pipeline {
	return &RpcPipeline{
		client: c.client,
		mux:    c.mux,
	}
}

func (c *RpcClient) Artifact() model.Artifact {
	return &RpcArtifact{
		client:   c.client,
		endpoint: DefaultArtifactEndpoint,
	}
}

func (c *RpcClient) Cache() model.Cache {
	return &RpcCache{
		client: c.client,
	}
}
