package logger

import (
	"log/slog"
	"os"
)

func New(isProduction bool) *slog.Logger {
	var handler slog.Handler
	if isProduction {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	} else {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	}
	return slog.New(handler).With("service", "medcord-backend")
}
