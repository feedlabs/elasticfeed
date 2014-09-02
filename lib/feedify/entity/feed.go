package entity

import (
	"time"
	"errors"
	"strconv"
)

func init() {
	Feeds = make(map[string]*Feed)
}

func AddFeed(feed Feed) (FeedId string) {
	feed.FeedId = strconv.FormatInt(time.Now().UnixNano(), 10)

	feed.Entries = make(map[string]*FeedEntry)
	Feeds[feed.FeedId] = &feed

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
