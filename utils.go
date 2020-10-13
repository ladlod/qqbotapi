package qqbotapi

import (
	"fmt"
	"time"
)

func LogError(format string, a ...interface{}) {
	format = fmt.Sprintf("[EROR][%v] %s\n", time.Now().Format("2006-01-02 15:04"), format)
	fmt.Printf(format, a)
	return
}

func LogInfo(format string, a ...interface{}) {
	format = fmt.Sprintf("[Info][%v] %s\n", time.Now().Format("2006-01-02 15:04"), format)
	fmt.Printf(format, a)
	return
}
