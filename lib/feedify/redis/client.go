package redis

import (
	"errors"

	"github.com/fzzy/radix/extra/pubsub"
	"github.com/fzzy/radix/redis"

	"github.com/feedlabs/feedify/lib/feedify/config"
)

type RedisClient struct {
	host     string
	port     string
	protocol string
}

func (r RedisClient) Cmd(command string, args ...interface{}) error {
	c, err := redis.Dial(r.protocol, r.host+":"+r.port)
	if err != nil {
		return errors.New("Redis dial error")
	}
	c.Cmd(command, args)
	return nil
}

func (r RedisClient) _subscribe(channel []string, callback func(bool, string, string)) error {
	c, err := redis.Dial(r.protocol, r.host+":"+r.port)
	if err != nil {
		return errors.New("Redis dial error")
	}

	psc := pubsub.NewSubClient(c)
	psr := psc.Subscribe(channel)
	for {
		psr = psc.Receive()
		callback(psr.Timeout(), psr.Message, psr.Channel)
	}

	return nil
}

func (r RedisClient) Subscribe(channel []string, callback func(bool, string, string)) {
	go r._subscribe(channel, callback)
}

func NewRedisClient() *RedisClient {
	host := config.GetConfigKey("redis::host")
	port := config.GetConfigKey("redis::port")
	protocol := config.GetConfigKey("redis::protocol")

	return &RedisClient{host, port, protocol}
}
