package logger

import (
	"io"
	"log"
	"os"
)

const (
	Red    = "\033[0;31mFATAL:\033[0m\t"
	Green  = "\033[0;32mINFO:\033[0m\t"
	Yellow = "\033[0;33mERROR:\033[0m\t"
)

type Logger struct {
	log *log.Logger
}

// InitLogger returns a new Logger instance.
// out is the io.Writer to which the log output will be written.
func InitLogger(out io.Writer) *Logger {
	if out == nil {
		out = os.Stdout
	}

	logger := &Logger{}
	logger.log = log.New(out, "", log.LstdFlags)
	return logger
}

// Info writes a log message with the info level.
func (l *Logger) Info(msg string) {
	l.log.Println(Green, msg)
}

// Infof writes a formatted log message with the info level.
func (l *Logger) Infof(format string, v ...interface{}) {
	l.log.Printf(Green+format, v...)
}

// Errorf prints a formatted error message to the logger's error output.
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.log.Printf(Yellow+format, v...)
}

// Fatalf prints a formatted error message to the logger's error output.
// NOTE: it then calls os.Exit(1).
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.log.Fatalf(Red+format, v...)
}
