package main

import (
	"strings"
	"testing"
)

func TestLogLevels(t *testing.T) {
	for i := -1; i < 6; i++ {
		var log = Init(i, false, false)

		var logMessages = [5]string{"", "", "", "", ""}

		log.Trace("4_log")
		logMessages[0] = log.GetLastLog()

		log.Debug("3_log")
		logMessages[1] = log.GetLastLog()

		log.Info("2_log")
		logMessages[2] = log.GetLastLog()

		log.Warn("1_log")
		logMessages[3] = log.GetLastLog()

		log.Error("0_log")
		logMessages[4] = log.GetLastLog()

		if i < 0 {
			if !(strings.Contains(logMessages[0], "") && strings.Contains(logMessages[1], "") && strings.Contains(logMessages[2], "") && strings.Contains(logMessages[3], "") && strings.Contains(logMessages[4], "")) {
				t.Errorf("Got:\n%s%s%s%s%sexpected contents: %s", logMessages[0], logMessages[1], logMessages[2], logMessages[3], logMessages[4], ", , , , ")
			}
		} else if i == 0 {
			if !(strings.Contains(logMessages[0], "") && strings.Contains(logMessages[1], "") && strings.Contains(logMessages[2], "") && strings.Contains(logMessages[3], "") && strings.Contains(logMessages[4], "0_log")) {
				t.Errorf("Got:\n%s%s%s%s%sexpected contents: %s", logMessages[0], logMessages[1], logMessages[2], logMessages[3], logMessages[4], ", , , , 0_log")
			}
		} else if i == 1 {
			if !(strings.Contains(logMessages[0], "") && strings.Contains(logMessages[1], "") && strings.Contains(logMessages[2], "") && strings.Contains(logMessages[3], "1_log") && strings.Contains(logMessages[4], "0_log")) {
				t.Errorf("Got:\n%s%s%s%s%sexpected contents: %s", logMessages[0], logMessages[1], logMessages[2], logMessages[3], logMessages[4], ", , , 1_log, 0_log")
			}
		} else if i == 2 {
			if !(strings.Contains(logMessages[0], "") && strings.Contains(logMessages[1], "") && strings.Contains(logMessages[2], "2_log") && strings.Contains(logMessages[3], "1_log") && strings.Contains(logMessages[4], "0_log")) {
				t.Errorf("Got:\n%s%s%s%s%sexpected contents: %s", logMessages[0], logMessages[1], logMessages[2], logMessages[3], logMessages[4], ", , 2_log, 1_log, 0_log")
			}
		} else if i == 3 {
			if !(strings.Contains(logMessages[0], "") && strings.Contains(logMessages[1], "3_log") && strings.Contains(logMessages[2], "2_log") && strings.Contains(logMessages[3], "1_log") && strings.Contains(logMessages[4], "0_log")) {
				t.Errorf("Got:\n%s%s%s%s%sexpected contents: %s", logMessages[0], logMessages[1], logMessages[2], logMessages[3], logMessages[4], ", 3_log, 2_log, 1_log, 0_log")
			}
		} else if i >= 4 {
			if !(strings.Contains(logMessages[0], "4_log") && strings.Contains(logMessages[1], "3_log") && strings.Contains(logMessages[2], "2_log") && strings.Contains(logMessages[3], "1_log") && strings.Contains(logMessages[4], "0_log")) {
				t.Errorf("Got:\n%s%s%s%s%sexpected contents: %s", logMessages[0], logMessages[1], logMessages[2], logMessages[3], logMessages[4], "4_log, 3_log, 2_log, 1_log, 0_log")
			}
		}
	}
}

func TestPrintFileName(t *testing.T) {
	const fileName = "rolloger.go"

	var log = Init(LEVEL_INFO, false, false)
	log.Info("test_log_1")
	var logMsgWithoutFile = log.GetLastLog()

	if strings.Contains(logMsgWithoutFile, fileName) {
		t.Errorf("Got: %s, expected no fileName", logMsgWithoutFile)
	}

	log = Init(LEVEL_INFO, true, false)
	log.Info("test_log_2")
	var logMsgWithFile = log.GetLastLog()

	if !strings.Contains(logMsgWithFile, fileName) {
		t.Errorf("Got: %s, expected fileName: %s", logMsgWithFile, fileName)
	}
}

func TestColoredLogs(t *testing.T) {
	const infoColor = "\033[36m"

	var log = Init(LEVEL_INFO, false, false)
	log.Info("test_log_without_color")
	var logMsgWithoutColor = log.GetLastLog()

	if strings.Contains(logMsgWithoutColor, infoColor) {
		t.Errorf("Got: %s, expected no color", logMsgWithoutColor)
	}

	log = Init(LEVEL_INFO, false, true)
	log.Info("test_log_with_color")
	var logMsgWithColor = log.GetLastLog()

	if !strings.Contains(logMsgWithColor, infoColor) {
		t.Errorf("Got: %s, expected color: %s", logMsgWithColor, infoColor)
	}
}

func TestTruncateString(t *testing.T) {
	const (
		maxLength = 5
		expected1 = "testa"
		expected2 = "testab"
		expected3 = "testabc"
		expected4 = "testabcd"
		expected5 = "testa..."
		expected6 = ""
		expected7 = "t"
	)

	var result1 = truncateString(maxLength, "testa")

	if result1 != expected1 {
		t.Errorf("Got: %s, expected: %s", result1, expected1)
	}

	var result2 = truncateString(maxLength, "testab")
	if result2 != expected2 {
		t.Errorf("Got: %s, expected: %s", result2, expected2)
	}

	var result3 = truncateString(maxLength, "testabc")
	if result3 != expected3 {
		t.Errorf("Got: %s, expected: %s", result3, expected3)
	}

	var result4 = truncateString(maxLength, "testabcd")
	if result4 != expected4 {
		t.Errorf("Got: %s, expected: %s", result4, expected4)
	}

	var result5 = truncateString(maxLength, "testabcde")
	if result5 != expected5 {
		t.Errorf("Got: %s, expected: %s", result5, expected5)
	}

	var result6 = truncateString(maxLength, "")
	if result6 != expected6 {
		t.Errorf("Got: %s, expected: %s", result6, expected6)
	}

	var result7 = truncateString(maxLength, "t")
	if result7 != expected7 {
		t.Errorf("Got: %s, expected: %s", result7, expected7)
	}
}

func TestGetCurrentLogLevel(t *testing.T) {
	var log = Init(LEVEL_INFO, false, false)

	var levelNum, levelString = log.GetCurrentLogLevel()

	if levelNum != LEVEL_INFO || levelString != LEVEL_NAMES[2] {
		t.Errorf("Got: %d, %s expected: %d, %s", levelNum, levelString, LEVEL_INFO, LEVEL_NAMES[2])
	}
}

func TestSetCurrentLogLevel(t *testing.T) {
	var log = Init(LEVEL_INFO, false, false)

	log.SetLogLevel(LEVEL_WARN)
	var levelNum, _ = log.GetCurrentLogLevel()
	if levelNum != LEVEL_WARN {
		t.Errorf("Got: %d expected: %d", levelNum, LEVEL_WARN)
	}
}
