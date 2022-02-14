# Rollogger
![Build Status](https://github.com/roland-burke/rollogger/actions/workflows/simple-workflow.yml/badge.svg) ![Go Report](/badge/github.com/roland-burke/rollogger)

Simple Logger written in GO.

## Install
`go get -d -u github.com/roland-burke/rollogger`

## Documentation
In addition to the main log functions there are other Functions to interact with the logger:

| Syntax      | Description |
| ----------- | ----------- |
| Init(level int, colorLogs bool) *Log	| Initilailzes the logger with the level, and sets if the message should be colored.	|
| GetCurrentLogLevel() (int, string)   	| Return the current log level as a number and string.					|
| GetLastLog() string  			| Returns the latest log message, mainly used for testing.				|
| SetLogLevel(newLevel int)   		| Sets the current log level.								|

Typical Log functions:

| Syntax      | Description |
| ----------- | ----------- |
| Trace(msg string, args ...interface{})	| Logs messages for tracing	|
| Debug(msg string, args ...interface{})   	| Logs debug messages		|
| Info(msg string, args ...interface{})  	| Logs info messages		|
| Warn(msg string, args ...interface{}) 	| Logs warning messages		|
| Error(msg string, args ...interface{}) 	| Logs error messages		|


## Example

```
package main

import (
	"github.com/roland-burke/rollogger"
)

func main() {
	logger = rollogger.Init(rollogger.LEVEL_INFO, true)
		logger.Trace("Trace log message")
		logger.Debug("Debug log message")
		logger.Info("Info log message %s: %d", "with parameter", 42)
		logger.Warn("Warn log message")
		logger.Error("Error log message")
}
```
