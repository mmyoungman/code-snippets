package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
)

var SERVER = "ws.ifelse.io"
var PORT = ""
var PATH = ""

var TIMESWAIT = 0
var TIMESWAITMAX = 5

var in = bufio.NewReader(os.Stdin)

func getInput(input chan string) {
	result, err := in.ReadString('\n')
	if err != nil {
		log.Println(err)
		return
	}
	input <- result
}

func main() {
	fmt.Println("Connecting to:", SERVER, "at", PATH)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	input := make(chan string, 1)
	go getInput(input)

	URL := url.URL{Scheme: "wss", Host: SERVER, Path: PATH}
	fmt.Println("url:", URL.String())
	conn, _, err := websocket.DefaultDialer.Dial(URL.String(), nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				closeErr := err.(*websocket.CloseError)
				if closeErr.Code == websocket.CloseNormalClosure {
					return
				}
				log.Println("ReadMessage() error: ", err)
				return
			}
			fmt.Printf("Received: %s\n", message)
		}
	}()

	for {
		select {
		case <-time.After(4 * time.Second):
			log.Println("Please give me input!", TIMESWAIT)
			TIMESWAIT++
			if TIMESWAIT > TIMESWAITMAX {
				syscall.Kill(syscall.Getpid(), syscall.SIGINT)
			}
		case <-done:
			return
		case t := <-input:
			err := conn.WriteMessage(websocket.TextMessage, []byte(t))
			if err != nil {
				log.Println("Write error:", err)
				return
			}
			TIMESWAIT = 0
			go getInput(input)
		case <-interrupt:
			log.Println("Disconnecting...")
			err := conn.WriteMessage(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("Write close error: ", err)
				return
			}
			select {
			case <-done:
			case <-time.After(2 * time.Second):
			}
			return
		}
	}
}
