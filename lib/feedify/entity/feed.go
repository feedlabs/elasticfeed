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
	Feeds = make(map[string]*Feed)
	Feeds["1"] = &Feed{"1", "foo"}
	Feeds["2"] = &Feed{"2", "bar"}
	Feeds["3"] = &Feed{"3", "foobar"}

	message, _ = stream.NewStreamMessage()
}

func AddFeed(feed Feed) (FeedId string) {
	feed.FeedId = strconv.FormatInt(time.Now().UnixNano(), 10)
	Feeds[feed.FeedId] = &feed

	message.Publish(feed.Data)

	return feed.FeedId
}

func GetFeed(FeedId string) (feed *Feed, err error) {
	if v, ok := Feeds[FeedId]; ok {
		return v, nil
	}
	return nil, errors.New("FeedId Not Exist")
}

func GetFeedList() map[string]*Feed {
	return Feeds
}

func UpdateFeed(FeedId string, Data string) (err error) {
	if v, ok := Feeds[FeedId]; ok {
		v.Data = Data
		return nil
	}
	return errors.New("FeedId Not Exist")
}

func DeleteFeed(FeedId string) {
	delete(Feeds, FeedId)
}
