package main

import (
	"fmt"
	"os"
	"time"
)

const (
	logFilePath = "./service.log"
)

var (
	logTimeFormat = "02-01-2006 15:04:05"
)

func appendLog(message string) {

	f, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		appendLog("[Append log]: Error: " + err.Error())
	}
	defer f.Close()

	if config.Log.Format != "" {
		logTimeFormat = config.Log.Format
	}

	logTime := time.Now().Format(logTimeFormat)
	logMessage := fmt.Sprintf("[%s]%s\n", logTime, message)

	if _, err := f.WriteString(logMessage); err != nil {
		appendLog("[Append log]: Error: " + err.Error())
	}

	checkLogFileSize()
}

func checkLogFileSize() {
	fi, err := os.Stat(logFilePath)
	if err != nil {
		appendLog("[Log size check]: Error: " + err.Error())
	}

	size := fi.Size()
	if size/1000000 == config.Log.MaxFileSize {
		os.Remove(logFilePath)
	}
}
