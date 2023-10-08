package main

import (
	"encoding/json"
	"strings"
)

type Filter struct {
	Ids     []string `json:"ids,omitempty"`
	Kinds   []int    `json:"kinds,omitempty"`
	Authors []string `json:"authors,omitempty"`
	Tags    []string `json:"-,omitempty"`
	Since   []string `json:"since,omitempty"`
	Until   []string `json:"until,omitempty"`
	Limit   []string `json:"limit,omitempty"`
}

type Filters []Filter

func (filters Filters) MarshalJSON() ([]byte, error) {
	var result strings.Builder

	for i, filter := range filters {
		if i > 1 {
			result.WriteString(",")
		}
		jsonValue, _ := json.Marshal(filter)
		result.WriteString(string(jsonValue))
	}

	return []byte(result.String()), nil
}
