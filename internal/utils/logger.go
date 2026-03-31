package utils

import (
	"fmt"
	"log"
)

var Logger = &logger{}

var colors = map[string]string{
	"INFO":    "\033[34m", // Blue
	"ERROR":   "\033[31m", // Red
	"DEBUG":   "\033[36m", // Cyan
	"WARN":    "\033[33m", // Yellow
	"LOG":     "\033[0m",  // White (default)
	"SUCCESS": "\033[32m", // Green
	"FATAL":   "\033[31m", // Red
	"PANIC":   "\033[31m", // Red
	"TRACE":   "\033[35m", // Magenta
	"RESET":   "\033[0m",  // Reset
}

type logger struct {
	Level   string `json:"level"`
	Message string `json:"message"`
}

func (l *logger) print() {
	var output string
	color, exists := colors[l.Level]
	if !exists {
		color = colors["RESET"]
	}
	output = fmt.Sprintf("%s[%s] %s%s", color, l.Level, l.Message, colors["RESET"])
	log.Println(output)
}

func (l *logger) Info(msg string, args ...interface{}) error {
	l.Level = "INFO"
	l.Message = fmt.Sprintf(msg, args...)
	l.print()
	return nil
}

func (l *logger) Error(msg string, args ...interface{}) error {
	l.Level = "ERROR"
	l.Message = fmt.Sprintf(msg, args...)
	l.print()
	return fmt.Errorf(l.Message)
}

func (l *logger) Debug(msg string, args ...interface{}) error {
	l.Level = "DEBUG"
	l.Message = fmt.Sprintf(msg, args...)
	l.print()
	return nil
}

func (l *logger) Warn(msg string, args ...interface{}) error {
	l.Level = "WARN"
	l.Message = fmt.Sprintf(msg, args...)
	l.print()
	return nil
}

func (l *logger) Log(msg string, args ...interface{}) error {
	l.Level = "LOG"
	l.Message = fmt.Sprintf(msg, args...)
	l.print()
	return nil
}

func (l *logger) Success(msg string, args ...interface{}) error {
	l.Level = "SUCCESS"
	l.Message = fmt.Sprintf(msg, args...)
	l.print()
	return nil
}

func (l *logger) Fatal(msg string, args ...interface{}) error {
	l.Level = "FATAL"
	l.Message = fmt.Sprintf(msg, args...)
	l.print()
	return fmt.Errorf(l.Message)
}

func (l *logger) Panic(msg string, args ...interface{}) {
	l.Level = "PANIC"
	l.Message = fmt.Sprintf(msg, args...)
	l.print()
	panic(l.Message)
}

func (l *logger) Trace(msg string, args ...interface{}) error {
	l.Level = "TRACE"
	l.Message = fmt.Sprintf(msg, args...)
	l.print()
	return nil
}
