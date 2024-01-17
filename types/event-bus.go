package types

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

type Subscriber func(data interface{}, eventName string) error

type SubscriberContext struct {
	SubscriberId uuid.UUID
}

type SubscriberDescriptor struct {
	Id         uuid.UUID
	Subscriber Subscriber
}

type EventHandler func(data interface{}, eventName string) error

type EmitData struct {
	EventName string
	Data      core.JSONB
	Options   core.JSONB
}
