package resources

import (
	"errors"
	"strconv"

	"github.com/feedlabs/feedify/graph"
)

func (this *Feed) AddEntry(entry Entry) (EntryId string, err error) {
	return AddEntry(entry, this.Id, this.ApplicationId)
}

func (this *Feed) GetEntryList() (entries []*Entry, err error) {
	return GetEntryList(this.Id, this.ApplicationId)
}

func GetFeedList(appId string) (feedList []*Feed, err error) {
	nodes, err := storage.FindNodesByLabel(RESOURCE_FEED_LABEL)
	if err != nil {
		nodes = nil
	}

	var feeds []*Feed

	for _, node := range nodes {
		data := node.Data["data"].(string)
		id := strconv.Itoa(node.Id)
		rels, _ := storage.RelationshipsNode(node.Id, "contains")

		feed := &Feed{id , appId, data, len(rels)}
		feeds = append(feeds, feed)
	}

	return feeds, err
}

func GetFeed(id string, applicationId string) (feed *Feed, err error) {
	_id, err := strconv.Atoi(id)
	node, err := storage.Node(_id)

	if err != nil {
		return nil, err
	}

	if node != nil && contains(node.Labels, RESOURCE_FEED_LABEL) {
		data := node.Data["data"].(string)
		rels, _ := storage.RelationshipsNode(node.Id, "contains")
		return &Feed{strconv.Itoa(node.Id), applicationId, data, len(rels)}, nil
	}

	return nil, errors.New("FeedId not exist")
}

func AddFeed(feed Feed, applicationId string) (id string, err error) {
	// get feed
	app, err := GetApplication(applicationId)
	if err != nil {
		return "0", err
	}

	// add feed
	properties := graph.Props{"applicationId": applicationId, "data": feed.Data}
	_feed, err := storage.NewNode(properties, RESOURCE_FEED_LABEL)

	if err != nil {
		return "0", err
	}

	// create relation
	_appId, _ := strconv.Atoi(app.Id)
	rel, err := storage.RelateNodes(_appId, _feed.Id, "contains", nil)

	if err != nil || rel.Type == "" {
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

func init() {
	Feeds = make(map[string]*Feed)
}
