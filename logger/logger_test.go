package logger

import (
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestLogger(t *testing.T) {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	sugar := logger.Sugar()
	url := "https://dd.com"
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)

	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

func TestWriter(t *testing.T) {
	l := NewLogger(LOG_OUT_ALL, "test")
	l.Debug("debug...")
	l.Debugf("debugf %s ...", "param")
	l.Info("Info...")
	l.Infof("Infof %s ...", "param")
	l.Error("Error...")
	l.Errorf("Errorf %s ...", "param")
}
