package logs

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func InitLogger() {
	logger = logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// Set log level from environment or default to info
	level := os.Getenv("LOG_LEVEL")
	if level == "" {
		level = "info"
	}
	SetLevel(level)
}

func SetLevel(level string) {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		logger.Warn("Invalid log level, defaulting to info")
		lvl = logrus.InfoLevel
	}
	logger.SetLevel(lvl)
}

func getSourceInfo() logrus.Fields {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return logrus.Fields{"source": "unknown"}
	}

	// Get last 3 path segments for cleaner output
	parts := strings.Split(file, "/")
	if len(parts) > 3 {
		file = strings.Join(parts[len(parts)-3:], "/")
	}

	return logrus.Fields{"source": fmt.Sprintf("%s:%d", file, line)}
}

func Debug(args ...interface{}) {
	logger.WithFields(getSourceInfo()).Debug(args...)
}

func Info(args ...interface{}) {
	logger.WithFields(getSourceInfo()).Info(args...)
}

func Warn(args ...interface{}) {
	logger.WithFields(getSourceInfo()).Warn(args...)
}

func Error(args ...interface{}) {
	logger.WithFields(getSourceInfo()).Error(args...)
}

func Fatal(args ...interface{}) {
	logger.WithFields(getSourceInfo()).Fatal(args...)
}

func Debugf(format string, args ...interface{}) {
	logger.WithFields(getSourceInfo()).Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	logger.WithFields(getSourceInfo()).Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	logger.WithFields(getSourceInfo()).Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	logger.WithFields(getSourceInfo()).Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	logger.WithFields(getSourceInfo()).Fatalf(format, args...)
}

func WithFields(fields logrus.Fields) *logrus.Entry {
	return logger.WithFields(fields).WithFields(getSourceInfo())
}

func WithField(key string, value interface{}) *logrus.Entry {
	return logger.WithField(key, value).WithFields(getSourceInfo())
}
