package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type EventMessage struct {
	Label          string
	SubscriptionId string
	Event          Event
}

func JsonToEventMessage(eventJson string) (string, Event) {
	fmt.Println(eventJson)
	var message []interface{}
	err := json.Unmarshal([]byte(eventJson), &message)
	if err != nil {
		log.Fatal(err)
	}
	if len(message) != 3 {
		log.Fatal("Event message should be an array of length 1")
	}
	if message[0] != "EVENT" {
		log.Fatal("Trying to process non-event message as an event message")
	}

	subscriptionId := message[1].(string)

	event := parseEventMap(message[2].(map[string]interface{}))

	if event.Id != GenerateEventId(event) { // @MarkFix
		log.Fatal("event.Id different to GenerateEventId")
	}

	return subscriptionId, event
}

func parseEventMap(eventMap map[string]interface{}) Event {
	tagsInterfaceArr := eventMap["tags"].(interface{}).([]interface{})
	var tags Tags
	for _, object := range tagsInterfaceArr {
		tagInterface := object.([]interface{})
		var tag Tag
		for _, tagElem := range tagInterface {
			tag = append(tag, tagElem.(string))

		}
		tags = append(tags, tag)
	}

	var event Event
	event.Id = eventMap["id"].(string)
	event.PubKey = eventMap["pubkey"].(string)
	event.CreatedAt = int64(eventMap["created_at"].(float64))
	event.Kind = int(eventMap["kind"].(float64))
	event.Tags = tags
	event.Content = eventMap["content"].(string)
	event.Sig = eventMap["sig"].(string)

	return event
}
