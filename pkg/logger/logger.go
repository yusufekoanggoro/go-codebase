package logger

import (
	"fmt"
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
)

const (
	DevProfile  = "dev"
	ProdProfile = "prod"
	AppName     = "service"
)

type Logger interface {
	Info(message, event, key string)
	Warn(message, event, key string)
	Error(message, event, key string)
	Debug(message, event, key string)
	Fatal(message, event, key string)
}

// Logger struct wraps around the logrus logger
type LoggerImpl struct {
	*logrus.Logger
}

// NewLogger is the constructor for Logger
func NewLogger() *LoggerImpl {
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
	return &LoggerImpl{logger}
}

// GetCallerInfo gets the function name and line number where the log is called from
func (l *LoggerImpl) GetCallerInfo() string {
	// Get caller information
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		return "Unknown Caller"
	}

	funcName := runtime.FuncForPC(pc).Name()
	return fmt.Sprintf("%s:%d in %s", file, line, funcName)
}

func mergeFields(fields1, fields2 logrus.Fields) logrus.Fields {
	merged := make(logrus.Fields)
	for k, v := range fields1 {
		merged[k] = v
	}
	for k, v := range fields2 {
		merged[k] = v
	}
	return merged
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

func (l *LoggerImpl) Info(message, event, key string) {
	// Get caller information
	callerInfo := l.GetCallerInfo()

	// Create fields with caller info
	fields := logrus.Fields{
		"caller": callerInfo,
	}

	// Set additional fields based on event and key
	dynamicFields := setFields(event, key)

	// Combine the fields
	allFields := mergeFields(fields, dynamicFields)

	// Log the message with the combined fields
	entry := l.WithFields(allFields)
	entry.Info(message)
}

func (l *LoggerImpl) Warn(message, event, key string) {
	// Get caller information
	callerInfo := l.GetCallerInfo()

	// Create fields with caller info
	fields := logrus.Fields{
		"caller": callerInfo,
	}

	// Set additional fields based on event and key
	dynamicFields := setFields(event, key)

	// Combine the fields
	allFields := mergeFields(fields, dynamicFields)

	// Log the message with the combined fields
	entry := l.WithFields(allFields)
	entry.Warn(message)
}

func (l *LoggerImpl) Error(message, event, key string) {
	// Get caller information
	callerInfo := l.GetCallerInfo()

	// Create fields with caller info
	fields := logrus.Fields{
		"caller": callerInfo,
	}

	// Set additional fields based on event and key
	dynamicFields := setFields(event, key)

	// Combine the fields
	allFields := mergeFields(fields, dynamicFields)

	// Log the message with the combined fields
	entry := l.WithFields(allFields)
	entry.Error(message)
}

func (l *LoggerImpl) Debug(message, event, key string) {
	// Get caller information
	callerInfo := l.GetCallerInfo()

	// Create fields with caller info
	fields := logrus.Fields{
		"caller": callerInfo,
	}

	// Set additional fields based on event and key
	dynamicFields := setFields(event, key)

	// Combine the fields
	allFields := mergeFields(fields, dynamicFields)

	// Log the message with the combined fields
	entry := l.WithFields(allFields)
	entry.Debug(message)
}

func (l *LoggerImpl) Fatal(message, event, key string) {
	// Get caller information
	callerInfo := l.GetCallerInfo()

	// Create fields with caller info
	fields := logrus.Fields{
		"caller": callerInfo,
	}

	// Set additional fields based on event and key
	dynamicFields := setFields(event, key)

	// Combine the fields
	allFields := mergeFields(fields, dynamicFields)

	// Log the message with the combined fields
	entry := l.WithFields(allFields)
	entry.Fatal(message)
}
