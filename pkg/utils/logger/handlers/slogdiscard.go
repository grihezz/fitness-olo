package slogpretty

import (
	"context"
	"golang.org/x/exp/slog"
)

// DiscardHandler a Handler handles log records produced by a Logger..
type DiscardHandler struct{}

// Handle handles the Record.
func (h *DiscardHandler) Handle(_ context.Context, _ slog.Record) error {
	// Просто игнорируем запись журнала
	return nil
}

// WithAttrs returns a new Handler whose attributes consist of
// both the receiver's attributes and the arguments.
func (h *DiscardHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	// Возвращает тот же обработчик, так как нет атрибутов для сохранения
	return h
}

// WithGroup returns a new Handler with the given group appended to
// the receiver's existing groups.
func (h *DiscardHandler) WithGroup(_ string) slog.Handler {
	// Возвращает тот же обработчик, так как нет группы для сохранения
	return h
}

// Enabled reports whether the handler handles records at the given level.
// The handler ignores records whose level is lower.
func (h *DiscardHandler) Enabled(_ context.Context, _ slog.Level) bool {
	// Всегда возвращает false, так как запись журнала игнорируется
	return false
}
