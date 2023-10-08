package main

import (
	"fmt"
	"strings"
)

type Tag []string
type Tags []Tag

func (tags Tags) String() string { // @MarkFix untested
	var result strings.Builder
	result.WriteString("[")

	for i, tag := range tags {
		if i > 0 {
			result.WriteString(",")
		}

		writeTag(&result, tag)
	}

	result.WriteString("]")
	return result.String()
}

func writeTag(result *strings.Builder, tag Tag) {
	result.WriteString("[")

	for i, str := range tag {
		if i > 0 {
			result.WriteString(",")
		}
		result.WriteString(fmt.Sprintf("\"%s\"", str))
	}

	result.WriteString("]")
}
