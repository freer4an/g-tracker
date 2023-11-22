package logger

import (
	"io"
	"log"
	"os"
)

const (
	Red    = "\033[0;31mERROR:\033[0m\t"
	Green  = "\033[0;32mINFO:\033[0m\t"
	Yellow = "\033[0;33mFATAL:\033[0m\t"
)

type Logger struct {
	info  *log.Logger
	error *log.Logger
	fatal *log.Logger
}

// InitLogger returns a new Logger instance.
// out is the io.Writer to which the log output will be written.
func InitLogger(out io.Writer) *Logger {
	if out == nil {
		out = os.Stdout
	}
	logger := &Logger{}
	logger.info = log.New(out, Green, log.LstdFlags)
	logger.error = log.New(out, Red, log.LstdFlags)
	logger.fatal = log.New(out, Yellow, log.LstdFlags)
	return logger
}

// Infof writes a formatted log message with the info level.
func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}

// Errorf prints a formatted error message to the logger's error output.
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.error.Printf(format, v...)
}

// Fatalf prints a formatted error message to the logger's error output.
// NOTE: it then calls os.Exit(1).
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.fatal.Fatalf(format, v...)
}
