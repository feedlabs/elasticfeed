package plugin

import (
	"net/rpc"
	"github.com/feedlabs/elasticfeed/plugin/model"
)

// An implementation of packer.Artifact where the RpcArtifact is actually
// available over an RPC connection.
type RpcArtifact struct {
	client   *rpc.Client
	endpoint string
}

// ArtifactRpcServer wraps a packer.Artifact implementation and makes it
// exportable as part of a Golang RPC server.
type ArtifactRpcServer struct {
	artifact model.Artifact
}

func (a *RpcArtifact) BuilderId() (result string) {
	a.client.Call(a.endpoint+".BuilderId", new(interface{}), &result)
	return
}

func (a *RpcArtifact) Files() (result []string) {
	a.client.Call(a.endpoint+".Files", new(interface{}), &result)
	return
}

func (a *RpcArtifact) Id() (result string) {
	a.client.Call(a.endpoint+".Id", new(interface{}), &result)
	return
}

func (a *RpcArtifact) String() (result string) {
	a.client.Call(a.endpoint+".String", new(interface{}), &result)
	return
}

func (a *RpcArtifact) State(name string) (result interface{}) {
	a.client.Call(a.endpoint+".State", name, &result)
	return
}

func (a *RpcArtifact) Destroy() error {
	var result error
	if err := a.client.Call(a.endpoint+".Destroy", new(interface{}), &result); err != nil {
		return err
	}

	return result
}

func (s *ArtifactRpcServer) BuilderId(args *interface{}, reply *string) error {
	*reply = s.artifact.BuilderId()
	return nil
}

func (s *ArtifactRpcServer) Files(args *interface{}, reply *[]string) error {
	*reply = s.artifact.Files()
	return nil
}

func (s *ArtifactRpcServer) Id(args *interface{}, reply *string) error {
	*reply = s.artifact.Id()
	return nil
}

func (s *ArtifactRpcServer) String(args *interface{}, reply *string) error {
	*reply = s.artifact.String()
	return nil
}

func (s *ArtifactRpcServer) State(name string, reply *interface{}) error {
	*reply = s.artifact.State(name)
	return nil
}

func (s *ArtifactRpcServer) Destroy(args *interface{}, reply *error) error {
	err := s.artifact.Destroy()
	if err != nil {
		err = NewBasicError(err)
	}

	*reply = err
	return nil
}
