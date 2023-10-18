package main

import (
	"fmt"
	"log"
	"mmyoungman/nostr_backup/json_wrapper"
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

func (em RelayEventMessage) ToJson() string {
	result := fmt.Sprintf("[\"EVENT\",\"%s\",%s]",
		em.SubscriptionId, em.Event.ToJson())
	DevBuildValidJson(result)
	return result
}

func (em RelayEoseMessage) ToJson() string {
	result := fmt.Sprintf("[\"EOSE\",\"%s\"]", em.SubscriptionId)
	DevBuildValidJson(result)
	return result
}

func (om RelayOkMessage) ToJson() string {
	result := fmt.Sprintf("[\"OK\",\"%s\",%t,\"%s\"]",
		om.EventId, om.Status, om.Message)
	DevBuildValidJson(result)
	return result
}

func (nm RelayNoticeMessage) ToJson() string {
	result := fmt.Sprintf("[\"NOTICE\",\"%s\"]", nm.Message)
	DevBuildValidJson(result)
	return result
}

func ProcessRelayMessage(messageJson string) (label string, message RawJsonArray) {
	if !json_wrapper.IsValidJson(messageJson) {
		log.Fatal("Message has invalid JSON", messageJson)
	}

	err := json_wrapper.UnmarshalJSON([]byte(messageJson), &message)
	if err != nil {
		log.Fatal("Could not unmarshal messageJson", err)
	}

	if len(message) < 2 {
		log.Fatal("Relay messages should be an array of at least length 2!", message)
	}

	err = json_wrapper.UnmarshalJSON(message[0], &label)
	if err != nil {
		log.Fatal(err)
	}

	return label, message[1:]
}
