//go:build !prod

package main

import "mmyoungman/nostr_backup/json_wrapper"

func UNUSED(x ...interface{}) {}

func DevBuildValidJson(str string) {
	if !json_wrapper.IsValidJson(str) {
		panic("Json is not valid!")
	}
}
