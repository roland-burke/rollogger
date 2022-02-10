# rollogger
![Build Status](https://github.com/roland-burke/rollogger/actions/workflows/simple-workflow.yml/badge.svg)

Simple Logger written in GO

## Install

`go get -d -u github.com/roland-burke/rollogger`

## Usage

```
import (
	"github.com/roland-burke/rollogger"
)

logger = rollogger.Init(rollogger.LEVEL_INFO, true)
	logger.Trace("Trace log message")
	logger.Debug("Debug log message")
	logger.Info("Info log message %s: %d", "with parameter", 42)
	logger.Warn("Warn log message")
	logger.Error("Error log message")
```
