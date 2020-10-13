package qqbotapi

import (
	"fmt"
	"time"
)

func LogError(format string, a ...interface{}) {
	format = fmt.Sprintf("[%v] %s\n", time.Now().Format("2006-01-02 15:04"), format)
	fmt.Printf(format, a)
	return
}
