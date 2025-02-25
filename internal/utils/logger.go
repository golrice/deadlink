package utils

import (
	"log"
	"os"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "[DeadLinkChecker] ", log.LstdFlags|log.Lshortfile)
}

func Debug(msg string) {
	logger.Println("[DEBUG] " + msg)
}

func Info(msg string) {
	logger.Println("[INFO] " + msg)
}

func Error(msg string, err error) {
	logger.Printf("[ERROR] %s: %v\n", msg, err)
}
