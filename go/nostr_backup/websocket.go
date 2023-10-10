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

func ReceiveMessages(conn *websocket.Conn, receivedMessage chan string, done chan error) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				done <- nil
				return
			}
			log.Println("ReceiveMessages() error: ", err)
			done <- err
			return
		}

		receivedMessage <- string(message)
	}
}
