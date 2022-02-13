package rollogger

import (
	"encoding/json"
	"fmt"
	"time"
)

var LEVEL_NAMES = [5]string{"[ERROR]", "[WARN]", "[INFO]", "[DEBUG]", "[TRACE]"}
var LEVEL_COLORS = [5]string{"\033[31m", "\033[33m", "\033[36m", "", ""}

// The higher the number, the noisier the logger is
// A negative level mutes the logger
const (
	TRACE_LEVEL = 4
	DEBUG_LEVEL = 3
	INFO_LEVEL  = 2
	WARN_LEVEL  = 1
	ERROR_LEVEL = 0

	MAX_LOG_MSG_LENGTH = 5000
)

type Log struct {
	rootLevel int
	colorLogs bool
	lastLog   string
}

func convertJsonObjectToString(object interface{}) string {
	var jsonObj, err = json.MarshalIndent(object, "", "\t")
	if err != nil {
		return fmt.Sprintf("Error during Marshalling: %s", err.Error())
	}
	return fmt.Sprintf(":\n%s", string(jsonObj))
}

func Init(level int, colorLogs bool) *Log {
	return &Log{
		rootLevel: level,
		colorLogs: colorLogs,
		lastLog:   "",
	}
}

func (l *Log) Trace(msg string, args ...interface{}) {
	if l.rootLevel >= TRACE_LEVEL {
		write(TRACE_LEVEL, fmt.Sprintf(msg, args...), l)
	}
}

func (l *Log) TraceObj(obj interface{}) {
	if l.rootLevel >= TRACE_LEVEL {
		write(TRACE_LEVEL, convertJsonObjectToString(obj), l)
	}
}

func (l *Log) Debug(msg string, args ...interface{}) {
	if l.rootLevel >= DEBUG_LEVEL {
		write(DEBUG_LEVEL, fmt.Sprintf(msg, args...), l)
	}
}

func (l *Log) DebugObj(obj interface{}) {
	if l.rootLevel >= DEBUG_LEVEL {
		write(DEBUG_LEVEL, convertJsonObjectToString(obj), l)
	}
}

func (l *Log) Info(msg string, args ...interface{}) {
	if l.rootLevel >= INFO_LEVEL {
		write(INFO_LEVEL, fmt.Sprintf(msg, args...), l)
	}
}

func (l *Log) InfoObj(obj interface{}) {
	if l.rootLevel >= INFO_LEVEL {
		write(INFO_LEVEL, convertJsonObjectToString(obj), l)
	}
}

func (l *Log) Warn(msg string, args ...interface{}) {
	if l.rootLevel >= WARN_LEVEL {
		write(WARN_LEVEL, fmt.Sprintf(msg, args...), l)
	}
}

func (l *Log) WarnObj(obj interface{}) {
	if l.rootLevel >= WARN_LEVEL {
		write(WARN_LEVEL, convertJsonObjectToString(obj), l)
	}
}

func (l *Log) Error(msg string, args ...interface{}) {
	if l.rootLevel >= ERROR_LEVEL {
		write(ERROR_LEVEL, fmt.Sprintf(msg, args...), l)
	}
}

func (l *Log) ErrorObj(obj interface{}) {
	if l.rootLevel >= ERROR_LEVEL {
		write(ERROR_LEVEL, convertJsonObjectToString(obj), l)
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

	if l.colorLogs {
		l.lastLog = fmt.Sprintf("%s %s%-7s\033[0m %s\n", time.Now().Format("02-01-2006 15:04:05.99 MST"), LEVEL_COLORS[msgLevel], LEVEL_NAMES[msgLevel], truncateString(MAX_LOG_MSG_LENGTH, msg))
	} else {
		l.lastLog = fmt.Sprintf("%s %-7s %s\n", time.Now().Format("02-01-2006 15:04:05.99 MST"), LEVEL_NAMES[msgLevel], truncateString(MAX_LOG_MSG_LENGTH, msg))
	}
	fmt.Print(l.lastLog)
}
