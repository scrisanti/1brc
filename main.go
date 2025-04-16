package main

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/user"
	"path/filepath"
)

func main() {

	// Get User Home Dir
	usr, err := user.Current()
	if err != nil {
		slog.Error(fmt.Sprintf("Couldn't Get User: %v", err))
	}
	home := usr.HomeDir

	// Log To Standard Dir
	logDir := filepath.Join(home, "logs")
	logFilename := "log_1brc.log"
	logFP := filepath.Join(logDir, logFilename)
	logFile, err := os.OpenFile(logFP, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	mw := io.MultiWriter(os.Stdout, logFile)
	handler := slog.NewTextHandler(mw, &slog.HandlerOptions{Level: slog.LevelDebug})
	logger := slog.New(handler)
	slog.SetDefault(logger)

	slog.Info("# ----- 1brc ----- # ")
	cwd, err := os.Getwd()
	if err != nil {
		slog.Error(fmt.Sprintf("Error Getting CWD: %v", err))
	}
	slog.Debug("User CWD:", "cwd", cwd)
}
