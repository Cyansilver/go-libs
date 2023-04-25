package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

type StandardLogger struct {
	*logrus.Logger
}

// NewLogger initializes the standard logger
func NewLogger() *StandardLogger {
	var baseLogger = logrus.New()

	var standardLogger = &StandardLogger{baseLogger}

	standardLogger.Formatter = &logrus.JSONFormatter{}

	standardLogger.SetReportCaller(true)
	standardLogger.SetOutput(os.Stdout)
	standardLogger.SetLevel(logrus.TraceLevel)

	return standardLogger
}

var Logger = NewLogger()
