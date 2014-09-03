package entity

import (
	"time"
	"errors"
	"strconv"

	"github.com/feedlabs/feedify/lib/feedify/stream"
)

var (
	message *stream.StreamMessage
)

func init() {
	message, _ = stream.NewStreamMessage()
}

func AddFeedEntry(feedEntry FeedEntry, FeedId string) (FeedEntryId string) {
	feedEntry.Id = strconv.FormatInt(time.Now().UnixNano(), 10)
	Feeds[FeedId].Entries[feedEntry.Id] = &feedEntry

	message.Publish(feedEntry.Data)

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
		return nil
	}
	return errors.New("FeedEntry id " + FeedEntryId + " for Feed id " + FeedId + " does not exist")
}

func DeleteFeedEntry(FeedEntryId string, FeedId string) {
	delete(Feeds[FeedId].Entries, FeedEntryId)
}
