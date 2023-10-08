package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type Event struct {
	Id        string `json:"id"`
	PubKey    string `json:"pubkey"`
	CreatedAt int64  `json:"created_at"`
	Kind      int    `json:"kind"`
	Tags      Tags   `json:"tags"`
	Content   string `json:"content"`
	Sig       string `json:"sig"`
}

func (e Event) String() string {
	eventJson, err := json.Marshal(e)
	if err != nil {
		log.Fatal(err)
	}
	return string(eventJson)

}

func JsonToEvent(eventJson string) Event {
	var event Event
	err := json.Unmarshal([]byte(eventJson), &event)
	if err != nil {
		log.Fatal(err)
	}
	return event
}


var asciiEscapes = []byte{'\\', '"', 'b', 'f', 'n', 'r', 't'}
var binaryEscapes = []byte{'\\', '"', '\b', '\f', '\n', '\r', '\t'}

func escapeByte(b *strings.Builder, c byte) {
	for i, esc := range binaryEscapes {
		if esc == c {
			b.WriteByte('\\')
			b.WriteByte(asciiEscapes[i])
			return
		}
	}
	if c < 0x20 {
		b.WriteString(fmt.Sprintf("\\u%04x", c))
		return
	}
	b.WriteByte(c)
}

func DecorateJsonStr(str string) string { // @MarkFix untested
	var result strings.Builder
	result.WriteByte('"')
	for _, c := range []byte(str) {
		escapeByte(&result, c)
	}
	result.WriteByte('"')
	return result.String()
}

func GenerateEventId(event Event) string {
	serializedEvent := fmt.Sprintf("[0,\"%s\",%d,%d,%s,%s]",
		event.PubKey,
		event.CreatedAt,
		event.Kind,
		event.Tags,
		DecorateJsonStr(event.Content))

	hash := sha256.Sum256([]byte(serializedEvent))
	return hex.EncodeToString(hash[:])
}

//func GenerateEventIdJsonPackageEscaping(event Event) string {
//	content, err := json.Marshal(event.Content)
//	if err != nil {
//		log.Fatal(err)
//	}
//	serializedEvent := fmt.Sprintf("[0,\"%s\",%d,%d,%s,%s]",
//		event.PubKey,
//		event.CreatedAt,
//		event.Kind,
//		event.Tags,
//		content)
//
//	fmt.Printf("Serialized event: %s\n", serializedEvent)
//
//	hash := sha256.Sum256([]byte(serializedEvent))
//	return hex.EncodeToString(hash[:])
//}

