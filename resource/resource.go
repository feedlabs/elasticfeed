package resource

import (
	"errors"
	"encoding/json"
	"time"
	"math/rand"

	"github.com/feedlabs/feedify/service"
	"github.com/feedlabs/feedify/graph"
	"github.com/feedlabs/feedify/stream"

	"github.com/feedlabs/elasticfeed/service/stream/controller/room"
	"github.com/feedlabs/elasticfeed/service/stream/model"
	"github.com/feedlabs/elasticfeed/plugin/plugins/pipeline"
)

const (
	RESOURCE_ORG_LABEL         = "org"
	RESOURCE_ADMIN_LABEL       = "admin"
	RESOURCE_TOKEN_LABEL       = "token"
	RESOURCE_APPLICATION_LABEL = "application"
	RESOURCE_FEED_LABEL        = "feed"
	RESOURCE_ENTRY_LABEL       = "entry"
	RESOURCE_METRIC_LABEL      = "metric"
	RESOURCE_VIEWER_LABEL      = "viewer"
	RESOURCE_WORKFLOW_LABEL    = "workflow"
	RESOURCE_PLUGIN_LABEL      = "plugin"
)

var (
	Orgs            map[string]*Org
	Admins            map[string]*Admin
	Tokens            map[string]*Token
	Applications    map[string]*Application
	Feeds           map[string]*Feed
	Entries            map[string]*Entry
	Metrics            map[string]*Metric
	Viewers            map[string]*Viewer
	Workflows            map[string]*Workflow
	Plugins            map[string]*Plugin

	message    *stream.StreamMessage
	storage    *graph.GraphStorage
)

type Org struct {
	Id               string
	Name             string
	Data             string

	Tokens            int
	Admins            int
	Applications      int
}

type Admin struct {
	Id               string
	Org              *Org

	Username              string
	Maintainer            bool
	Whitelist             []string
	Data                  string

	Tokens                int
}

type Token struct {
	Id                  string
	Admin                *Admin
	Data                string
}

type Application struct {
	Id        string
	Org       *Org
	Data      string
	Feeds     int
}

type Feed struct {
	Id            string
	Application   *Application
	Data          string

	Entries         int
	Workflows       int
}

type Entry struct {
	Id        string
	Feed      *Feed
	Data      string
}

type Viewer struct {}

type Metric struct {}

type Workflow struct {
	Id             string
	Feed           *Feed
	Default        bool
	Data           string
}

type Plugin struct {
	Id            string
	Name          string
	Group         string
	Version       string
	Path          string
	License       string
}

func ResourceStreamManager() {
	for {
		select {
		case socketEvent := <-room.ResourceEvent:

			go ResourceStreamRequest(socketEvent)
		}
	}
}

func ResourceStreamRequest(socketEvent model.SocketEvent) {

	// *******************************************************************
	// here should be implemented REAL CONTENT IMPROVEMENT
	// based on connected user (viewer) or users (audience)!: habits, behaviours, stats etc.
	// PIPE: filtering, customization
	// SCENARIO-ENGINE: scenarios
	// *******************************************************************

	// *******************************************************************
	// SCENARIO AND RULES/METRICS
	// should use go routine with time limit to query filter rules
	// if in specific time there is no rules the results should be sent
	// client feed. After this the next package should be sent with
	// rules which entries should be remove/hidden from the view!
	// *******************************************************************

	timeout := make(chan bool, 1)
	results := make(chan []*Entry, 1)

	list, _ := GetEntryList(socketEvent.FeedId, socketEvent.AppId, socketEvent.OrgId)

	// PIPE TIMEOUT
	go func() {
		amt := time.Duration(rand.Intn(100))
		time.Sleep(amt * time.Millisecond)
		timeout <- true
	}()

	// SHOULD BE A FILTER IMPLEMENTATION
	go func(list []*Entry, socketEvent model.SocketEvent) {
		list = pipeline.Filter(list).([]*Entry)
		results <- list
	}(list, socketEvent)

	select {

		// IF PIPE TAKES TOO MUCH TIME, DATA DELAYED
	case <-timeout:

		event := room.NewFeedEvent(room.FEED_ENTRY_NEW, socketEvent.FeedId, "{Content:\"tiemout\"}")
		data, _ := json.Marshal(event)

		if socketEvent.Ws != nil {
			amt := time.Duration(rand.Intn(500)) * 1000
			time.Sleep(amt * time.Microsecond)
			socketEvent.Ws.WriteMessage(1, data)
		}

		if socketEvent.Ch != nil {
			socketEvent.Ch <- data
		}

		// IF DATA ARRIVES WITHOUT DELAY
	case list := <-results:

		// *********************************************************************
		// register socket handler
		// needs to send notiffication to long pooling + ws
		// join should generate uniqe ID and client should use it
		// maybe sessionID could be as uniqeID ?
		// room.FeedSubscribers[socketEvent.FeedId][channelID] = socketEvent
		// *********************************************************************
		//
		// timeout with channels and routines?
		// http://blog.golang.org/go-concurrency-patterns-timing-out-and

		//		amt := time.Duration(rand.Intn(500)) * 10000
		//		time.Sleep(amt * time.Microsecond)

		d, _ := json.Marshal(list)
		event := room.NewFeedEvent(room.FEED_ENTRY_INIT, socketEvent.FeedId, string(d))
		data, _ := json.Marshal(event)

		if socketEvent.Ws != nil {
			socketEvent.Ws.WriteMessage(1, data)
		}

		if socketEvent.Ch != nil {
			socketEvent.Ch <- data
		}

	}

}

func Contains(s []string, e string) bool {
	for _, a := range s { if a == e { return true } }
	return false
}

func ConvertInterfaceToStringArray(d interface{}) []string {
	data := d.([]interface{})
	output := make([]string, len(data))
	for i := 0; i < len(data); i++ {
		output[i] = data[i].(string)
	}
	return output
}

func InitResources() {
	Admins = make(map[string]*Admin)
	Applications = make(map[string]*Application)
	Feeds = make(map[string]*Feed)
	Entries = make(map[string]*Entry)
	Orgs = make(map[string]*Org)
	Tokens = make(map[string]*Token)
	Metrics = make(map[string]*Metric)
	Viewers = make(map[string]*Viewer)
	Plugins = make(map[string]*Plugin)
}

func InitStorage() {
	graph_service, _ := service.NewGraph()
	if graph_service == nil {
		panic(errors.New("Cannot create graph service"))
	}
	storage = graph_service.Storage
}

func InitStreamCommunicator() {
	go ResourceStreamManager()
}
