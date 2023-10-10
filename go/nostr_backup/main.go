package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func main() {
	//npub := "1f0rwg0z2smrkggypqn7gctscevu22z6thch243365xt0tz8fw9uqupzj2x"
	npubHex := "4bc6e43c4a86c764208104fc8c2e18cb38a50b4bbe2eaac63aa196f588e97178"

	event := Event{
		PubKey:    npubHex,
		Kind:      KindTextNote,
		CreatedAt: 0,
		Tags:      make([]Tag, 0),
		Content:   "Test!\n❤️‍🔥\"b\\😅  <html>",
	}
	event.Id = GenerateEventId(event)

	eventJson := event.ToJson()

	fmt.Printf("Event JSON: %s\n", eventJson)

	var eventStruct Event
	_ = json.Unmarshal([]byte(eventJson), &eventStruct)

	fmt.Println(
		"eventStruct: ",
		eventStruct.Id,
		eventStruct.PubKey,
		eventStruct.CreatedAt,
		eventStruct.Kind,
		eventStruct.Tags,
		eventStruct.Content,
		eventStruct.Sig)

	filters := Filters{{
		Authors: []string{npubHex},
		//Kinds: []int{KindTextNote,KindRepost,KindReaction},
	}}

	clientReqMessage := ClientReqMessage{
		SubscriptionId: uuid.New().String(),
		Filters: filters,
	}

	clientReqJson := clientReqMessage.ToJson()
	fmt.Printf("clientReqJson: %s\n", clientReqJson)

	//conn := Connect("nos.lol")
	conn := Connect("nostr.mom")

	receivedMessage := make(chan string)
	receivedMessagesDone := make(chan error)

	go ReceiveMessages(conn, receivedMessage, receivedMessagesDone)

	err := conn.WriteMessage(websocket.TextMessage, []byte(clientReqJson))
	if err != nil {
		log.Fatal(err)
	}

	numOfMessages := 0
	for {
		label, message := ProcessRelayMessage(<-receivedMessage)
		numOfMessages++

		switch label {
		case "EVENT":
			var eventMessage RelayEventMessage
			err = json.Unmarshal(message[0], &eventMessage.SubscriptionId)
			if err != nil {
				log.Fatal(err)
			}

			err = json.Unmarshal(message[1], &eventMessage.Event)
			if err != nil {
				log.Fatal(err)
			}
			generatedEventId := GenerateEventId(eventMessage.Event)
			if generatedEventId != eventMessage.Event.Id {
				log.Fatal("Incorrect Id received!")
			}
			var eventJson = eventMessage.ToJson()
			fmt.Printf("RelayEventMessage: %s\n", eventJson)

		case "OK":
			var okMessage RelayOkMessage
			err = json.Unmarshal(message[0], &okMessage.EventId)
			if err != nil {
				log.Fatal(err)
			}

			err = json.Unmarshal(message[1], &okMessage.Status)
			if err != nil {
				log.Fatal(err)
			}

			err = json.Unmarshal(message[2], &okMessage.Message)
			if err != nil {
				log.Fatal(err)
			}
			var okJson, _ = json.Marshal(okMessage)
			fmt.Printf("RelayOkMessage: %s\n", okJson)

		case "EOSE":
			var eoseMessage RelayEoseMessage
			err = json.Unmarshal(message[0], &eoseMessage.SubscriptionId)
			if err != nil {
				log.Fatal(err)
			}
			var eoseJson = eoseMessage.ToJson()
			fmt.Printf("RelayEoseMessage: %s\n", eoseJson)
			goto end

		case "NOTICE":
			var noticeMessage RelayNoticeMessage
			err = json.Unmarshal(message[0], &noticeMessage.Message)
			if err != nil {
				log.Fatal(err)
			}
			var noticeJson  = noticeMessage.ToJson()
			fmt.Printf("RelayNoticeMessage: %s\n", noticeJson)
			goto end

		default:
			log.Fatalf("Unknown Relay Message type: \"%s\"", label)
		}
	}
end:
	fmt.Println("NumOfMessages: ", numOfMessages)
	_ = conn.WriteMessage(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	select {
	case err := <-receivedMessagesDone:
		if err != nil {
			log.Fatal("receivedMessages exited with error: ", err)
		}
	case <-time.After(10 * time.Second):
		log.Fatal("recievedMessages didn't close after 10 seconds")
	}
}
