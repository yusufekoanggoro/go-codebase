package logger

import (
	"fmt"
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
)

const (
	AppName = "service"
)

// Logger struct wraps around the logrus logger
type Logger struct {
	*logrus.Logger
}

// NewLogger is the constructor for Logger
func NewLogger() *Logger {
	logger := logrus.New()

	// Set log format to Text (default format for terminal)
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,  // Optional: add timestamps to each log entry
		DisableColors: false, // Enable colors in terminal output
		ForceQuote:    true,  // Quote the log message to make it clearer
	})

	// Set log level to Info (you can change it to Debug, Warn, etc.)
	logger.SetLevel(logrus.InfoLevel)

	// Set output to the terminal (standard output)
	logger.SetOutput(os.Stdout)

	// Return the Logger object
	return &Logger{logger}
}

// GetCallerInfo gets the function name and line number where the log is called from
func (l *Logger) GetCallerInfo() string {
	// Get caller information
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		return "Unknown Caller"
	}

	funcName := runtime.FuncForPC(pc).Name()
	return fmt.Sprintf("%s:%d in %s", file, line, funcName)
}

func setFields(event, key string) logrus.Fields {
	// Add or update fields based on event and key
	fields := logrus.Fields{
		"topic": AppName,
		"event": event,
		"key":   key,
	}
	return fields
}

func (l *Logger) Info(message, event, key string) {
	// Get caller information
	callerInfo := l.GetCallerInfo()

	// Create fields with caller info
	fields := logrus.Fields{
		"caller": callerInfo,
	}

	// Set additional fields based on event and key
	entry := l.WithFields(fields)
	entry.Info(message)

	// Dynamically set fields using the setFields function and log the message
	dynamicFields := setFields(event, key)
	entry.WithFields(dynamicFields).Info(message)
}

func (l *Logger) Warn(message, event, key string) {
	// Get caller information
	callerInfo := l.GetCallerInfo()

	// Create fields with caller info
	fields := logrus.Fields{
		"caller": callerInfo,
	}

	// Set additional fields based on event and key
	entry := l.WithFields(fields)
	entry.Warn(message)

	// Dynamically set fields using the setFields function and log the message
	dynamicFields := setFields(event, key)
	entry.WithFields(dynamicFields).Warn(message)
}

func (l *Logger) Error(message, event, key string) {
	// Get caller information
	callerInfo := l.GetCallerInfo()

	// Create fields with caller info
	fields := logrus.Fields{
		"caller": callerInfo,
	}

	// Set additional fields based on event and key
	entry := l.WithFields(fields)
	entry.Error(message)

	// Dynamically set fields using the setFields function and log the message
	dynamicFields := setFields(event, key)
	entry.WithFields(dynamicFields).Error(message)
}

func (l *Logger) Debug(message, event, key string) {
	// Get caller information
	callerInfo := l.GetCallerInfo()

	// Create fields with caller info
	fields := logrus.Fields{
		"caller": callerInfo,
	}

	// Set additional fields based on event and key
	entry := l.WithFields(fields)
	entry.Debug(message)

	// Dynamically set fields using the setFields function and log the message
	dynamicFields := setFields(event, key)
	entry.WithFields(dynamicFields).Debug(message)
}

func (l *Logger) Fatal(message, event, key string) {
	// Get caller information
	callerInfo := l.GetCallerInfo()

	// Create fields with caller info
	fields := logrus.Fields{
		"caller": callerInfo,
	}

	// Set additional fields based on event and key
	entry := l.WithFields(fields)
	entry.Fatal(message)

	// Dynamically set fields using the setFields function and log the message
	dynamicFields := setFields(event, key)
	entry.WithFields(dynamicFields).Fatal(message)
}
