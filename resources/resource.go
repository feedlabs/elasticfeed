package resources

import (
	"errors"

	"github.com/feedlabs/feedify/service"
)

var (
	Feeds map[string]*Feed

	stream	*service.StreamService
	graph	*service.GraphService
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
	stream, _ = service.NewStream()
	if stream == nil {
		panic(errors.New("Cannot create stream service"))
	}
	graph, _ = service.NewGraph()
	if graph == nil {
		panic(errors.New("Cannot create graph service"))
	}
}
