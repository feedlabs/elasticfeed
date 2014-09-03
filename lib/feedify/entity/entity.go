package entity

var (
	Feeds map[string]*Feed
)

type Feed struct {
	Id		string
	Data	string
	Entries	map[string]*FeedEntry
}

type FeedEntry struct {
	Id		string
	Data	string
}
