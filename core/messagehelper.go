package core

const (
	// ApplyToPayloadString defines the reserved word denoting the payload field
	// when using the parameter "ApplyTo" in formatters
	ApplyToPayloadString = "payload"
)

// GetAppliedContent is a func() to get message content from payload or meta data
// for later handling by plugins
type GetAppliedContent func(msg *Message) []byte

// SetAppliedContent is a func() to store message content to payload or meta data
type SetAppliedContent func(msg *Message, content []byte)

// GetAppliedContentGetFunction returns a GetAppliedContent function
func GetAppliedContentGetFunction(applyTo string) GetAppliedContent {
	if applyTo != "" && applyTo != ApplyToPayloadString {
		return func(msg *Message) []byte {
			return msg.MetaData().GetValue(applyTo)
		}
	}

	return func(msg *Message) []byte {
		return msg.Data()
	}
}

// GetAppliedContentSetFunction returns SetAppliedContent function to store message content
func GetAppliedContentSetFunction(applyTo string) SetAppliedContent {
	if applyTo != "" && applyTo != ApplyToPayloadString {
		return func(msg *Message, content []byte) {
			msg.MetaData().SetValue(applyTo, content)
			return
		}
	}

	return func(msg *Message, content []byte) {
		msg.Store(content)
		return
	}
}
