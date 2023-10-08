package main

import "fmt"

type Tag []string
type Tags []Tag

func (tags Tags) String() string { // @MarkFix untested
	result := "["

	for i, tag := range tags {
		if i > 0 {
			result += ","
		}

		result += tag.String()
	}

	return result + "]"
}

func (tag Tag) String() string {
	result := "["

	for i, str := range tag {
		if i > 0 {
			result += ","
		}
		result += fmt.Sprintf("\"%s\"", str)
	}

	return result + "]"
}

