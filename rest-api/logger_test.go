package main

import (
	log "github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestLoggerHook(t *testing.T) {
	initializeLogger()
	log.Error("1")
	log.Error("2")
	log.Error("3")
	log.Error("4")
	log.Error("5")
	time.Sleep(time.Second)
}
