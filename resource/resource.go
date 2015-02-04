package resource

import (
	"errors"
	"encoding/json"

	"github.com/feedlabs/feedify/service"
	"github.com/feedlabs/feedify/graph"
	"github.com/feedlabs/feedify/stream"

	"github.com/feedlabs/elasticfeed/service/stream/controller/room"
	"github.com/feedlabs/elasticfeed/service/stream/model"
)

const RESOURCE_ORG_LABEL = "org"
const RESOURCE_ADMIN_LABEL = "admin"
const RESOURCE_TOKEN_LABEL = "token"
const RESOURCE_APPLICATION_LABEL = "application"
const RESOURCE_FEED_LABEL = "feed"
const RESOURCE_ENTRY_LABEL = "entry"

var (
	Orgs            map[string]*Org
	Admins        map[string]*Admin
	Tokens        map[string]*Token
	Applications    map[string]*Application
	Feeds            map[string]*Feed
	Entries        map[string]*Entry

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

	Entries       int
}

type Entry struct {
	Id        string
	Feed      *Feed
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

	go ResourceStreamManager()
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

	list, err := GetEntryList(socketEvent.FeedId, socketEvent.AppId, socketEvent.OrgId)

	if err == nil {

		// *********************************************************************
		// register socket handler
		// needs to send notiffication to long pooling + ws
		// join should generate uniqe ID and client should use it
		// maybe sessionID could be as uniqeID ?
		// room.FeedSubscribers[socketEvent.FeedId][channelID] = socketEvent
		// *********************************************************************

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
