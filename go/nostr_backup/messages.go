package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type EventMessage struct {
	Label string
	SubscriptionId string
	Event Event
}

type ReqMessage struct {
	Label string
	SubscriptionId string
	Filters Filters
}

func (em EventMessage) MarshalJSON() ([]byte, error) {
	type localEventMessage EventMessage
	valueJson, err := json.Marshal(localEventMessage(em))
	if err != nil {
		return nil, err
	}

	return []byte(fmt.Sprintf("[%s]", valueJson)), nil
}

func (em ReqMessage) MarshalJSON() ([]byte, error) {
	type localReqMessage ReqMessage
	valueJson, err := json.Marshal(localReqMessage(em))
	if err != nil {
		return nil, err
	}

	return []byte(fmt.Sprintf("[%s]", valueJson)), nil
}

func JsonToEventMessage(eventJson string) (subscriptionId string, event Event) {
	if !json.Valid([]byte(eventJson)) {
		log.Fatal("Trying to process invalid JSON for an event message") // @MarkFix
	}

	var message []json.RawMessage
	err := json.Unmarshal([]byte(eventJson), &message)
	if err != nil {
		log.Fatal(err)
	}
	if len(message) != 3 {
		log.Fatal("Event message should be an array of length 3")
	}
	var label string
	err = json.Unmarshal(message[0], &label)
	if err != nil {
		log.Fatal(err)
	}
	if label != "EVENT" {
		log.Fatal("Trying to process non-event message as an event message")
	}

	err = json.Unmarshal(message[1], &subscriptionId)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(message[2], &event)
	if err != nil {
		log.Fatal(err)
	}

	generatedId := GenerateEventId(event)
	if event.Id != generatedId { // @MarkFix
		log.Fatalf("event.Id different to GenerateEventId, %s / %s\n", event.Id, generatedId)
	}

	return subscriptionId, event
}
