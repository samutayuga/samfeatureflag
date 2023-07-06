package main

import (
	"log"
	"samfeatureflag/cmd"
	"samfeatureflag/tracer"

	"go.uber.org/zap"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if err := cmd.Execute(); err != nil {
		tracer.Logger.DPanic("error while while executing the command", zap.String("error", err.Error()))
	}
}
