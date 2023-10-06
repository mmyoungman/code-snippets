package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Message struct {
	One string
	Two string
	Three int64
}

func main() {
	m := Message{"Alice", "Hello", 182947947914}
	jsonBytes, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("json string: ", string(jsonBytes))

	var msgStruct Message
	err = json.Unmarshal(jsonBytes, &msgStruct)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("msgStruct: .Name, .Body, .Time: ", msgStruct.One, msgStruct.Two, msgStruct.Three)


}
