package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func main() {
	event := Event{
		PubKey:    "",
		Kind:      KindTextNote,
		CreatedAt: 0,
		Tags:      make([]Tag, 0),
		Content:   "Test!\n❤️‍🔥\"b\\😅",
	}
	event.Id = GenerateEventId(event)

	eventJson := event.String()

	fmt.Println("Event JSON: ", eventJson)

	eventStruct := JsonToEvent(eventJson)

	fmt.Println(
		"eventStruct: ",
		eventStruct.Id,
		eventStruct.PubKey,
		eventStruct.CreatedAt,
		eventStruct.Kind,
		eventStruct.Tags,
		eventStruct.Content,
		eventStruct.Sig)

	//npub := "1f0rwg0z2smrkggypqn7gctscevu22z6thch243365xt0tz8fw9uqupzj2x"
	npubHex := "4bc6e43c4a86c764208104fc8c2e18cb38a50b4bbe2eaac63aa196f588e97178"

	filter := Filter{
		Authors: []string{npubHex},
		Kinds: []int{KindTextNote,KindRepost,KindReaction},
	}

	subscriptionId := uuid.New()
	message := fmt.Sprintf("[\"REQ\", \"%s\", %s]", subscriptionId, filter.String())

	fmt.Println("Message: ", message)

	//conn := Connect("nos.lol")
	conn := Connect("nostr.mom")

	receivedMessages := make(chan string)

	go ReceiveMessages(conn, receivedMessages)
	defer conn.Close()

	err := conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		log.Fatal(err)
	}

	for {
		newMessage := <-receivedMessages
		if strings.HasPrefix(newMessage, "[\"EVENT\",") {
			eventSubId, event := JsonToEventMessage(newMessage)
			if eventSubId != subscriptionId.String() {
				log.Fatal("Event subscriptionId incorrect?")
			}

			fmt.Printf("Processed EVENT: \n%s\n", event.String())
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

		//messageStr := string(message)
		//switch {
		//case strings.HasPrefix(messageStr, "[\"EVENT\","):
		//	fmt.Println("Got an EVENT message!")
		//case strings.HasPrefix(messageStr, "[\"OK\","):
		//	fmt.Println("Got an OK message!")
		//case strings.HasPrefix(messageStr, "[\"EOSE\","):
		//	fmt.Println("Got an EOSE message!")
		//case strings.HasPrefix(messageStr, "[\"NOTICE\","):
		//	fmt.Println("Got an NOTICE message!")
		//default:
		//	fmt.Println("Received message of unknown type!")
		//}
