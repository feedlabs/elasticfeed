package stream

type StreamMessage struct {
	message	string
	channel string
	client StreamAdapterStore
}

func (m StreamMessage) Send() {
	m.client.Publish(m.message, m.channel)
}

func (m StreamMessage) SetMessage(message string) {
	m.message = message
}

func (m StreamMessage) SetChannel(channel string) {
	m.channel = channel
}

func (m StreamMessage) SetClient(client StreamAdapterStore) {
	m.client = client
}

func NewStreamMessage(message string, channel string, client StreamAdapterStore) *StreamMessage {
	return &StreamMessage{message, channel, client}
}
