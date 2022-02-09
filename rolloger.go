package main

import (
	"fmt"
	"runtime"
	"strings"
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
	rootLevel     int
	printFileName bool
	colorLogs     bool
	lastLog       string
}

func Init(level int, printFileName bool, colorLogs bool) *Log {
	return &Log{
		rootLevel:     level,
		printFileName: printFileName,
		colorLogs:     colorLogs,
		lastLog:       "",
	}
}

func (l *Log) Trace(msg string) {
	if l.rootLevel >= LEVEL_TRACE {
		write(LEVEL_TRACE, msg, l)
	}
}

func (l *Log) Debug(msg string) {
	if l.rootLevel >= LEVEL_DEBUG {
		write(LEVEL_DEBUG, msg, l)
	}
}

func (l *Log) Info(msg string) {
	if l.rootLevel >= LEVEL_INFO {
		write(LEVEL_INFO, msg, l)
	}
}

func (l *Log) Warn(msg string) {
	if l.rootLevel >= LEVEL_WARN {
		write(LEVEL_WARN, msg, l)
	}
}

func (l *Log) Error(msg string) {
	if l.rootLevel >= LEVEL_ERROR {
		write(LEVEL_ERROR, msg, l)
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

func getFileNameString(colorLogs bool) string {
	_, file, _, ok := runtime.Caller(1)
	if ok {
		var filePath = strings.Split(file, "/")
		var fileName = filePath[len(filePath)-1]

		if colorLogs {
			return fmt.Sprintf("\033[32m%s\033[0m: ", truncateString(MAX_FILE_NAME_LENGTH, fileName))
		} else {
			return fmt.Sprintf("%s: ", truncateString(MAX_FILE_NAME_LENGTH, fileName))
		}
	}
	return ""
}

func write(msgLevel int, msg string, l *Log) {
	var fileName = ""

	if l.printFileName {
		fileName = getFileNameString(l.colorLogs)
	}

	if l.colorLogs {
		l.lastLog = fmt.Sprintf("%s %s[%-5s]\033[0m %s%s\n", time.Now().Format("02-01-2006 15:04:05.99 MST"), LEVEL_COLORS[msgLevel], LEVEL_NAMES[msgLevel], fileName, truncateString(MAX_LOG_MSG_LENGTH, msg))
	} else {
		l.lastLog = fmt.Sprintf("%s [%-5s] %s%s\n", time.Now().Format("02-01-2006 15:04:05.99 MST"), LEVEL_NAMES[msgLevel], fileName, truncateString(MAX_LOG_MSG_LENGTH, msg))
	}
	fmt.Print(l.lastLog)
}

func main() {
	var log = Init(LEVEL_TRACE, true, true)
	log.Trace("Trace log message")
	log.Debug("Debug log message")
	log.Info("Info log message")
	log.Warn("Warn log message")
	log.Error("Error log message")
}
