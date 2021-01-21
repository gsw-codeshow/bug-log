package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {
	url := "127.0.0.1"
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
	return
}
