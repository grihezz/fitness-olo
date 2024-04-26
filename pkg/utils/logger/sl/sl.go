package sl

import (
	"log/slog"
)

// Err creates a slog attribute for logging errors.
func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
