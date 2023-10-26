package main

import "log"

type ConnectionListMessage struct {
	Connection *Connection
	Message    string
}

type ConnectionList struct {
	Connections [](*Connection)
	DoneChans   []chan error
	MessageChan chan ConnectionListMessage
}

func CreateConnectionList() *ConnectionList {
	var connList ConnectionList
	connList.MessageChan = make(chan ConnectionListMessage, 100)
	return &connList
}

func messageAggregator(cpMessageChan chan ConnectionListMessage,
	conn *Connection, doneChan chan error) {
	for {
		select {
		case newMessage := <-conn.MessageChan:
			cpMessage := ConnectionListMessage{
				Connection: conn,
				Message:    newMessage,
			}
			cpMessageChan <- cpMessage
		case err := <-doneChan:
			if err != nil {
				log.Fatal("messageAggregator failed", err)
			}
			return
		}
	}
}

func (cp *ConnectionList) AddConnection(server string) {
	newConn := Connect(server)
	cp.Connections = append(cp.Connections, newConn)
	doneChan := make(chan error)
	cp.DoneChans = append(cp.DoneChans, doneChan)

	//assert len(Connections) == len(DoneChans)

	go messageAggregator(cp.MessageChan, newConn, doneChan)
}

func (cp *ConnectionList) Close() {
	for i := range cp.Connections {
		cp.DoneChans[i] <- nil
		cp.Connections[i].Close()
	}
	close(cp.MessageChan)
}

func (cp *ConnectionList) CloseConnection(server string) {
	for i := range cp.Connections {
		if cp.Connections[i].Server == server {
			cp.DoneChans[i] <- nil
			cp.Connections[i].Close()

			//assert len(Connections) == len(DoneChans)

			// remove connection from connList arrays
			numConns := len(cp.Connections)
			cp.Connections[i] = cp.Connections[numConns-1]
			cp.DoneChans[i] = cp.DoneChans[numConns-1]
			cp.Connections = cp.Connections[:numConns-1]
			cp.DoneChans = cp.DoneChans[:numConns-1]
			return
		}
	}
	log.Fatal("Cannot close connection", server, " as not in connection list")
}

func (cp *ConnectionList) CreateSubscriptions(subscriptionId string, filters Filters) {
	for i := range cp.Connections {
		cp.Connections[i].CreateSubscription(subscriptionId, filters)
	}
}

func (cp *ConnectionList) CloseSubscription(server string, subscriptionId string) {
	for i := range cp.Connections {
		if cp.Connections[i].Server == server {
			cp.Connections[i].CloseSubscription(subscriptionId)
			return
		}
	}
	log.Fatal("CloseSubscription fail! Could not find subscriptionId", subscriptionId, "for server", server)
}

func (cp *ConnectionList) HasAllSubsEosed() bool {
	for i := range cp.Connections {
		if !cp.Connections[i].HasAllSubsEosed() {
			return false
		}
	}
	return true
}
