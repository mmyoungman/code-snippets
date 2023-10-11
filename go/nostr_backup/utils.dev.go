//go:build !prod

package main

func UNUSED(x ...interface{}) {}

func DevBuildValidJson(str string) {
	if !IsValidJson(str) {
		panic("Json is not valid!")
	}
}
