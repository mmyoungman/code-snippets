package main

import (
	"log"
	"mmyoungman/nostr_backup/internal/uuid"
	"mmyoungman/nostr_backup/internal/websocket"
	"time"
)

type Connection struct {
	Server        string
	WSConnection  websocket.Connection
	Subscriptions []Subscription
	MessageChan   chan string
	DoneChan      chan error
}

func Connect(server string) *Connection {
	var conn Connection
	conn.WSConnection = websocket.Connect(server)
	conn.MessageChan = make(chan string)
	conn.DoneChan = make(chan error)

	go websocket.ReceiveMessages(conn.WSConnection,
		conn.MessageChan, conn.DoneChan)

	return &conn
}

func (conn *Connection) Close() {
	websocket.WriteCloseMessage(conn.WSConnection)
	select {
	case err := <-conn.DoneChan:
		if err != nil {
			log.Fatal("receivedMessages exited with error: ", err)
		}
	case <-time.After(10 * time.Second):
		log.Fatal("recievedMessages didn't close after 10 seconds")
	}
	close(conn.MessageChan)
	close(conn.DoneChan)
}

func (conn *Connection) CreateSubscription(filters Filters) (subscriptionId string) {
	var subscription Subscription
	subscription.Id = uuid.NewUuid()
	subscription.Filters = filters

	clientReqMessage := ClientReqMessage{
		SubscriptionId: subscription.Id,
		Filters:        filters,
	}
	websocket.WriteMessage(
		conn.WSConnection, clientReqMessage.ToJson())

	conn.Subscriptions = append(
		conn.Subscriptions, subscription)

	return subscriptionId
}

func (conn *Connection) HasAllSubsEosed() bool {
	for _, sub := range conn.Subscriptions {
		if !sub.Eose {
			return false
		}
	}
	return true
}

func (conn *Connection) EoseSubscription(subscriptionId string) {
	for i := range conn.Subscriptions {
		if conn.Subscriptions[i].Id == subscriptionId {
			conn.Subscriptions[i].Eose = true
		}
	}

}

func (conn *Connection) CloseSubscription(subscriptionId string) {
	for i := range conn.Subscriptions {
		if conn.Subscriptions[i].Id == subscriptionId {
			conn.Subscriptions[i] = conn.Subscriptions[len(conn.Subscriptions)-1]
			conn.Subscriptions = conn.Subscriptions[:len(conn.Subscriptions)-1]
			goto closeWSConnection
		}
	}
	log.Fatal("Subscription not found", subscriptionId, "for connection", conn.Server)

closeWSConnection:
	clientCloseMessage := ClientCloseMessage{
		SubscriptionId: subscriptionId,
	}
	websocket.WriteMessage(conn.WSConnection, clientCloseMessage.ToJson())
}
