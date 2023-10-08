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

func (e Event) MarshalJSON() ([]byte, error) {
	tagsJson, err := json.Marshal(e.Tags)
	if err != nil {
		log.Fatal(err)
	}

	//if e.Id == "" { // @MarkFix
	//	e.Id = GenerateEventId(e)
	//}

	return json.Marshal(map[string]interface{}{
		"id":         e.Id,
		"pubkey":     e.PubKey,
		"created_at": e.CreatedAt,
		"kind":       e.Kind,
		"tags":       tagsJson,
		"content":    e.Content,
		"sig":        e.Sig,
	})

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

func DecorateJsonStr(str string) string {
	var result strings.Builder
	result.WriteByte('"')
	for _, c := range []byte(str) {
		escapeByte(&result, c)
	}
	result.WriteByte('"')
	return result.String()
}

func GenerateEventId(event Event) string {
	tagsBytes, _ := event.Tags.MarshalJSON()

	serializedEvent := fmt.Sprintf("[0,\"%s\",%d,%d,%s,%s]",
		event.PubKey,
		event.CreatedAt,
		event.Kind,
		tagsBytes,
		DecorateJsonStr(event.Content))

	hash := sha256.Sum256([]byte(serializedEvent))
	return hex.EncodeToString(hash[:])
}
