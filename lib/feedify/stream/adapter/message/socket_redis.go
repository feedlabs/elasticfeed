package message

import (
	"github.com/feedlabs/feedify/lib/feedify/redis"
	"github.com/feedlabs/feedify/lib/feedify/stream"
)

func init() {
	stream.RegisterAdapterStore("socket_redis", newAdapterStore, createNewRedisClient)
}

func createNewRedisClient(options stream.Options) error {
	return nil
}

func newAdapterStore(options stream.Options) (stream.StreamAdapterStore, error) {
	client := redis.NewRedisClient()
	return &StreamAdapterStore{client}, nil
}

type StreamAdapterStore struct {
	client *redis.RedisClient
}

func (m StreamAdapterStore) Publish(message string, channel string) {
	m.client.Cmd("publish", channel, message)
}

func (m StreamAdapterStore) Name() string {
	return "socket_redis"
}

func (m StreamAdapterStore) Subscribe(channels []string, callback func(bool, string, string)) {
	m.client.Subscribe(channels, callback)
}

func (m StreamAdapterStore) Connect() {}
