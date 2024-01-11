package types

import "github.com/driver005/gateway/core"

type Subscriber func(data interface{}, eventName string) error

type SubscriberContext struct {
	SubscriberId string
}

type SubscriberDescriptor struct {
	Id         string
	Subscriber Subscriber
}

type EventHandler func(data interface{}, eventName string) error

type EmitData struct {
	EventName string
	Data      core.JSONB
	Options   core.JSONB
}
