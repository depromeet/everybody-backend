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

//func Test(t *testing.T) {
//	type A interface {
//	}
//	type B interface {
//	}
//
//	var (
//		a A
//		b B
//	)
//
//	a = 3
//	b = a
//	num, ok := a.(int)
//	fmt.Println(num, ok)
//	convert, ok := a.(B)
//	fmt.Println(convert, ok)
//	fmt.Println(b)
//
//}
