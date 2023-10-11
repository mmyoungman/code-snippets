package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
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

func (event Event) ToJson() string {
	result := fmt.Sprintf(
		"{\"id\":\"%s\",\"pubkey\":\"%s\",\"created_at\":%d,\"kind\":%d,\"tags\":%s,\"content\":%s,\"sig\":\"%s\"}",
		event.Id,
		event.PubKey,
		event.CreatedAt,
		event.Kind,
		event.Tags.ToJson(),
		DecorateJsonStr(event.Content),
		event.Sig)

	DevBuildValidJson(result)

	return result
}

func (event Event) GenerateEventId() string {
	serializedEvent := fmt.Sprintf("[0,\"%s\",%d,%d,%s,%s]",
		event.PubKey,
		event.CreatedAt,
		event.Kind,
		event.Tags.ToJson(),
		DecorateJsonStr(event.Content))

	DevBuildValidJson(serializedEvent)

	hash := sha256.Sum256([]byte(serializedEvent))
	return hex.EncodeToString(hash[:])
}
