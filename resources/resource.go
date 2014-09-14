package resources

import (
	"errors"

	"github.com/feedlabs/feedify/service"
	"github.com/feedlabs/feedify/graph"
	"github.com/feedlabs/feedify/stream"
)

var (
	Feeds map[string]*Feed

	message	*stream.StreamMessage
	storage	*graph.GraphStorage
)

type Feed struct {
	Id      string
	Data    string
	Entries map[string]*FeedEntry
}

type FeedEntry struct {
	Id   string
	Data string
}

func init() {
	stream_service, _ := service.NewStream()
	if stream_service == nil {
		panic(errors.New("Cannot create stream service"))
	}
	message = stream_service.Message

	graph_service, _ := service.NewGraph()
	if graph_service == nil {
		panic(errors.New("Cannot create graph service"))
	}
	storage = graph_service.Storage
}
