//go:build prod

package utils

import (
	"log/slog"
)

var IsProd = true

// No declaration of UNUSED func here, as we don't want it used in prod builds

func Assert(condition bool) {
	if !condition {
		slog.Error("Assert failed") // @MarkFix What to do here?
	}
}
