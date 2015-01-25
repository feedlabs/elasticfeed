package resource

import (
	"errors"
	"strconv"
	"encoding/json"

	"github.com/feedlabs/feedify/graph"

	"github.com/feedlabs/elasticfeed/stream/controller/room"
	"github.com/feedlabs/elasticfeed/stream/model"
)

func (this *Feed) AddEntry(entry Entry) (EntryId string, err error) {
	return AddEntry(entry, this.Id, this.Application.Id, this.Application.Org.Id)
}

func (this *Feed) GetEntryList() (entries []*Entry, err error) {
	return GetEntryList(this.Id, this.Application.Id, this.Application.Org.Id)
}

func GetFeedList(ApplicationId string, OrgId string) (feedList []*Feed, err error) {
	app, err := GetApplication(ApplicationId, OrgId)
	if err != nil {
		return nil, err
	}

	_id, _ := strconv.Atoi(app.Id)
	_rels, _ := storage.RelationshipsNode(_id, "feed")

	var feeds []*Feed

	for _, rel := range _rels {
		data := rel.EndNode.Data["data"].(string)
		id := strconv.Itoa(rel.EndNode.Id)
		rels, _ := storage.RelationshipsNode(rel.EndNode.Id, "entry")

		feed := &Feed{id , app, data, len(rels)}
		feeds = append(feeds, feed)
	}

	if feeds == nil {
		feeds = make([]*Feed, 0)
	}

	return feeds, err
}

func GetFeed(id string, applicationId string, orgId string) (feed *Feed, err error) {
	app, err := GetApplication(applicationId, orgId)
	if err != nil {
		return nil, err
	}

	_id, err := strconv.Atoi(id)
	node, err := storage.Node(_id)

	if err != nil {
		return nil, err
	}

	if node != nil && Contains(node.Labels, RESOURCE_FEED_LABEL) && app.Id == node.Data["applicationId"].(string) {
		data := node.Data["data"].(string)
		rels, _ := storage.RelationshipsNode(node.Id, "entry")
		return &Feed{strconv.Itoa(node.Id), app, data, len(rels)}, nil
	}

	return nil, errors.New("FeedId not exist for ApplicationId `"+applicationId+"`")
}

func AddFeed(feed Feed, applicationId string, orgId string) (id string, err error) {
	// get feed
	app, err := GetApplication(applicationId, orgId)
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
	rel, err := storage.RelateNodes(_appId, _feed.Id, "feed", nil)

	if err != nil || rel.Type == "" {
		return "0", err
	}

	feed.Id = strconv.Itoa(_feed.Id)

	// notify
	data, _ := json.Marshal(feed)
	room.Publish <- room.NewEvent(model.EVENT_MESSAGE, "system", string(data))

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
