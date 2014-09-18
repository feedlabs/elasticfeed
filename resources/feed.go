package resources

import (
	"errors"
	"strconv"

	"github.com/feedlabs/feedify/graph"
)

func init() {
	Feeds = make(map[string]*Feed)
}

func AddFeed(feed Feed) (id string) {
	properties := graph.Props{"data": feed.Data}
	label := "feed"
	node, _ := storage.NewNode(properties, label)

	feed.Id = strconv.Itoa(node.Id)

	return feed.Id
}

func GetFeed(id string) (feed *Feed, err error) {
	_id, err := strconv.Atoi(id)
	node, err := storage.Node(_id)

	if err != nil {
		return nil, err
	}

	if node != nil {
		data := node.Data["data"].(string)
		return &Feed{strconv.Itoa(node.Id), data, nil}, nil
	}

	return nil, errors.New("Id not exist")
}

func GetFeedList() map[string]*Feed {
	nodes, err := storage.FindNodesByLabel("feed")
	if err != nil {
		nodes = nil
	}

	Feeds = make(map[string]*Feed)

	for _, node := range nodes {
		data := node.Data["data"].(string)
		id := strconv.Itoa(node.Id)
		Feeds[id] = &Feed{id , data, nil}
		Feeds[id].Entries = make(map[string]*FeedEntry)
	}

	return Feeds
}

func UpdateFeed(id string, data string) (err error) {
	_id, _ := strconv.Atoi(id)
	return storage.SetPropertyNode(_id, "data", data)
}

func DeleteFeed(id string) (error) {
	_id, _ := strconv.Atoi(id)
	return storage.DeleteNode(_id)
}
