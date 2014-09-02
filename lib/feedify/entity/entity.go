package entity

var (
	Users map[string]*User
	Feeds map[string]*Feed
)

type User struct {
	UserId	string
	Name	string
}

type Feed struct {
	FeedId	string
	Data	string
	Entries map[string]*FeedEntry
}

type FeedEntry struct {
	Id		string
	Data	string
}
