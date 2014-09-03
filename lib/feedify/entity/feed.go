package entity

import (
	"time"
	"errors"
	"strconv"
)

func init() {
	Feeds = make(map[string]*Feed)
}

func AddFeed(feed Feed) (id string) {
	feed.Id = strconv.FormatInt(time.Now().UnixNano(), 10)

	feed.Entries = make(map[string]*FeedEntry)
	Feeds[feed.Id] = &feed

	return feed.Id
}

func GetFeed(id string) (feed *Feed, err error) {
	if v, ok := Feeds[id]; ok {
		return v, nil
	}
	return nil, errors.New("Id not exist")
}

func GetFeedList() map[string]*Feed {
	return Feeds
}

func UpdateFeed(id string, data string) (err error) {
	if v, ok := Feeds[id]; ok {
		v.Data = data
		return nil
	}
	return errors.New("Feed id " + id + " does not exist")
}

func DeleteFeed(id string) {
	delete(Feeds, id)
}
