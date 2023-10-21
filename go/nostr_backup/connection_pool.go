package main

import "log"

type ConnectionPoolMessage struct {
	Connection *Connection
	Server     string
	Message    string
}

type ConnectionPool struct {
	Connections []Connection
	DoneChans   []*chan error
	MessageChan chan ConnectionPoolMessage
}

func messageAggregator(
	cpMessageChan chan ConnectionPoolMessage, conn *Connection,
	messageChan *chan string, doneChan *chan error) {
	for {
		select {
		case newMessage := <-*messageChan:
			cpMessage := ConnectionPoolMessage{
				Connection: conn,
				Server:     conn.Server,
				Message:    newMessage,
			}
			cpMessageChan <- cpMessage
		case err := <-*doneChan:
			if err != nil {
				log.Fatal("messageAggregator failed", err)
			}
			return
		}
	}
}

func (cp *ConnectionPool) AddConnection(server string) {
	newConn := Connect(server)
	cp.Connections = append(cp.Connections, *newConn)
	doneChan := make(chan error)
	cp.DoneChans = append(cp.DoneChans, &doneChan)

	go messageAggregator(cp.MessageChan, newConn, &newConn.MessageChan, &doneChan)
}

func (cp *ConnectionPool) Close() {
	for i := range cp.Connections {
		*cp.DoneChans[i] <- nil
		cp.Connections[i].Close()
	}
	close(cp.MessageChan)
}

func (cp *ConnectionPool) CloseConnection(server string) {
	for i := range cp.Connections {
		if cp.Connections[i].Server == server {
			*cp.DoneChans[i] <- nil
			cp.Connections[i].Close()

			//assert numConns(Connections) == numConns(DoneChans)

			numConns := len(cp.Connections)
			cp.Connections[i] = cp.Connections[numConns-1]
			cp.DoneChans[i] = cp.DoneChans[numConns-1]
			cp.Connections = cp.Connections[:numConns-1]
			cp.DoneChans = cp.DoneChans[:numConns-1]
		}
	}
	log.Fatal("Cannot close connection", server, "as not in connection pool")
}

func (cp *ConnectionPool) CreateSubscriptions(subscriptionId string, filters Filters) {
	for i := range cp.Connections {
		cp.Connections[i].CreateSubscription(subscriptionId, filters)
		//conn := cp.Connections[i]
		//(&conn).CreateSubscription(subscriptionId, filters)
	}
}

func (cp *ConnectionPool) EoseSubscription(server string, subscriptionId string) {
	for i := range cp.Connections {
		if cp.Connections[i].Server == server {
			cp.Connections[i].EoseSubscription(subscriptionId)
			return
		}
	}
	log.Fatal("EoseSubscription fail! Could not find subscriptionId", subscriptionId, "for server", server)
}

func (cp *ConnectionPool) CloseSubscription(server string, subscriptionId string) {
	for i := range cp.Connections {
		if cp.Connections[i].Server == server {
			cp.Connections[i].CloseSubscription(subscriptionId)
			return
		}
	}
	log.Fatal("CloseSubscription fail! Could not find subscriptionId", subscriptionId, "for server", server)
}

func (cp *ConnectionPool) HasAllSubsEosed() bool {
	for i := range cp.Connections {
		if !cp.Connections[i].HasAllSubsEosed() {
			return false
		}
	}
	return true
}
