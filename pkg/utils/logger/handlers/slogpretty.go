// Package slogpretty provides logging utilities with pretty formatting for the slog package.
//
// This package includes implementations for creating pretty loggers and handlers.
package slogpretty

import (
	"context"
	"encoding/json"
	"io"
	stdLog "log"
	"log/slog"

	"github.com/fatih/color"
)

// PrettyHandlerOptions contains options for configuring the PrettyHandler.
type PrettyHandlerOptions struct {
	SlogOpts *slog.HandlerOptions // Slog handler options
}

// PrettyHandler represents a handler that formats log entries in a pretty way.
type PrettyHandler struct {
	opts         PrettyHandlerOptions // Pretty handler options
	slog.Handler                      // Underlying slog handler
	l            *stdLog.Logger       // Standard logger
	attrs        []slog.Attr          // Attributes
}

// NewPrettyHandler creates a new instance of PrettyHandler with the provided options and output writer.
func (opts PrettyHandlerOptions) NewPrettyHandler(out io.Writer) *PrettyHandler {
	h := &PrettyHandler{
		Handler: slog.NewJSONHandler(out, opts.SlogOpts),
		l:       stdLog.New(out, "", 0),
	}

	return h
}

// Handle implements the Handle method of the slog.Handler interface.
// It formats and prints the log entry in a pretty way.
func (h *PrettyHandler) Handle(_ context.Context, r slog.Record) error {
	level := r.Level.String() + ":"

	switch r.Level {
	case slog.LevelDebug:
		level = color.MagentaString(level)
	case slog.LevelInfo:
		level = color.BlueString(level)
	case slog.LevelWarn:
		level = color.YellowString(level)
	case slog.LevelError:
		level = color.RedString(level)
	}

	fields := make(map[string]interface{}, r.NumAttrs())

	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()

		return true
	})

	for _, a := range h.attrs {
		fields[a.Key] = a.Value.Any()
	}

	var b []byte
	var err error

	if len(fields) > 0 {
		b, err = json.MarshalIndent(fields, "", "  ")
		if err != nil {
			return err
		}
	}

	timeStr := r.Time.Format("[15:05:05.000]")
	msg := color.CyanString(r.Message)

	h.l.Println(
		timeStr,
		level,
		msg,
		color.WhiteString(string(b)),
	)

	return nil
}

// WithAttrs implements the WithAttrs method of the slog.Handler interface.
// It returns a new PrettyHandler with the given attributes.
func (h *PrettyHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &PrettyHandler{
		Handler: h.Handler,
		l:       h.l,
		attrs:   attrs,
	}
}

// WithGroup implements the WithGroup method of the slog.Handler interface.
// It returns a new PrettyHandler with the given group.
func (h *PrettyHandler) WithGroup(name string) slog.Handler {
	// TODO: implement
	return &PrettyHandler{
		Handler: h.Handler.WithGroup(name),
		l:       h.l,
	}
}
