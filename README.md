# Rollogger ![Build Status](https://github.com/roland-burke/rollogger/actions/workflows/build-and-test.yml/badge.svg) [![Coverage Status](https://coveralls.io/repos/github/roland-burke/rollogger/badge.svg?branch=master)](https://coveralls.io/github/roland-burke/rollogger?branch=master) ![Go Report Card](https://goreportcard.com/badge/github.com/roland-burke/rollogger)

This is my logger for go. It is desgined to be small, fast and minimalistic.

## Release Notes

### 1.1.0
- Added feature to override the log level with Environment variable
- Added a function to set the color setting

### 1.0.0
Initial release. Added basic features like logging with levels: Trace, Debug, Info, Warn and Error. Added also some utility functions and the ability to log with color and formatted.

## Install
`go get -d -u github.com/roland-burke/rollogger`

## Example

```
package main

import (
	"github.com/roland-burke/rollogger"
)

var logger *rollogger.Log

type Example struct {
	Value1 int
	Value2 bool
	Value3 string
}

func main() {

	var exampleObj = Example{
		Value1: 42,
		Value2: true,
		Value3: "Test",
	}

	logger = rollogger.Init(rollogger.INFO_LEVEL, true, true)
		logger.Trace("Trace log message")
		logger.Debug("Debug log message")
		logger.Info("Info log message %s: %d", "with parameter", 42)

		logger.InfoObj(exampleObj)
		logger.SetPrettyPrint(false)
		logger.InfoObj(exampleObj)

		logger.Warn("Warn log message")
		logger.Error("Error log message")
}
```
![screenshot](https://user-images.githubusercontent.com/56251366/153864038-f20aad06-ec05-49a6-a37e-49bc3d123d63.png)

## Documentation
Here is an overview of settings functions of the logger:

| Syntax      | Description |
| ----------- | ----------- |
| Init(level int, colorLogs bool, jsonPrettyPrint bool) *Log	| Initilailzes the logger with the level, the colorLogs and the json Pretty Print Flags	|
| GetCurrentLogLevel() (int, string)   							| Return the current log level as a number and string					|
| GetLastLog() string  											| Returns the latest log message, mainly used for testing				|
| SetLogLevel(newLevel int)   									| Sets the current log level								|
| SetPrettyPrint(newValue bool)					| Sets if Object logs should print pretty |
| SetColorLogs(newValue bool)                   | Sets if the logs should be colored |

Overview of how to set the LogLevel:

| Environment Variable      | Description |
| ----------- | ----------- |
| ROLLOGER_LOG_LEVEL						| Overrides the log level. Allowed values are: TRACE, DEBUG, INFO, WARN, ERROR. The values are not case-sensitive.		|


| Syntax      | Description |
| ----------- | ----------- |
| Trace(msg string, args ...interface{})						| Logs trace messages		|
| Debug(msg string, args ...interface{})   						| Logs debug messages		|
| Info(msg string, args ...interface{})  						| Logs info messages		|
| Warn(msg string, args ...interface{}) 						| Logs warning messages		|
| Error(msg string, args ...interface{}) 						| Logs error messages		|
| TraceObj(obj interface{})										| Logs an object on trace level		|
| DebugObj(obj interface{})   									| Logs an object on debug level		|
| InfoObj(obj interface{})  									| Logs an object on info level		|
| WarnObj(obj interface{}) 										| Logs an object on warn level		|
| ErrorObj(obj interface{}) 									| Logs an object on error level		|

