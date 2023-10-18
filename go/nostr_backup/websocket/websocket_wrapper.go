package websocket

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

func WSConnect(server string) *websocket.Conn {
	URL := url.URL{Scheme: "wss", Host: server}
	conn, _, err := websocket.DefaultDialer.Dial(URL.String(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return conn
}

func WSReceieveMessages(conn *websocket.Conn, receivedMessage chan string, done chan error) {
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

func WSWriteMessage(conn *websocket.Conn, message string) {
	err := conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		log.Fatal("Failed to write websocket message!", err)
	}
}

func WSSendCloseMessage(conn *websocket.Conn) {
	err := conn.WriteMessage(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Fatal("Writing websocket close message failed!", err)
	}
}
