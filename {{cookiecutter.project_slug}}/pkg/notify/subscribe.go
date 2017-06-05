package notify

import (
	bitnotify "github.com/bitly/go-notify"
)

// A Message must have a unique ID that can be passed to
// SubscriberFunc's with added data
type Message interface {
	ID() string
}

// A SubscriberFunc handles receiving a Message
// e.g. busts a cache when the Message contains a cache key
type SubscriberFunc func(Message)

// Subscribe registers a SubscriberFunc for a given messageID.
// When a Message arrives with that ID() it will be passed
// to the SubscriberFunc.
// Subscribe starts the listener and SubscriberFunc in a separate Goroutine
func Subscribe(messageID string, sf SubscriberFunc) {
	eventChan := make(chan interface{})
	bitnotify.Start(messageID, eventChan)
	go func() {
		for {
			data := <-eventChan
			sf(data.(Message))
		}
	}()
}

// Notify receives a Message and notifies any SubscriberFuncs
// listening for Message.ID()
func Notify(message Message) {
	bitnotify.Post(message.ID(), message)
}
