package main

// Insert log statements in code using go.uber.org/zap
//
// See github.com/framps/golang_gotchas for latest code
//
// Copyright (C) 2017 framp at linux-tips-and-tricks dot de

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func hookD(e zapcore.Entry) error {
	fmt.Print("*** DDD Hook called ***\n")
	return nil
}

func hookP(e zapcore.Entry) error {
	fmt.Print("*** PPP Hook called ***\n")
	return nil
}

func main() {

	zap.L().Info("ZAP")

	loggerd, err := zap.NewDevelopment(
		zap.Fields(zap.String("name", "Dev")),
		zap.Hooks(hookD))
	if err != nil {
		panic(err)
	}

	loggerd.Info("Starting DDD")

	loggerp, err := zap.NewProduction(
		zap.Fields(zap.String("name", "Prod")),
		zap.Hooks(hookP))
	if err != nil {
		panic(err)
	}

	loggerp.Info("Starting PPP")

}
