package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func main() {
	//npub := "1f0rwg0z2smrkggypqn7gctscevu22z6thch243365xt0tz8fw9uqupzj2x"
	npubHex := "4bc6e43c4a86c764208104fc8c2e18cb38a50b4bbe2eaac63aa196f588e97178"

	//event := Event{
	//	PubKey:    npubHex,
	//	Kind:      KindTextNote,
	//	CreatedAt: 0,
	//	Tags:      make([]Tag, 0),
	//	Content:   "Test!\n❤️‍🔥\"b\\😅",
	//}
	//event.Id = GenerateEventId(event)

	//eventJson := event.String()

	//fmt.Println("Event JSON: ", eventJson)

	//eventStruct := JsonToEvent(eventJson)

	//fmt.Println(
	//	"eventStruct: ",
	//	eventStruct.Id,
	//	eventStruct.PubKey,
	//	eventStruct.CreatedAt,
	//	eventStruct.Kind,
	//	eventStruct.Tags,
	//	eventStruct.Content,
	//	eventStruct.Sig)

	filter := Filter{
		Authors: []string{npubHex},
		//Kinds: []int{KindTextNote,KindRepost,KindReaction},
	}

	subscriptionId := uuid.New().String()
	filterJson, _ := json.Marshal(filter)
	reqMessage := fmt.Sprintf("[\"REQ\", \"%s\", %s]", subscriptionId, filterJson)

	//conn := Connect("nos.lol")
	conn := Connect("nostr.mom")

	receivedMessage := make(chan string)

	go ReceiveMessages(conn, receivedMessage)
	defer conn.Close()

	err := conn.WriteMessage(websocket.TextMessage, []byte(reqMessage))
	if err != nil {
		log.Fatal(err)
	}

	for {
		newMessage := <-receivedMessage
		// @MarkFix make generic message handler?
		if strings.HasPrefix(newMessage, "[\"EVENT\",") {
			fmt.Printf("Received: \n%s\n", newMessage)
			eventSubId, event := JsonToEventMessage(newMessage)
			if eventSubId != subscriptionId {
				log.Fatal("Event subscriptionId incorrect?")
			}

			eventJson, _ := json.Marshal(event)
			fmt.Printf("Processed EVENT: \n%s\n", eventJson)
		}
		if strings.HasPrefix(newMessage, "[\"EOSE\",") {
			break
		}
	}

	fmt.Println("Press Enter to quit")
	in := bufio.NewReader(os.Stdin)
	_, err = in.ReadString('\n')
	if err != nil {
		log.Println(err)
		return
	}
}
