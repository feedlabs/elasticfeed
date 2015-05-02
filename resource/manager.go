package resource

import (
	"encoding/json"
	"time"
	"math/rand"

	emodel "github.com/feedlabs/elasticfeed/elasticfeed/model"
	smodel "github.com/feedlabs/elasticfeed/service/stream/model"
	pmodel "github.com/feedlabs/elasticfeed/plugin/model"

	"github.com/feedlabs/elasticfeed/service/stream/controller/room"
)

var (
	pluginManagerAnn pmodel.Pipeline
	entryListCache []*Entry
)

type ResourceManager struct {
	engine emodel.Elasticfeed
}

func (this * ResourceManager) Init() {
	this.InitStreamCommunicator()
}

func (this * ResourceManager) InitStreamCommunicator() {
	// should bind service-stream-controllers to get handler to channel
	// should pass it down to listen for events on streaming controllers
	go this.ResourceStreamManager()
}

func (this * ResourceManager) ResourceStreamManager() {
	for {
		select {
		case socketEvent := <-this.engine.GetServiceManager().GetStreamService().GetFeedRoomManager().ResourceEvent:

			go this.ResourceStreamRequest(socketEvent)
		}
	}
}

func (this * ResourceManager) ResourceStreamRequest(socketEvent smodel.SocketEvent) {

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

	// COLLECTING ENTRIES
	if entryListCache == nil {
		entryListCache, _ = GetEntryList(socketEvent.FeedId, socketEvent.AppId, socketEvent.OrgId)
	}

	// WORKFLOW TIMEOUT
	// !! SHOULD BE CONFIGURABLE OVER RUNTIME SETTING
	// !! DEFAULT VALUE SHOULD BE IN CONFIG FILE
	go func() {
		amt := time.Duration(100)
		time.Sleep(amt * time.Millisecond)
		timeout <- true
	}()

	// WORKFLOW PIPELINE
	go func(list []*Entry, socketEvent smodel.SocketEvent) {

		if pluginManagerAnn == nil {
			pluginManagerAnn, _ = this.engine.GetPluginManager().LoadPipeline("ann")
			pluginManagerAnn.Prepare()
		}

		newList, _ := pluginManagerAnn.Run(list)

		var newEntryList []*Entry

		for _, v := range newList.([]interface{}) {
			Id := ""
			Data := ""
			for k, vv := range v.(map[interface{}]interface{}) {
				if k == "Id" {
					Id = vv.(string)
				}
				if k == "Data" {
					Data = vv.(string)
				}
			}
			if Id != "" && Data != "" {
				newEntryList = append(newEntryList, &Entry{Id, nil, Data})
			}
		}

		list = newEntryList

		results <- list
	}(entryListCache, socketEvent)

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

func NewResourceManager(engine emodel.Elasticfeed) emodel.ResourceManager {
	return &ResourceManager{engine}
}
