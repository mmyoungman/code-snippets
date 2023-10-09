package main

import (
	"encoding/json"
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
	eventJson, err := json.Marshal(em.Event)
	if err != nil {
		return nil, err
	}

	return []byte(fmt.Sprintf("[\"EVENT\",%s]", eventJson)), nil
}

func (rm ClientReqMessage) MarshalJSON() ([]byte, error) {
	filtersJson, err := json.Marshal(rm.Filters)
	if err != nil {
		return nil, err
	}

	return []byte(fmt.Sprintf("[\"REQ\",\"%s\",%s]", rm.SubscriptionId, filtersJson)), nil
}

func (cm ClientCloseMessage) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("[\"CLOSE\",\"%s\"]", cm.SubscriptionId)), nil
}

