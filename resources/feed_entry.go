package resources

import (
	"errors"
	"strconv"

	"github.com/feedlabs/feedify/graph"
)

const RESOURCE_ENTRY_LABEL = "entry"

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

func AddFeedEntry(feedEntry FeedEntry, FeedId string) (FeedEntryId string, err error) {
	// get feed
	feed, err := GetFeed(FeedId)
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

func GetFeedEntry(id string, FeedId string) (feedEntry *FeedEntry, err error) {
	feed, err := GetFeed(FeedId)
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
		return &FeedEntry{strconv.Itoa(entry.Id), FeedId, data}, nil
	}

	return nil, errors.New("EntryId not exist")
}

func GetFeedEntryList(FeedId string) (feedEntries []*FeedEntry, err error) {
	feed, err := GetFeed(FeedId)

	if err != nil {
		return nil, err
	}

	_id, _ := strconv.Atoi(feed.Id)
	_rels, _ := storage.RelationshipsNode(_id, "contains")

	var entries []*FeedEntry

	for _, rel := range _rels {
		data := rel.EndNode.Data["data"].(string)
		entry := &FeedEntry{strconv.Itoa(rel.EndNode.Id), FeedId, data}
		entries = append(entries, entry)
	}

	return entries, nil
}

func UpdateFeedEntry(id string, FeedId string, data string) (err error) {
	entry, err := GetFeedEntry(id, FeedId)

	if err != nil {
		return err
	}

	_data := BODY_HEADER + `"Id": "` + entry.Id + `", "Action": "update", "Tag": {}, "Data": ` + strconv.Quote(data) + BODY_BOTTOM
	message.Publish(_data)

	_id, _ := strconv.Atoi(entry.Id)
	return storage.SetPropertyNode(_id, "data", data)
}

func DeleteFeedEntry(id string, FeedId string) (error) {
	entry, err := GetFeedEntry(id, FeedId)

	if err != nil {
		return err
	}

	_data := BODY_HEADER + `"Id": "` + entry.Id + `", "Action": "remove"` + BODY_BOTTOM
	message.Publish(_data)

	_id, _ := strconv.Atoi(id)
	return storage.DeleteNode(_id)
}
