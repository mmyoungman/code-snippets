package main

import (
	"fmt"
	"log"
	"mmyoungman/nostr_backup/internal/json"
	"mmyoungman/nostr_backup/internal/uuid"
	"mmyoungman/nostr_backup/internal/websocket"
	"time"
)

func main() {
	//npub := "1f0rwg0z2smrkggypqn7gctscevu22z6thch243365xt0tz8fw9uqupzj2x"
	npubHex := "4bc6e43c4a86c764208104fc8c2e18cb38a50b4bbe2eaac63aa196f588e97178"

	//event := Event{
	//	PubKey:    npubHex,
	//	Kind:      KindTextNote,
	//	CreatedAt: 0,
	//	Tags:      make([]Tag, 0),
	//	Content:   "Test!\n❤️‍🔥\"b\\😅  <html>",
	//}
	//event.Id = event.GenerateEventId()

	//eventJson := event.ToJson()

	//fmt.Printf("Event JSON: %s\n", eventJson)

	//var eventStruct Event
	//_ = json.UnmarshalJSON([]byte(eventJson), &eventStruct)

	//fmt.Println(
	//	"eventStruct: ",
	//	eventStruct.Id,
	//	eventStruct.PubKey,
	//	eventStruct.CreatedAt,
	//	eventStruct.Kind,
	//	eventStruct.Tags,
	//	eventStruct.Content,
	//	eventStruct.Sig)

	db := DBConnect()
	defer db.Close()

	//DBInsertEvent(db, eventStruct)

	//events := DBGetEvents(db)

	//for _, event := range events {
	//	fmt.Println(event.ToJson())
	//}

	filters := Filters{{
		Authors: []string{npubHex},
		//Kinds: []int{KindTextNote,KindRepost,KindReaction},
	}}

	clientReqMessage := ClientReqMessage{
		SubscriptionId: uuid.NewUuid(),
		Filters:        filters,
	}

	clientReqJson := clientReqMessage.ToJson()
	fmt.Printf("clientReqJson: %s\n", clientReqJson)

	//conn := websocket.WSConnect("nos.lol")
	conn := websocket.Connect("nostr.mom")
	defer conn.Close()

	receivedMessage := make(chan string)
	receivedMessagesDone := make(chan error)

	go websocket.ReceiveMessages(conn, receivedMessage, receivedMessagesDone)

	websocket.WriteMessage(conn, clientReqJson)

	numOfMessages := 0
	for {
		var newMessage string
		select {
		case newMessage = <-receivedMessage:
		case <-time.After(10 * time.Second):
			fmt.Println("No new message received in 10 seconds")
			goto end
		}
		//newMessage := <-receivedMessage
		label, message := ProcessRelayMessage(newMessage)
		numOfMessages++

		switch label {
		case "EVENT":
			var eventMessage RelayEventMessage
			err := json.UnmarshalJSON(message[0], &eventMessage.SubscriptionId)
			if err != nil {
				log.Fatal("Failed to unmarshal RelayEventMessage.SubscriptionId", err)
			}

			err = json.UnmarshalJSON(message[1], &eventMessage.Event)
			if err != nil {
				log.Fatal("Failed to unmarshal RelayEventMessage.Event", err)
			}
			generatedEventId := eventMessage.Event.GenerateEventId()
			if generatedEventId != eventMessage.Event.Id {
				log.Fatal("Incorrect Id received!")
			}

			eventHasValidSig := eventMessage.Event.IsSigValid()
			if !eventHasValidSig {
				log.Fatal("Event has invalid sig: ",
					eventMessage.Event.ToJson())
			}

			DBInsertEvent(db, eventMessage.Event)

			//eventJson := eventMessage.ToJson()
			//fmt.Printf("RelayEventMessage: %s\n", eventJson)

		case "OK":
			var okMessage RelayOkMessage
			err := json.UnmarshalJSON(message[0], &okMessage.EventId)
			if err != nil {
				log.Fatal("Failed to unmarshal RelayOkMessage.EventId", err)
			}

			err = json.UnmarshalJSON(message[1], &okMessage.Status)
			if err != nil {
				log.Fatal("Failed to unmarshal RelayOkMessage.Status", err)
			}

			err = json.UnmarshalJSON(message[2], &okMessage.Message)
			if err != nil {
				log.Fatal("Failed to unmarshal RelayOkMessage.Message", err)
			}
			okJson := okMessage.ToJson()
			fmt.Printf("RelayOkMessage: %s\n", okJson)

		case "EOSE":
			var eoseMessage RelayEoseMessage
			err := json.UnmarshalJSON(message[0], &eoseMessage.SubscriptionId)
			if err != nil {
				log.Fatal("Failed to unmarshal RelayEoseMessage.SubscriptionId", err)
			}
			eoseJson := eoseMessage.ToJson()
			fmt.Printf("RelayEoseMessage: %s\n", eoseJson)
			goto end

		case "NOTICE":
			var noticeMessage RelayNoticeMessage
			err := json.UnmarshalJSON(message[0], &noticeMessage.Message)
			if err != nil {
				log.Fatal("Failed to unmarshal RelayNoticeMessage.Message", err)
			}
			noticeJson := noticeMessage.ToJson()
			fmt.Printf("RelayNoticeMessage: %s\n", noticeJson)
			goto end

		default:
			log.Fatalf("Unknown Relay Message type: \"%s\"", label)
		}
	}
end:
	fmt.Println("NumOfMessages: ", numOfMessages)
	websocket.WriteCloseMessage(conn)

	//events := DBGetEvents(db)
	//for _, event := range events {
	//	fmt.Println(event.ToJson())
	//}

	select {
	case err := <-receivedMessagesDone:
		if err != nil {
			log.Fatal("receivedMessages exited with error: ", err)
		}
	case <-time.After(10 * time.Second):
		log.Fatal("recievedMessages didn't close after 10 seconds")
	}
}
