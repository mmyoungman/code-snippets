package main

import (
	"encoding/json"
	"log"
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

func (filter Filter) String() string {
	json, err := json.Marshal(filter)
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

type Filters []Filter

func (filters Filters) String() string {
	json, err := json.Marshal(filters)
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
