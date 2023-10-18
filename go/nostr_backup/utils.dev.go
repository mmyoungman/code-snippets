//go:build !prod

package main

import "mmyoungman/nostr_backup/json"

func UNUSED(x ...interface{}) {}

func DevBuildValidJson(str string) {
	if !json.IsValidJson(str) {
		panic("Json is not valid!")
	}
}
