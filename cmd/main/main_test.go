package rollogger

import (
	"math"
	"os"
	"strings"
	"testing"
)

type Example struct {
	Value1 int
	Value2 bool
	Value3 string
}

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

func TestColoredLogs(t *testing.T) {
	const infoColor = "\033[36m"
	const warnColor = "\033[33m"
	const errorColor = "\033[31m"

	var log = Init(INFO_LEVEL, false, false)
	log.Info("test_log_without_color")
	var logMsgWithoutColor = log.GetLastLog()

	if strings.Contains(logMsgWithoutColor, infoColor) {
		t.Errorf("Got: %s, expected no color", logMsgWithoutColor)
	}

	log = Init(INFO_LEVEL, true, false)
	log.Info("test_log_with_color")
	var logMsgWithColor = log.GetLastLog()

	if !strings.Contains(logMsgWithColor, infoColor) {
		t.Errorf("Got: %s, expected color: %s", logMsgWithColor, infoColor)
	}

	log = Init(INFO_LEVEL, true, false)
	log.Warn("test_log_with_color_warn")
	var logMsgWithColorWarn = log.GetLastLog()

	if !strings.Contains(logMsgWithColorWarn, warnColor) {
		t.Errorf("Got: %s, expected color: %s", logMsgWithColorWarn, warnColor)
	}

	log = Init(INFO_LEVEL, true, false)
	log.Error("test_log_with_color_error")
	var logMsgWithColorError = log.GetLastLog()

	if !strings.Contains(logMsgWithColorError, errorColor) {
		t.Errorf("Got: %s, expected color: %s", logMsgWithColorError, errorColor)
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
	var log = Init(INFO_LEVEL, false, false)

	var levelNum, levelString = log.GetCurrentLogLevel()

	if levelNum != INFO_LEVEL || levelString != LEVEL_NAMES[2] {
		t.Errorf("Got: %d, %s expected: %d, %s", levelNum, levelString, INFO_LEVEL, LEVEL_NAMES[2])
	}
}

func TestSetCurrentLogLevel(t *testing.T) {
	var log = Init(INFO_LEVEL, false, false)

	log.SetLogLevel(WARN_LEVEL)
	var levelNum, _ = log.GetCurrentLogLevel()
	if levelNum != WARN_LEVEL {
		t.Errorf("Got: %d expected: %d", levelNum, WARN_LEVEL)
	}
}

func TestSetPrettyPrintLogLevel(t *testing.T) {
	var log = Init(INFO_LEVEL, false, false)

	log.SetPrettyPrint(true)

	if !log.jsonPrettyPrint {
		t.Errorf("Got: %t expected: %t", log.jsonPrettyPrint, true)
	}
}

func TestSetColorLogs(t *testing.T) {
	var log = Init(INFO_LEVEL, false, false)

	log.SetColorLogs(true)

	if !log.colorLogs {
		t.Errorf("Got: %t expected: %t", log.colorLogs, true)
	}
}

func TestStringParemeter(t *testing.T) {
	var log = Init(INFO_LEVEL, false, false)
	const expected = "with parameter: 42, second"

	log.Info("with parameter: %d, %s", 42, "second")
	var msg = log.GetLastLog()
	if !strings.Contains(msg, expected) {
		t.Errorf("Got: %s expected: %s", msg, expected)
	}
}

func TestConvertObjToStringPrettyPrint(t *testing.T) {
	var expected = ":\n{\n\t\"Value1\": 42,\n\t\"Value2\": true,\n\t\"Value3\": \"Moin\"\n}"
	var expected2 = "Error during Marshalling: json: unsupported value: +Inf"

	var exampleObj = Example{
		Value1: 42,
		Value2: true,
		Value3: "Moin",
	}

	var objString = convertJsonObjectToString(exampleObj, true)

	if !strings.Contains(objString, expected) {
		t.Errorf("Got: %s expected: %s", objString, expected)
	}

	var objString2 = convertJsonObjectToString(math.Inf(1), true)

	if !strings.Contains(objString2, expected2) {
		t.Errorf("Got: %s expected: %s", objString2, expected)
	}

}

func TestConvertObjToStringNormalPrint(t *testing.T) {
	var expected = "{Value1:42 Value2:true Value3:Moin}"
	var expected2 = "+Inf"

	var exampleObj = Example{
		Value1: 42,
		Value2: true,
		Value3: "Moin",
	}

	var objString = convertJsonObjectToString(exampleObj, false)

	if !strings.Contains(objString, expected) {
		t.Errorf("Got: %s expected: %s", objString, expected)
	}

	var objString2 = convertJsonObjectToString(math.Inf(1), false)

	if !strings.Contains(objString2, expected2) {
		t.Errorf("Got: %s expected: %s", objString2, expected2)
	}

}

func TestLogObjMethods(t *testing.T) {
	var expected = ":\n{\n\t\"Value1\": 42,\n\t\"Value2\": true,\n\t\"Value3\": \"Moin\"\n}"

	var log = Init(TRACE_LEVEL, false, true)

	var exampleObj = Example{
		Value1: 42,
		Value2: true,
		Value3: "Moin",
	}

	log.TraceObj(exampleObj)
	var lastTrace = log.GetLastLog()
	log.DebugObj(exampleObj)
	var lastDebug = log.GetLastLog()
	log.InfoObj(exampleObj)
	var lastInfo = log.GetLastLog()
	log.WarnObj(exampleObj)
	var lastWarn = log.GetLastLog()
	log.ErrorObj(exampleObj)
	var lastError = log.GetLastLog()

	if !strings.Contains(lastTrace, expected) {
		t.Errorf("Got: %s expected: %s", lastTrace, expected)
	}

	if !strings.Contains(lastDebug, expected) {
		t.Errorf("Got: %s expected: %s", lastTrace, expected)
	}

	if !strings.Contains(lastInfo, expected) {
		t.Errorf("Got: %s expected: %s", lastInfo, expected)
	}

	if !strings.Contains(lastWarn, expected) {
		t.Errorf("Got: %s expected: %s", lastWarn, expected)
	}

	if !strings.Contains(lastError, expected) {
		t.Errorf("Got: %s expected: %s", lastError, expected)
	}
}

func TestSetEnvironmentVariableShouldOverrideLogLevel(t *testing.T) {
	// Test uppercase
	os.Setenv("ROLLOGER_LOG_LEVEL", "DEBUG")
	var log = Init(INFO_LEVEL, false, false)

	level, levelname := log.GetCurrentLogLevel()

	var expectedLevel = 3
	var expectedLevelName = "DEBUG"

	if strings.Compare(levelname, expectedLevelName) != 0 {
		t.Errorf("Got: %s expected: %s", levelname, expectedLevelName)
	}

	if level != expectedLevel {
		t.Errorf("Got: %d expected: %d", level, expectedLevel)
	}

	// Test lowercase
	os.Setenv("ROLLOGER_LOG_LEVEL", "warn")
	log = Init(INFO_LEVEL, false, false)

	level, levelname = log.GetCurrentLogLevel()

	expectedLevel = 1
	expectedLevelName = "WARN"

	if strings.Compare(levelname, expectedLevelName) != 0 {
		t.Errorf("Got: %s expected: %s", levelname, expectedLevelName)
	}

	if level != expectedLevel {
		t.Errorf("Got: %d expected: %d", level, expectedLevel)
	}
}

func TestSetEnvironmentVariableWrongShouldNotOverrideLogLevel(t *testing.T) {
	os.Setenv("ROLLOGER_LOG_LEVEL", "someInvalidValue")
	var log = Init(INFO_LEVEL, false, false)

	level, levelname := log.GetCurrentLogLevel()

	var expectedLevel = 2
	var expectedLevelName = "INFO"

	if strings.Compare(levelname, expectedLevelName) != 0 {
		t.Errorf("Got: %s expected: %s", levelname, expectedLevelName)
	}

	if level != expectedLevel {
		t.Errorf("Got: %d expected: %d", level, expectedLevel)
	}
}
