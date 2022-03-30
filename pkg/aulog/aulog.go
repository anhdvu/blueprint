package aulog

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"runtime/debug"
	"sync"
	"time"
)

// Level represents security level for log entry
type Level uint8

// Security levels
const (
	LevelInfo Level = iota
	LevelError
	LevelFatal
	LevelOff
)

func (l Level) String() string {
	switch l {
	case LevelInfo:
		return "INFO"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	default:
		return ""
	}
}

// Logger represents a custom Aulab logger.
type Logger struct {
	service  string
	out      io.Writer
	minLevel Level
	mu       sync.Mutex
}

// New returns a newly initialized Logger with designated output destination and minimum severity.
func New(name string, out io.Writer, minlevel Level) *Logger {
	return &Logger{
		service:  name,
		out:      out,
		minLevel: minlevel,
	}
}

func (l *Logger) print(level Level, message string, properties map[string]string) (int, error) {
	if level < l.minLevel {
		return 0, nil
	}

	logEntry := struct {
		Service    string            `json:"service,omitempty"`
		Level      string            `json:"level,omitempty"`
		Time       string            `json:"time,omitempty"`
		Message    string            `json:"message,omitempty"`
		Properties map[string]string `json:"properties,omitempty"`
		Trace      string            `json:"trace,omitempty"`
	}{
		Service:    l.service,
		Level:      level.String(),
		Time:       time.Now().UTC().Format(time.RFC3339),
		Message:    message,
		Properties: properties,
	}

	if level >= LevelError {
		logEntry.Trace = string(debug.Stack())
	}

	logs, err := json.Marshal(logEntry)
	if err != nil {
		errLine := fmt.Sprintf("%s: unable to marshal log message %s\n", LevelError.String(), err.Error())
		logs = []byte(errLine)
	}
	logs = append(logs, '\n')

	l.mu.Lock()
	defer l.mu.Unlock()
	return l.out.Write(logs)
}

func (l *Logger) Write(message []byte) (n int, err error) {
	return l.print(LevelError, string(message), nil)
}

// Info prints log at INFO level.
func (l *Logger) Info(message string, properties map[string]string) {
	l.print(LevelInfo, message, properties)
}

// Errors prints log at ERROR level.
func (l *Logger) Errors(err error, properties map[string]string) {
	l.print(LevelError, err.Error(), properties)
}

// Fatal prints log at FATAL level and terminate the application.
func (l *Logger) Fatal(err error, properties map[string]string) {
	l.print(LevelFatal, err.Error(), properties)
	os.Exit(0)
}
