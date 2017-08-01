package utils

import (
	"fmt"
	"runtime"
	"strings"
)

func FuncName(offset int) string {
	pc, _, _, _ := runtime.Caller(offset)
	fullName := runtime.FuncForPC(pc).Name()
	parts := strings.Split(fullName, "/")
	return parts[len(parts)-1:][0]
}

func Debugf(format string, a ...interface{}) {
	info := fmt.Sprintf(format, a...)
	fmt.Printf("%s: %s", FuncName(2), info)
}

func Debugln(a ...interface{}) {
	info := fmt.Sprintln(a...)
	fmt.Printf("%s: %s", FuncName(2), info)
}
