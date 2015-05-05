package resource

import (
	"errors"
	"strconv"
	"encoding/json"

	"github.com/feedlabs/feedify/graph"
	"github.com/feedlabs/elasticfeed/service/stream/controller/room"
)

func GetEntryList(FeedId string, ApplicationId string, OrgId string) (feedEntries []*Entry, err error) {
	feed, err := GetFeed(FeedId, ApplicationId, OrgId)

	if err != nil {
		return nil, err
	}

	_id, _ := strconv.Atoi(feed.Id)
	_rels, _ := storage.RelationshipsNode(_id, "entry")

	var entries []*Entry

	for _, rel := range _rels {
		data := rel.EndNode.Data["data"].(string)
		entry := &Entry{strconv.Itoa(rel.EndNode.Id), feed, data}
		if entry != nil && Contains(rel.EndNode.Labels, RESOURCE_ENTRY_LABEL) && feed.Id == rel.EndNode.Data["feedId"].(string) {
			entries = append(entries, entry)
		}
	}

	if entries == nil {
		entries = make([]*Entry, 0)
	}

	return entries, nil
}

func GetEntry(id string, FeedId string, ApplicationId string, OrgId string) (feedEntry *Entry, err error) {
	feed, err := GetFeed(FeedId, ApplicationId, OrgId)
	if err != nil {
		return nil, err
	}

	_id, err := strconv.Atoi(id)
	entry, err := storage.Node(_id)

	if err != nil {
		return nil, err
	}

	if entry != nil && Contains(entry.Labels, RESOURCE_ENTRY_LABEL) && feed.Id == entry.Data["feedId"].(string) {
		data := entry.Data["data"].(string)
		return &Entry{strconv.Itoa(entry.Id), feed, data}, nil
	}

	return nil, errors.New("EntryId `"+id+"` not exist")
}

func AddEntry(feedEntry Entry, FeedId string, ApplicationId string, OrgId string) (EntryId string, err error) {
	// get feed
	feed, err := GetFeed(FeedId, ApplicationId, OrgId)
	if err != nil {
		return "0", err
	}

	// add feed-entry
	properties := graph.Props{
		"feedId": feed.Id,
		"data": feedEntry.Data,
	}
	entry, err := storage.NewNode(properties, RESOURCE_ENTRY_LABEL)

	if err != nil {
		return "0", err
	}

	// create relation
	_feedId, _ := strconv.Atoi(feed.Id)
	rel, err := storage.RelateNodes(_feedId, entry.Id, "entry", nil)

	if err != nil || rel.Type == "" {
		return "0", err
	}

	feedEntry.Id = strconv.Itoa(entry.Id)

	// notify
	d, _ := json.Marshal(feedEntry)
	// SHOULD CREATE AND TRIGGER EVENT VIA SYSTEM EVENT MANAGER
	// STREAM SERVICE SHOUD LISTEN FOR IT AND STREAM TO CONNECTED CLIENTS
	room.FeedRoom.Publish <- room.NewFeedEvent(room.FEED_ENTRY_NEW, feed.Id, string(d))

	return feedEntry.Id, nil
}

func UpdateEntry(id string, FeedId string, ApplicationId string, OrgId string, data string) (err error) {
	entry, err := GetEntry(id, FeedId, ApplicationId, OrgId)

	if err != nil {
		return err
	}

	// update entry
	entry.Data = data

	// notify
	d, _ := json.Marshal(entry)
	// SHOULD CREATE AND TRIGGER EVENT VIA SYSTEM EVENT MANAGER
	// STREAM SERVICE SHOUD LISTEN FOR IT AND STREAM TO CONNECTED CLIENTS
	room.FeedRoom.Publish <- room.NewEntryEvent(room.ENTRY_UPDATE, entry.Id, string(d))

	_id, _ := strconv.Atoi(entry.Id)
	return storage.SetPropertyNode(_id, "data", data)
}

func DeleteEntry(id string, FeedId string, ApplicationId string, OrgId string) (error) {
	entry, err := GetEntry(id, FeedId, ApplicationId, OrgId)

	if err != nil {
		return err
	}

	_id, _ := strconv.Atoi(entry.Id)
	_rels, _ := storage.RelationshipsNode(_id, "entry")

	for _, rel := range _rels {
		storage.DeleteRelation(rel.Id)
	}

	// notify
	d, _ := json.Marshal(entry)
	// SHOULD CREATE AND TRIGGER EVENT VIA SYSTEM EVENT MANAGER
	// STREAM SERVICE SHOUD LISTEN FOR IT AND STREAM TO CONNECTED CLIENTS
	room.FeedRoom.Publish <- room.NewEntryEvent(room.ENTRY_DELETE, entry.Id, string(d))

	return storage.DeleteNode(_id)
}
