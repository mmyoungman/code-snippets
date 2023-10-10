//go:build !prod

package main

import "encoding/json"

func UNUSED(x ...interface{}) {}

func DevBuildValidJson(str string) {
	if !json.Valid([]byte(str)) {
		panic("Json is not valid!")
	}
}
