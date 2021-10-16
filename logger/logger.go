package logger

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
)

var (
	info        = color.New(color.FgCyan).PrintlnFunc()
	err         = color.New(color.FgRed).PrintlnFunc()
	fatal       = color.New(color.BgRed).PrintlnFunc()
	currentTime = time.Now()
)

func Info(contents ...interface{}) {
	content := formatMessage(contents...)
	info(currentTime.Format("[01-02 15:04:05.000000]"), "| ~INFO~ | ", content)
}

func Error(contents ...interface{}) {
	content := formatMessage(contents...)
	err(currentTime.Format("[01-02 15:04:05.000000]"), "| ~WARNING~ | ", content)
}

func Fatal(contents ...interface{}) {
	content := formatMessage(contents...)
	fatal(currentTime.Format("[01-02 15:04:05.000000]"), "| ~FATAL~ | ", content)
}

func formatMessage(args ...interface{}) string {
	msg := fmt.Sprintln(args...)
	msg = strings.TrimRight(msg, " \n\r")
	return msg
}
