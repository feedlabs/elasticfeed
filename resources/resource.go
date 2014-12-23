package resources

import (
	"errors"

	"github.com/feedlabs/feedify/service"
	"github.com/feedlabs/feedify/graph"
	"github.com/feedlabs/feedify/stream"
)

const RESOURCE_ADMIN_LABEL = "admin"
const RESOURCE_TOKEN_LABEL = "token"
const RESOURCE_APPLICATION_LABEL = "application"
const RESOURCE_FEED_LABEL = "feed"
const RESOURCE_ENTRY_LABEL = "entry"

var (
	Admins          map[string]*Admin
	Tokens          map[string]*Token
	Applications    map[string]*Application
	Feeds           map[string]*Feed
	Entries         map[string]*Entry

	message    *stream.StreamMessage
	storage    *graph.GraphStorage
)

type Admin struct {
	Id           string
	Data         string
	Applications int
}

type Token struct {
	Id             string
	Data           string
}

type Application struct {
	Id        string
	Data      string
	Feeds     int
}

type Feed struct {
	Id      string
	Data    string
	Entries int
}

type Entry struct {
	Id        string
	FeedId    string
	Data      string
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

func contains(s []string, e string) bool {
	for _, a := range s { if a == e { return true } }
	return false
}
