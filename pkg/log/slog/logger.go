package slog

import (
	"log"
	"log/slog"
	"os"
)

const (
	ENV_LOCALE = "local"
	ENV_DEV    = "dev"
	ENV_PROD   = "prod"
)

func NewLogger(env string) (func(), error) {
	var logger *slog.Logger
	logFileCloseAction := func() {}
	switch env {
	case ENV_LOCALE:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case ENV_DEV:
		logFile, err := os.OpenFile("logs.json", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		logFileCloseAction = func() {
			logFile.Close()
		}
		if err != nil {
			log.Fatalf("Cannot find or create log file!")
			return logFileCloseAction, err
		}
		logger = slog.New(slog.NewJSONHandler(logFile, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case ENV_PROD:
		logFile, err := os.OpenFile("logs.json", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		logFileCloseAction = func() {
			logFile.Close()
		}
		if err != nil {
			log.Fatalf("Cannot find or create log file!")
			return logFileCloseAction, err
		}
		logger = slog.New(slog.NewJSONHandler(logFile, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	slog.SetDefault(logger)
	return logFileCloseAction, nil
}
