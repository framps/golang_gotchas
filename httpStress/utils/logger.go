package utils

import "fmt"

var debug bool

func Log(format string, args ...interface{}) (n int, err error) {
	if debug {
		n, err = fmt.Printf(format, args...)
	}
	return
}

func Logln(args ...interface{}) (n int, err error) {
	if debug {
		n, err = fmt.Println(args...)
	}
	return
}

func LogEnable() {
	debug = true
}
