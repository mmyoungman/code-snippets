package main

import (
	"fmt"
)

type ClientEventMessage struct {
	Event Event
}

type ClientReqMessage struct {
	SubscriptionId string
	Filters Filters
}

type ClientCloseMessage struct {
	SubscriptionId string
}

func (em ClientEventMessage) MarshalJSON() ([]byte, error) {
	panic("Use ToJson")
}

func (rm ClientReqMessage) MarshalJSON() ([]byte, error) {
	panic("Use ToJson")
}

func (cm ClientCloseMessage) MarshalJSON() ([]byte, error) {
	panic("Use ToJson")
}

func (em ClientEventMessage) ToJson() string {
	return fmt.Sprintf("[\"EVENT\",%s]", em.Event.ToJson())
}

func (rm ClientReqMessage) ToJson() string {
	return fmt.Sprintf("[\"REQ\",\"%s\",%s]",
		rm.SubscriptionId, rm.Filters.ToJson())
}

func (cm ClientCloseMessage) ToJson() string {
	return fmt.Sprintf("[\"CLOSE\",\"%s\"]", cm.SubscriptionId)
}

