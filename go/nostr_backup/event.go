package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
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

func (event Event) MarshalJSON() ([]byte, error) {
	panic("Use ToJson")
}

func (event Event) ToJson() string {
	var result strings.Builder
	result.WriteString(fmt.Sprintf("{\"id\":\"%s\",", event.Id))
	result.WriteString(fmt.Sprintf("\"pubkey\":\"%s\",", event.PubKey))
	result.WriteString(fmt.Sprintf("\"created_at\":%d,", event.CreatedAt))
	result.WriteString(fmt.Sprintf("\"kind\":%d,", event.Kind))
	result.WriteString(fmt.Sprintf("\"tags\":%s,", event.Tags.ToJson()))
	result.WriteString(
		fmt.Sprintf("\"content\":%s,", DecorateJsonStr(event.Content)))
	result.WriteString(fmt.Sprintf("\"sig\":\"%s\"}", event.Sig))

	// @MarkFix Dev build only check that JSON is valid (along with all other ToJson
	// functions

	return result.String()
}

func GenerateEventId(event Event) string {
	serializedEvent := fmt.Sprintf("[0,\"%s\",%d,%d,%s,%s]",
		event.PubKey,
		event.CreatedAt,
		event.Kind,
		event.Tags.ToJson(),
		DecorateJsonStr(event.Content))

	hash := sha256.Sum256([]byte(serializedEvent))
	return hex.EncodeToString(hash[:])
}
