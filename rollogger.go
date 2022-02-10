package rollogger

import (
	"fmt"
	"time"
)

var LEVEL_NAMES = [5]string{"ERROR", "WARN", "INFO", "DEBUG", "TRACE"}
var LEVEL_COLORS = [5]string{"\033[31m", "\033[33m", "\033[36m", "", ""}

// The higher the number, the noisier the logger is
// A negative level mutes the logger
const (
	LEVEL_TRACE = 4
	LEVEL_DEBUG = 3
	LEVEL_INFO  = 2
	LEVEL_WARN  = 1
	LEVEL_ERROR = 0

	MAX_FILE_NAME_LENGTH = 25
	MAX_LOG_MSG_LENGTH   = 4000
)

type Log struct {
	rootLevel int
	colorLogs bool
	lastLog   string
}

func Init(level int, colorLogs bool) *Log {
	return &Log{
		rootLevel: level,
		colorLogs: colorLogs,
		lastLog:   "",
	}
}

func (l *Log) Trace(msg string, args ...interface{}) {
	if l.rootLevel >= LEVEL_TRACE {
		write(LEVEL_TRACE, fmt.Sprintf(msg, args...), l)
	}
}

func (l *Log) Debug(msg string, args ...interface{}) {
	if l.rootLevel >= LEVEL_DEBUG {
		write(LEVEL_DEBUG, fmt.Sprintf(msg, args...), l)
	}
}

func (l *Log) Info(msg string, args ...interface{}) {
	if l.rootLevel >= LEVEL_INFO {
		write(LEVEL_INFO, fmt.Sprintf(msg, args...), l)
	}
}

func (l *Log) Warn(msg string, args ...interface{}) {
	if l.rootLevel >= LEVEL_WARN {
		write(LEVEL_WARN, fmt.Sprintf(msg, args...), l)
	}
}

func (l *Log) Error(msg string, args ...interface{}) {
	if l.rootLevel >= LEVEL_ERROR {
		write(LEVEL_ERROR, fmt.Sprintf(msg, args...), l)
	}
}

func (l *Log) GetCurrentLogLevel() (int, string) {
	return l.rootLevel, LEVEL_NAMES[l.rootLevel]
}

func (l *Log) GetLastLog() string {
	return l.lastLog
}

func (l *Log) SetLogLevel(newLevel int) {
	l.rootLevel = newLevel
}

func truncateString(maxLength int, msg string) string {
	const offset = 3
	if maxLength > offset && len(msg) > maxLength+offset {
		return fmt.Sprintf("%s...", msg[:maxLength])
	}
	return msg
}

func write(msgLevel int, msg string, l *Log) {
	var fileName = ""

	if l.colorLogs {
		l.lastLog = fmt.Sprintf("%s %s[%-5s]\033[0m %s%s\n", time.Now().Format("02-01-2006 15:04:05.99 MST"), LEVEL_COLORS[msgLevel], LEVEL_NAMES[msgLevel], fileName, truncateString(MAX_LOG_MSG_LENGTH, msg))
	} else {
		l.lastLog = fmt.Sprintf("%s [%-5s] %s%s\n", time.Now().Format("02-01-2006 15:04:05.99 MST"), LEVEL_NAMES[msgLevel], fileName, truncateString(MAX_LOG_MSG_LENGTH, msg))
	}
	fmt.Print(l.lastLog)
}
