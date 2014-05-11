package models

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

var (
	Storages map[string]*Storage
)

type Storage struct {
	StorageId	string
	Name		string
}

var (
	Streams map[string]*Stream
)

type Stream struct {
	StreamId	string
	Name		string
}
