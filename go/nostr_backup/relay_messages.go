package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type RelayEventMessage struct {
	SubscriptionId string
	Event          Event
}

type RelayOkMessage struct {
	EventId string
	Status  bool
	Message string
}

type RelayEoseMessage struct {
	SubscriptionId string
}

type RelayNoticeMessage struct {
	Message string
}

func ProcessRelayMessage(messageJson string) (label string, message []json.RawMessage) {
	if !json.Valid([]byte(messageJson)) {
		log.Fatal("Message has invalid JSON", messageJson)
	}

	err := json.Unmarshal([]byte(messageJson), &message)
	if err != nil {
		log.Fatal("Could not unmarshal messageJson", err)
	}

	if len(message) < 2 {
		log.Fatal("Relay messages should be an array of at least length 2!", message)
	}

	err = json.Unmarshal(message[0], &label)
	if err != nil {
		log.Fatal(err)
	}

	return label, message[1:]
}

func (em RelayEventMessage) MarshalJSON() ([]byte, error) {
	eventJson, err := json.Marshal(em.Event)
	if err != nil {
		return nil, err
	}

	return []byte(fmt.Sprintf("[\"EVENT\",\"%s\",%s]", em.SubscriptionId, eventJson)), nil
}

func (em RelayEoseMessage) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("[\"EOSE\",\"%s\"]", em.SubscriptionId)), nil
}

func (om RelayOkMessage) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("[\"OK\",\"%s\",%t,\"%s\"]", om.EventId, om.Status, om.Message)), nil
}

func (nm RelayNoticeMessage) MarshalJson() ([]byte, error) {
	return []byte(fmt.Sprintf("[\"NOTICE\",\"%s\"]", nm.Message)), nil
}
