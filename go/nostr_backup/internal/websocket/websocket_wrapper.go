package websocket

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

type WSConnection = *websocket.Conn

func Connect(server string) WSConnection {
	URL := url.URL{Scheme: "wss", Host: server}
	conn, _, err := websocket.DefaultDialer.Dial(URL.String(), nil)
	if err != nil {
		log.Fatal("Failed to connect to websocket", server, err)
	}

	return conn
}

func ReceiveMessages(conn WSConnection, messageChan chan string, doneChan chan error) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				doneChan <- nil
				return
			}
			doneChan <- err
			return
		}

		messageChan <- string(message)
	}
}

func WriteMessage(conn WSConnection, message string) {
	err := conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		log.Fatal("Failed to write websocket message!", err)
	}
}

func WriteCloseMessage(conn WSConnection) {
	err := conn.WriteMessage(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Fatal("Writing websocket close message failed!", err)
	}
}
