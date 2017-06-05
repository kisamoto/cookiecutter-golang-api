package cache

import (
	"log"

	"{{cookiecutter.repo}}/pkg/notify"
)

const (
	cacheBustMessageID = "CACHE_BUST"
)

// NewBust creates a new notify.Message with the given
// keyExpression (any Redis supported cache key expression)
func NewBust(keyExpression string) BustMessage {
	return BustMessage{
		messageID:     cacheBustMessageID,
		KeyExpression: keyExpression,
	}
}

// NotifyNewBust takes a Key Expression, generates a BustMessage
// and notifies any subscriber
func NotifyNewBust(keyExpression string) {
	notify.Notify(NewBust(keyExpression))
}

// BustMessage contains a KeyExpression to remove keys
// from the Cache
type BustMessage struct {
	messageID     string
	KeyExpression string
}

// The ID implements notify.Message
func (bm BustMessage) ID() string {
	return bm.messageID
}

// BusterSubscriber returns a notify.SubscriberFunc that will
// listen for BustMessages and delete keys from the cache
// with the given expression in the BustMessage
func BusterSubscriber(cache Cache) notify.SubscriberFunc {
	return notify.SubscriberFunc(func(m notify.Message) {
		err := cache.DeleteRaw([]byte(m.(BustMessage).KeyExpression))
		if err != nil {
			log.Println(err)
		}
	})
}
