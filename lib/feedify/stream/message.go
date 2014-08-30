package stream

import (
	"errors"
)

type StreamMessage struct {
	channel string
	adapter StreamAdapterStore
}

func (m StreamMessage) Publish(data string) {
	m.adapter.Publish(data, m.channel)
}

func (m StreamMessage) Subscribe(channels []string, callback func(bool, string, string)) {
	m.adapter.Subscribe(channels, callback)
}

func (m StreamMessage) SetChannel(channel string) {
	m.channel = channel
}

func (m StreamMessage) SetAdapter(adapter StreamAdapterStore) {
	m.adapter = adapter
}

func (m StreamMessage) GetAdapter() StreamAdapterStore {
	return m.adapter
}

func NewStreamMessage(channel string) (*StreamMessage, error) {
	adapter, err := NewAdapterStore("socket_redis", nil)
	if err != nil {
		return nil, errors.New("Cannot load stream message adapter")
	}

	return &StreamMessage{channel, adapter}, nil
}
