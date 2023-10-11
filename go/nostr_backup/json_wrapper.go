package main

import "encoding/json"

type rawJsonArray []json.RawMessage

func IsValidJson(str string) bool {
	return json.Valid([]byte(str))
}

func UnmarshalJSON(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
