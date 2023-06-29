package utils

import (
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	once   sync.Once
	logger *logrus.Logger
)

// GetLogger returns the singleton logger instance.
func GetLogger() *logrus.Logger {
	once.Do(func() {
		logger = createLogger()
	})
	return logger
}

func createLogger() *logrus.Logger {
	logFile, err := os.OpenFile("logs.json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatal("Failed to open log file: ", err)
	}

	logger := logrus.New()
	logger.SetOutput(logFile)
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.DebugLevel)

	logger.Info("Logger initialized")

	return logger
}

type LogFormat struct {
	Package  string `json:"package"`
	OnAction string `json:"on_action"`
	Message  string `json:"message"`
}
