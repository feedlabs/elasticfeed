package models

import (
	"errors"
	"strconv"
	"time"
)

func init() {
	Feeds = make(map[string]*Feed)
	Feeds["1"] = &Feed{"1", "foo"}
	Feeds["2"] = &Feed{"2", "bar"}
	Feeds["3"] = &Feed{"3", "foobar"}
}

func AddFeed(feed Feed) (FeedId string) {
	feed.FeedId = strconv.FormatInt(time.Now().UnixNano(), 10)
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

func UpdateFeed(FeedId string, Name string) (err error) {
	if v, ok := Feeds[FeedId]; ok {
		v.Name = Name
		return nil
	}
	return errors.New("FeedId Not Exist")
}

func DeleteFeed(FeedId string) {
	delete(Feeds, FeedId)
}
