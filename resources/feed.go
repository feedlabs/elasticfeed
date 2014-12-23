package resources

import (
	"errors"
	"strconv"

	"github.com/feedlabs/feedify/graph"
)

func init() {
	Feeds = make(map[string]*Feed)
}

func GetFeedList() []*Feed {
	nodes, err := storage.FindNodesByLabel(RESOURCE_FEED_LABEL)
	if err != nil {
		nodes = nil
	}

	var feeds []*Feed

	for _, node := range nodes {
		data := node.Data["data"].(string)
		id := strconv.Itoa(node.Id)
		rels, _ := storage.RelationshipsNode(node.Id, "contains")

		feed := &Feed{id , data, len(rels)}
		feeds = append(feeds, feed)
	}

	return feeds
}

func GetFeed(id string) (feed *Feed, err error) {
	_id, err := strconv.Atoi(id)
	node, err := storage.Node(_id)

	if err != nil {
		return nil, err
	}

	if node != nil && contains(node.Labels, RESOURCE_FEED_LABEL) {
		data := node.Data["data"].(string)
		rels, _ := storage.RelationshipsNode(node.Id, "contains")
		return &Feed{strconv.Itoa(node.Id), data, len(rels)}, nil
	}

	return nil, errors.New("FeedId not exist")
}

func AddFeed(feed Feed) (id string, err error) {
	properties := graph.Props{"data": feed.Data}
	_feed, err := storage.NewNode(properties, RESOURCE_FEED_LABEL)

	if err != nil {
		return "0", err
	}

	feed.Id = strconv.Itoa(_feed.Id)

	return feed.Id, nil
}

func UpdateFeed(id string, data string) (err error) {
	_id, _ := strconv.Atoi(id)
	return storage.SetPropertyNode(_id, "data", data)
}

func DeleteFeed(id string) (error) {
	_id, _ := strconv.Atoi(id)
	return storage.DeleteNode(_id)
}
