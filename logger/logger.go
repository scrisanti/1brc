package logger

import (
	"context"
	"log/slog"
	"os"
)

type SplitLevelHandler struct {
	infoHandler  slog.Handler
	debugHandler slog.Handler
}

func (h *SplitLevelHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.infoHandler.Enabled(ctx, level) || h.debugHandler.Enabled(ctx, level)
}

func (h *SplitLevelHandler) Handle(ctx context.Context, record slog.Record) error {
	if h.debugHandler.Enabled(ctx, record.Level) {
		if err := h.debugHandler.Handle(ctx, record); err != nil {
			return err
		}
	}
	if h.infoHandler.Enabled(ctx, record.Level) {
		if err := h.infoHandler.Handle(ctx, record); err != nil {
			return err
		}
	}
	return nil
}

func (h *SplitLevelHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &SplitLevelHandler{
		infoHandler:  h.infoHandler.WithAttrs(attrs),
		debugHandler: h.debugHandler.WithAttrs(attrs),
	}
}

func (h *SplitLevelHandler) WithGroup(name string) slog.Handler {
	return &SplitLevelHandler{
		infoHandler:  h.infoHandler.WithGroup(name),
		debugHandler: h.debugHandler.WithGroup(name),
	}
}

// InitLogger sets the default slog logger with stdout+file split behavior
func InitLogger(debugFilePath string) error {
	logFile, err := os.OpenFile(debugFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	infoHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	debugHandler := slog.NewTextHandler(logFile, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	splitHandler := &SplitLevelHandler{
		infoHandler:  infoHandler,
		debugHandler: debugHandler,
	}

	slog.SetDefault(slog.New(splitHandler))
	return nil
}
