package resources

import (
	"errors"
	"strconv"
	"time"
)

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

func AddFeedEntry(feedEntry FeedEntry, FeedId string) (FeedEntryId string) {
	feedEntry.Id = strconv.FormatInt(time.Now().UnixNano(), 10)
	Feeds[FeedId].Entries[feedEntry.Id] = &feedEntry

	_data := BODY_HEADER + `"Id": "` + feedEntry.Id + `", "Action": "add", "Tag": {}, "Data": "` + feedEntry.Data + `"` + BODY_BOTTOM
	message.Publish(_data)

	return feedEntry.Id
}

func GetFeedEntry(FeedEntryId string, FeedId string) (feedEntry *FeedEntry, err error) {
	if v, ok := Feeds[FeedId].Entries[FeedEntryId]; ok {
		return v, nil
	}
	return nil, errors.New("FeedEntryId for FeedId Not Exist")
}

func GetFeedEntryList(FeedId string) map[string]*FeedEntry {
	return Feeds[FeedId].Entries
}

func UpdateFeedEntry(FeedEntryId string, FeedId string, data string) (err error) {
	if v, ok := Feeds[FeedId].Entries[FeedEntryId]; ok {
		v.Data = data

		_data := BODY_HEADER + `"Id": "` + FeedEntryId + `", "Action": "update", "Tag": {}, "Data": "` + data + `"` + BODY_BOTTOM
		message.Publish(_data)

		return nil
	}
	return errors.New("FeedEntry id " + FeedEntryId + " for Feed id " + FeedId + " does not exist")
}

func DeleteFeedEntry(FeedEntryId string, FeedId string) {

	_data := BODY_HEADER + `"Id": "` + FeedEntryId + `", "Action": "delete"` + BODY_BOTTOM
	message.Publish(_data)

	delete(Feeds[FeedId].Entries, FeedEntryId)
}
