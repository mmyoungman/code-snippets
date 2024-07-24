//go:build !prod

package utils

import "log"

var IsProd = false

func UNUSED(x ...interface{}) {}

func Assert(condition bool) {
	if !condition {
		log.Fatal("Assert failed")
	}
}
