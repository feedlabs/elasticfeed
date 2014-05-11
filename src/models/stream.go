package models

import (
	"errors"
	"strconv"
	"time"
)

func init() {
	Streams = make(map[string]*Stream)
	Streams["1"] = &Stream{"1", "foo"}
	Streams["2"] = &Stream{"2", "bar"}
	Streams["3"] = &Stream{"3", "foobar"}
}

func AddStream(stream Stream) (StreamId string) {
	stream.StreamId = strconv.FormatInt(time.Now().UnixNano(), 10)
	Streams[stream.StreamId] = &stream
	return stream.StreamId
}

func GetStream(StreamId string) (stream *Stream, err error) {
	if v, ok := Streams[StreamId]; ok {
		return v, nil
	}
	return nil, errors.New("StreamId Not Exist")
}

func GetStreamList() map[string]*Stream {
	return Streams
}

func UpdateStream(StreamId string, Name string) (err error) {
	if v, ok := Streams[StreamId]; ok {
		v.Name = Name
		return nil
	}
	return errors.New("StreamId Not Exist")
}

func DeleteStream(StreamId string) {
	delete(Streams, StreamId)
}
