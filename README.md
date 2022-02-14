# Rollogger ![Build Status](https://github.com/roland-burke/rollogger/actions/workflows/simple-workflow.yml/badge.svg) ![Go Report Card](https://goreportcard.com/badge/github.com/roland-burke/rollogger)

This is my logger for go. It is desgined to be small and minimalistic.

## Install
`go get -d -u github.com/roland-burke/rollogger`

## Example

```
package main

import (
	"github.com/roland-burke/rollogger"
)

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

Overview of all logging functions:

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

