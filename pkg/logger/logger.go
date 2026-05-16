package logger

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"time"
)

type Level int

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
)

var levelNames = map[Level]string{
	DEBUG: "DEBUG",
	INFO:  "INFO",
	WARN:  "WARN",
	ERROR: "ERROR",
}

type Logger struct {
	level  Level
	output io.Writer
}

func New(level Level, output io.Writer) *Logger {
	return &Logger{
		level:  level,
		output: output,
	}
}

func NewDefault() *Logger {
	return New(INFO, os.Stdout)
}

func (l *Logger) log(level Level, format string, v ...interface{}) {
	if level < l.level {
		return
	}

	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "unknown"
		line = 0
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	prefix := formatLogPrefix(timestamp, levelNames[level], file, line)
	message := format

	if len(v) > 0 {
		message = formatLogMessage(format, v...)
	}

	logLine := prefix + " " + message + "\n"
	l.output.Write([]byte(logLine))
}

func (l *Logger) Debug(format string, v ...interface{}) {
	l.log(DEBUG, format, v...)
}

func (l *Logger) Info(format string, v ...interface{}) {
	l.log(INFO, format, v...)
}

func (l *Logger) Warn(format string, v ...interface{}) {
	l.log(WARN, format, v...)
}

func (l *Logger) Error(format string, v ...interface{}) {
	l.log(ERROR, format, v...)
}

func formatLogPrefix(timestamp, level, file string, line int) string {
	return timestamp + " [" + level + "] " + file + ":" + formatInt(line)
}

func formatInt(n int) string {
	return strconv.Itoa(n)
}

func formatLogMessage(format string, v ...interface{}) string {
	return fmt.Sprintf(format, v...)
}
