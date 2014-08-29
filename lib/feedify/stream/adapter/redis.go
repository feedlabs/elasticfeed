package adapter

import (
	"fmt"
	"github.com/fzzy/radix/redis"
	"github.com/fzzy/radix/extra/pubsub"
	"github.com/feedlabs/feedify/lib/feedify/stream"
)

func init() {
	stream.RegisterAdapterStore("redis", newAdapterStore, createNewRedisClient)
}

func createNewRedisClient(options stream.Options) error {

	return nil
}

func newAdapterStore(options stream.Options) (stream.StreamAdapterStore, error) {
	return &StreamAdapterStore{"localhost", "6379", "tcp"}, nil
}

type StreamAdapterStore struct {
	host string
	port string
	protocol string
}

func (m StreamAdapterStore) Publish(message string, channel string) {
	c, err := redis.Dial(m.protocol, m.host + ":" + m.port)
	if err != nil {
		fmt.Println("error:", err)
	}
	c.Cmd("publish", channel, message)
}

func (m StreamAdapterStore) Name() string {
	return "redis"
}

func (m StreamAdapterStore) Subscribe(channels []string) {
	go m._subscribe(channels)
}

func (m StreamAdapterStore) Connect() {}

func (m StreamAdapterStore) _subscribe(channel []string) {
	c, err := redis.Dial(m.protocol, m.host + ":" + m.port)
	if err != nil {
		fmt.Println("error:", err)
	}

	psc := pubsub.NewSubClient(c)
	psr := psc.Subscribe(channel)
	for {
		psr = psc.Receive() //Blocks until reply is received or timeout is tripped
		if !psr.Timeout() {
			fmt.Println("publish:", psr.Message, " channel:", psr.Channel)
		} else {
			fmt.Println("error: sub timedout")
			return
		}
	}
}
