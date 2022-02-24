package main

import (
	"go-logger/logger"
	"log"
)

func main() {
	OutLogger, err := logger.New("out.log", 1, 3, 0)
	if err != nil {
		log.Panic("Create Outlogger file error. " + err.Error())
	}

	OutLogger.Info("Hello %d", 1)

	//second type log
	ErrLogger, err := logger.New("err.log", 2, 100, 10)
	if err != nil {
		log.Panic("Create ErrLogger file error. " + err.Error())
	}

	ErrLogger.Error("Hello %d", 1)
}
