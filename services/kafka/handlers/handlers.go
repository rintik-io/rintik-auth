package handlers

import (
	"github.com/fahmyabdul/golibs/kafka"
)

// Handlers :
type Handlers struct {
	TopicName  string
	JSONString []byte
	Message    kafka.UniformPublishMessage
}

// Handle :
func (a *Handlers) Handle(topicName string, message kafka.UniformPublishMessage) error {
	a = &Handlers{}
	a.TopicName = topicName
	a.Message = message

	return nil
}
