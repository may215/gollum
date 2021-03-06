// Copyright 2015-2017 trivago GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import (
	"github.com/golang/protobuf/proto"
	"github.com/trivago/tgo/tcontainer"
	"time"
)

// MessageData is a container for the message payload, streamID and an optional message key
// The struct is used by Message.data for the current message data and orig for the original message data
type MessageData struct {
	payload  []byte
	streamID MessageStreamID
	MetaData MetaData
}

// Message is a container used for storing the internal state of messages.
// This struct is passed between consumers and producers.
type Message struct {
	data         MessageData
	orig         MessageData
	prevStreamID MessageStreamID
	source       MessageSource
	timestamp    time.Time
}

var (
	// MessageDataPool is the pool used for message payloads.
	// This pool should be used to allocate temporary buffers for e.g.
	// formatters.
	MessageDataPool = tcontainer.NewBytePool()
)

// NewMessage creates a new message from a given data stream by copying data.
func NewMessage(source MessageSource, data []byte, streamID MessageStreamID) *Message {
	buffer := getPayloadCopy(data)
	origBuffer := getPayloadCopy(data)

	message := &Message{
		source:       source,
		prevStreamID: streamID,
		timestamp:    time.Now(),
	}

	message.data.payload = buffer
	message.data.streamID = streamID
	message.data.MetaData = MetaData{}

	message.orig.payload = origBuffer
	message.orig.streamID = streamID
	message.orig.MetaData = MetaData{}

	return message
}

// NewMessageWithSize creates a new message with a buffer of a given size.
// The buffer may contain data from previous messages.
func NewMessageWithSize(source MessageSource, dataSize int, sequence uint64, streamID MessageStreamID) *Message {
	buffer := MessageDataPool.Get(dataSize)
	origBuffer := getPayloadCopy(buffer)

	message := &Message{
		source:       source,
		prevStreamID: streamID,
		timestamp:    time.Now(),
	}

	message.data.payload = buffer
	message.data.streamID = streamID
	message.data.MetaData = MetaData{}

	message.orig.payload = origBuffer
	message.orig.streamID = streamID
	message.orig.MetaData = MetaData{}

	return message
}

// getPayloadCopy return a copy of the data byte array
func getPayloadCopy(data []byte) []byte {
	buffer := MessageDataPool.Get(len(data))
	copy(buffer, data)

	return buffer
}

// Created returns the time when this message was created.
func (msg *Message) Created() time.Time {
	return msg.timestamp
}

// StreamID returns the stream this message is currently routed to.
func (msg *Message) StreamID() MessageStreamID {
	return msg.data.streamID
}

// PreviousStreamID returns the last "hop" of this message.
func (msg *Message) PreviousStreamID() MessageStreamID {
	return msg.prevStreamID
}

// GetRouter returns the stream object behind the current StreamID.
func (msg *Message) GetRouter() Router {
	return StreamRegistry.GetRouterOrFallback(msg.StreamID())
}

// GetPreviousRouter returns the stream object behind the previous StreamID.
func (msg *Message) GetPreviousRouter() Router {
	return StreamRegistry.GetRouterOrFallback(msg.prevStreamID)
}

// SetStreamID sets a new stream and stores the current one in the previous
// stream field.
func (msg *Message) SetStreamID(streamID MessageStreamID) {
	msg.prevStreamID = msg.StreamID()
	msg.data.streamID = streamID
}

// Source returns the message's source (can be nil).
func (msg *Message) Source() MessageSource {
	return msg.source
}

// String implements the stringer interface
func (msg *Message) String() string {
	return string(msg.data.payload)
}

// Data returns the stored data
func (msg *Message) Data() []byte {
	return msg.data.payload
}

// Len returns the length of the current data buffer
func (msg *Message) Len() int {
	return len(msg.data.payload)
}

// Cap returns the capacity of the current data buffer
func (msg *Message) Cap() int {
	return cap(msg.data.payload)
}

// MetaData returns the current MetaData
func (msg *Message) MetaData() MetaData {
	return msg.data.MetaData
}

// Store copies data into the hold data buffer. If the buffer can hold data
// it is resized, otherwise a new buffer will be allocated.
func (msg *Message) Store(data []byte) {
	copy(msg.Resize(len(data)), data)
}

// Offset moves the slice start offset of the currently stored data to the
// given position. This can be used to e.g. efficiently crop of the beginning
// of a message.
func (msg *Message) Offset(offset int) {
	msg.data.payload = msg.data.payload[offset:]
}

// Resize changes the size of the stored buffer. The current content is not
// guaranteed to be preserved. If content needs to be preserved use Extend.
func (msg *Message) Resize(size int) []byte {
	switch {
	case size == len(msg.data.payload):
	case size <= cap(msg.data.payload):
		msg.data.payload = msg.data.payload[:size]
	default:
		msg.data.payload = MessageDataPool.Get(size)
	}

	return msg.data.payload
}

// Extend changes the size of the stored buffer. The current content will be
// preserved. If content does not need to be preserved use Resize.
func (msg *Message) Extend(size int) []byte {
	switch {
	case size == len(msg.data.payload):
	case size <= cap(msg.data.payload):
		msg.data.payload = msg.data.payload[:size]
	default:
		old := msg.data.payload
		msg.data.payload = MessageDataPool.Get(size)
		copy(msg.data.payload, old)
	}

	return msg.data.payload
}

// Clone returns a copy of this message, i.e. the payload is duplicated.
// The created timestamp is copied, too.
func (msg *Message) Clone() *Message {
	clone := *msg

	clone.data.payload = MessageDataPool.Get(len(msg.data.payload))
	copy(clone.data.payload, msg.data.payload)

	return &clone
}

// CloneOriginal returns a copy of this message with the original payload and stream
// The created timestamp is copied, too.
func (msg *Message) CloneOriginal() *Message {
	clone := *msg

	clone.data.payload = MessageDataPool.Get(len(msg.orig.payload))
	copy(clone.data.payload, msg.orig.payload)

	clone.SetStreamID(msg.orig.streamID)

	return &clone
}

// Serialize generates a string containing all data that can be preserved over
// shutdown (i.e. no data directly referencing runtime components).
func (msg Message) Serialize() ([]byte, error) {
	serializable := &SerializedMessage{
		StreamID:     proto.Uint64(uint64(msg.data.streamID)),
		PrevStreamID: proto.Uint64(uint64(msg.prevStreamID)),
		Timestamp:    proto.Int64(msg.timestamp.UnixNano()),
		Data:         msg.data.payload,
		MetaData:     msg.data.MetaData,
	}

	return proto.Marshal(serializable)
}

// DeserializeMessage generates a message from a string produced by
// Message.Serialize.
func DeserializeMessage(data []byte) (Message, error) {
	serializable := new(SerializedMessage)
	err := proto.Unmarshal(data, serializable)

	msg := Message{
		prevStreamID: MessageStreamID(serializable.GetPrevStreamID()),
		timestamp:    time.Unix(0, serializable.GetTimestamp()),
	}

	msg.data.streamID = MessageStreamID(serializable.GetStreamID())
	msg.orig.streamID = msg.data.streamID

	msg.data.payload = serializable.GetData()
	copy(msg.orig.payload, msg.data.payload)

	msg.data.MetaData = serializable.GetMetaData()
	msg.orig.MetaData = msg.data.MetaData.Clone()

	return msg, err
}
