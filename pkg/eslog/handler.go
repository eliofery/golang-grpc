package eslog

import "log/slog"

// AttrFn ...
type AttrFn func(groups []string, attr slog.Attr) slog.Attr

// Handler ...
type Handler interface {
	slog.Handler
}
