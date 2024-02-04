package util

import "fmt"

const (
	ERR = Red + "Err" + ResetColor
	WARN = Yellow + "Warn" + ResetColor
	OK = Green + "Ok" + ResetColor
)

func LogOk(msg string) {
	log(OK, msg)
}

func LogWarn(msg string) {
	log(WARN, msg)
}

func LogErr(msg string) {
	log(ERR, msg)
}

func log(level string, msg string) {
	fmt.Printf("[%s] > %s", level, msg)
}
