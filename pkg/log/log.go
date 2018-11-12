// Package log provide common logger interface for logging messages and errors.
package log

import "os"

// Logger common loger interface.
type Logger interface {
	// Info writes a information message.
	Info(...interface{})
	// Infof writes a formated information message.
	Infof(string, ...interface{})
	// Warn writes a warning message.
	Warn(...interface{})
	// Warnf writes a formated warning message.
	Warnf(string, ...interface{})
	// Error writes an error message.
	Error(...interface{})
	// Errorf writes a formated error message.
	Errorf(string, ...interface{})
	// Debug writes a debug message.
	Debug(...interface{})
	// Debugf writes a formated debug message.
	Debugf(string, ...interface{})
}

// Default main logger of the package.
// Is initialized with the function Set.
var logger Logger

// Set sets the default logger.
func Set(l Logger) {
	logger = l
}

// Info writes the information message using the default logger.
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Infof writes the formated information message using the default logger.
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

// Warn writes the warning message using the default logger.
func Warn(args ...interface{}) {
	logger.Warn(args...)
}

// Warnf writes the formated warning message using the default logger.
func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

// Error writes the error message using the default logger.
func Error(args ...interface{}) {
	logger.Error(args...)
}

// Errorf writes the formated error message using the default logger.
func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

// Debug writes the debug message using the default logger.
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

// Debugf writes the formated debug message using the default logger.
func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

// Fatal is equivalent to Error() followed by a call to os.Exit(1).
func Fatal(args ...interface{}) {
	logger.Error(args...)
	os.Exit(1)
}

// Fatalf is equivalent to Errorf() followed by a call to os.Exit(1).
func Fatalf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
	os.Exit(1)
}
