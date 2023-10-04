package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

var SERVER = "ws.ifelse.io"

var in = bufio.NewReader(os.Stdin)

func main() {
	fmt.Println("Connecting to ", SERVER)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	URL := url.URL{Scheme: "wss", Host: SERVER}
	conn, _, err := websocket.DefaultDialer.Dial(URL.String(), nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	defer fmt.Println("Disconnected!")

	fmt.Println("Connected! Ctrl+C to exit")

	// Process received messages
	receiveMessagesDone := make(chan struct{})
	receiveMessages := func() {
		defer close(receiveMessagesDone)

		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
					return
				}
				log.Println("ReadMessage() error: ", err)
				return
			}
			fmt.Printf("Received: %s\n", message)
			fmt.Print("> ")
		}
	}
	go receiveMessages()

	// Process input
	input := make(chan string, 1)
	getInput := func() {
		fmt.Print("> ")
		result, err := in.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}
		input <- result
	}
	go getInput()

	for {
		select {
		case t := <-input: // on new user input
			err := conn.WriteMessage(websocket.TextMessage, []byte(t))
			if err != nil {
				log.Println("Write error:", err)
				os.Exit(1)
			}
			go getInput()
		case <-interrupt: // on Ctrl+C press
			log.Println("Disconnecting...")
			err := conn.WriteMessage(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("Write close error: ", err)
				os.Exit(1)
			}
			select {
			// wait for receiveMessages to end
			case <-receiveMessagesDone:
				os.Exit(0)
			case <-time.After(10 * time.Second):
				log.Println("receiveMessages failed to finish in 10 seconds")
				os.Exit(1)
			}
		case <-receiveMessagesDone: // on receieveMessages#ReadMessage err
			os.Exit(1)
		}
	}
}
