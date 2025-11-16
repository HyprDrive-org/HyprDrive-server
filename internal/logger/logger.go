package logger

import (
	"io"
	"log"
	"os"
)

// Logger wraps the standard logger so it can be swapped out later.
type Logger struct {
	*log.Logger
}

// New creates a new logger that writes to stdout.
func New() *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, "[hyprdrive] ", log.LstdFlags|log.Lshortfile),
	}
}

// WithOutput allows redirecting the logger output, useful for tests.
func (l *Logger) WithOutput(w io.Writer) {
	l.SetOutput(w)
}
