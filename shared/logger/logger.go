package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Logger *logrus.Logger

// init - init the logger on startup
func init() {
	// Create a new instance of the logger
	Logger = logrus.New()

	// Set log format (e.g., JSON for structured logging)
	Logger.SetFormatter(&logrus.JSONFormatter{})

	// Set output to stdout
	Logger.SetOutput(os.Stdout)

	// Set the default log level
	Logger.SetLevel(logrus.InfoLevel)
}
