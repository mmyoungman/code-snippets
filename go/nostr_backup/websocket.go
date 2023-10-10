package main

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

func Connect(server string) *websocket.Conn {
	URL := url.URL{Scheme: "wss", Host: server}
	conn, _, err := websocket.DefaultDialer.Dial(URL.String(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return conn
}

func ReceiveMessages(conn *websocket.Conn, receivedMessage chan string, done chan bool) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				done <- true
				return
			}
			log.Println("ReceiveMessages() error: ", err)
			return
		}

		receivedMessage <- string(message)
	}
}
