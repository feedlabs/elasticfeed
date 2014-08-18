package entity

var (
	Users map[string]*User
)

type User struct {
	UserId	string
	Name	string
}

var (
	Feeds map[string]*Feed
)

type Feed struct {
	FeedId	string
	Name	string
}
