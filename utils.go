package rebotapi

import (
	"fmt"
	"time"
)

func logError(format string, a ...interface{}) {
	format = fmt.Sprintf("[%v] %s", time.Now().Format("2006-01-02 15:04"), format)
	fmt.Printf(format, a)
	return
}
