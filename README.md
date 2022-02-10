# rollogger
Simple Logger written in GO

## Install

`go get -u github.com/roland-burke/rolloger`

## Usage

```
var log = Init(LEVEL_TRACE, true, true)
	log.Trace("Trace log message")
	log.Debug("Debug log message")
	log.Info("Info log message %s: %d", "with parameter", 42)
	log.Warn("Warn log message")
	log.Error("Error log message")
```