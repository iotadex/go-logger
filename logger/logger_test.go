package logger

import (
	"log"
	"testing"
)

func TestLogger(t *testing.T) {
	OutLogger, err := New("out.log", 1, 3, 0, INFO)
	if err != nil {
		log.Panic("Create Outlogger file error. " + err.Error())
	}

	OutLogger.Info("Hello %d", 1)

	//second type log
	ErrLogger, err := New("err.log", 2, 100, 10, INFO)
	if err != nil {
		log.Panic("Create ErrLogger file error. " + err.Error())
	}

	ErrLogger.Error("Hello %d", 1)
}
