package resources

import (
	"errors"
	"strconv"

	"github.com/feedlabs/feedify/graph"
)

// user_feed_token = channel_id + feed_id => e.g aabbccddee + aabbcc
// for private feeds there will be 1 websocket connection
// for public company feeds will be 1 websocket connection
// basically for each channel is 1 websocket connection
// private and public channel will stream through multiple feed-pages events
//
// channel => channel_id
// event => 'feed:' + feed_id
// data => [{				// action object
//		id => string		// entryId
//		tags => strings...	// array of strings
//		action => string	// add/delete/update
//		data => string		// entry data as content; string e.g. json.stringify
// }, {}, {}]

const BODY_HEADER = `{
  "channel": "iO5wshd5fFE5YXxJ/hfyKQ==:17",
  "event": "CM_Action_Abstract:SEND:31",
  "data": {
    "action": {
      "actor": {
        "_type": 33,
        "_id": {
          "id": "1"
        },
        "id": 1,
        "displayName": "user1",
        "visible": true,
        "_class": "Feed_Model_User"
      },
      "verb": 13,
      "type": 31,
      "_class": "Feed_Action_Feed"
    },
    "model": {
      "_type": 33,
      "_id": {
        "id": "1"
      },
      "id": 1,
      "displayName": "user1",
      "visible": true,
      "_class": "Feed_Model_User"
    },
    "data": {`

const BODY_BOTTOM = `
    }
  }
}`

func GetEntryList(FeedId string, ApplicationId string) (feedEntries []*Entry, err error) {
	feed, err := GetFeed(FeedId, ApplicationId)

	if err != nil {
		return nil, err
	}

	_id, _ := strconv.Atoi(feed.Id)
	_rels, _ := storage.RelationshipsNode(_id, "contains")

	var entries []*Entry

	for _, rel := range _rels {
		data := rel.EndNode.Data["data"].(string)
		entry := &Entry{strconv.Itoa(rel.EndNode.Id), FeedId, data}
		if entry != nil && contains(rel.EndNode.Labels, RESOURCE_ENTRY_LABEL) && feed.Id == rel.EndNode.Data["feedId"].(string) {
			entries = append(entries, entry)
		}
	}

	return entries, nil
}

func GetEntry(id string, FeedId string, ApplicationId string) (feedEntry *Entry, err error) {
	feed, err := GetFeed(FeedId, ApplicationId)
	if err != nil {
		return nil, err
	}

	_id, err := strconv.Atoi(id)
	entry, err := storage.Node(_id)

	if err != nil {
		return nil, err
	}

	if entry != nil && contains(entry.Labels, RESOURCE_ENTRY_LABEL) && feed.Id == entry.Data["feedId"].(string) {
		data := entry.Data["data"].(string)
		return &Entry{strconv.Itoa(entry.Id), FeedId, data}, nil
	}

	return nil, errors.New("EntryId `"+id+"` not exist")
}

func AddEntry(feedEntry Entry, FeedId string, ApplicationId string) (EntryId string, err error) {
	// get feed
	feed, err := GetFeed(FeedId, ApplicationId)
	if err != nil {
		return "0", err
	}

	// add feed-entry
	properties := graph.Props{"feedId": FeedId, "data": feedEntry.Data}
	entry, err := storage.NewNode(properties, RESOURCE_ENTRY_LABEL)

	if err != nil {
		return "0", err
	}

	// create relation
	_feedId, _ := strconv.Atoi(feed.Id)
	rel, err := storage.RelateNodes(_feedId, entry.Id, "contains", nil)

	if err != nil || rel.Type == "" {
		return "0", err
	}

	_data := BODY_HEADER + `"Id": "` + feedEntry.Id + `", "Action": "add", "Tag": {}, "Data": ` + strconv.Quote(feedEntry.Data) + BODY_BOTTOM
	message.Publish(_data)

	feedEntry.Id = strconv.Itoa(entry.Id)

	return feedEntry.Id, nil
}

func UpdateEntry(id string, FeedId string, ApplicationId string, data string) (err error) {
	entry, err := GetEntry(id, FeedId, ApplicationId)

	if err != nil {
		return err
	}

	_data := BODY_HEADER + `"Id": "` + entry.Id + `", "Action": "update", "Tag": {}, "Data": ` + strconv.Quote(data) + BODY_BOTTOM
	message.Publish(_data)

	_id, _ := strconv.Atoi(entry.Id)
	return storage.SetPropertyNode(_id, "data", data)
}

func DeleteEntry(id string, FeedId string, ApplicationId string) (error) {
	entry, err := GetEntry(id, FeedId, ApplicationId)

	if err != nil {
		return err
	}

	_id, _ := strconv.Atoi(entry.Id)
	_rels, _ := storage.RelationshipsNode(_id, "contains")

	for _, rel := range _rels {
		storage.DeleteRelation(rel.Id)
	}

	_data := BODY_HEADER + `"Id": "` + entry.Id + `", "Action": "remove"` + BODY_BOTTOM
	message.Publish(_data)

	return storage.DeleteNode(_id)
}