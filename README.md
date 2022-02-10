# rollogger
![Build Status](https://github.com/roland-burke/rollogger/actions/workflows/simple-workflow.yml/badge.svg)

Simple Logger written in GO

## Install

`go get -d -u github.com/roland-burke/rollogger`

## Usage

```
var log = Init(LEVEL_TRACE, true, true)
	log.Trace("Trace log message")
	log.Debug("Debug log message")
	log.Info("Info log message %s: %d", "with parameter", 42)
	log.Warn("Warn log message")
	log.Error("Error log message")
```
