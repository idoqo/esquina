package logger

import (
	"fmt"
	"log"
	"time"
)

func Info(message string) {
	instant := time.Now().Format("Mon Jan _2 15:04:05 2020")
	fmt.Printf(`%s ▶ INFO: %s`, instant, message)
	fmt.Println()
}

func Fatal(message string) {
	instant := time.Now().Format("Mon Jan _2 15:04:05 2020")
	log.Fatalf(`%s ▶ FATAL: %s`, instant, message)
}

func Panic(message string) {
	instant := time.Now().Format("Mon Jan _2 15:04:05 2020")
	log.Panicf(`%s ▶ PANIC: %s`, instant, message)
}