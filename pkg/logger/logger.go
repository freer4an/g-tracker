package logger

import (
	"io"
	"log"
)

type Logger struct {
	info  *log.Logger
	error *log.Logger
	fatal *log.Logger
}

// InitLogger returns a new Logger instance.
// out is the io.Writer to which the log output will be written.
func InitLogger(out io.Writer) *Logger {
	logger := &Logger{}
	logger.info = log.New(out, "INFO: ", log.LstdFlags)
	logger.error = log.New(out, "ERROR: ", log.LstdFlags|log.Llongfile)
	logger.fatal = log.New(out, "FATAL: ", log.LstdFlags|log.Llongfile)
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
