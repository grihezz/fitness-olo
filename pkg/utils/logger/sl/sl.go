package sl

import (
	"fmt"
	"log/slog"
)

// Err creates a slog attribute for logging errors.
func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}

func Wrap(op string, err error) error {
	return fmt.Errorf("%s : %w", op, err)
}
