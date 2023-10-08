package main

import (
	"encoding/json"
	"log"
)

func JsonToEventMessage(eventJson string) (subscriptionId string, event Event) {
	if !json.Valid([]byte(eventJson)) {
		log.Fatal("Trying to process invalid JSON for an event message") // @MarkFix
	}

	var message []interface{}
	err := json.Unmarshal([]byte(eventJson), &message)
	if err != nil {
		log.Fatal(err)
	}
	if len(message) != 3 {
		log.Fatal("Event message should be an array of length 3")
	}
	if message[0] != "EVENT" {
		log.Fatal("Trying to process non-event message as an event message")
	}

	subscriptionId = message[1].(string)

	event = parseEventMap(message[2].(map[string]interface{}))

	generatedId := GenerateEventId(event)
	if event.Id != generatedId { // @MarkFix
		log.Fatalf("event.Id different to GenerateEventId, %s / %s\n", event.Id, generatedId)
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
