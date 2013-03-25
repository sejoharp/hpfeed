package helper

import (
	"log"
)

func LogInfo(msg string)  { log.Println(msg) }
func LogError(msg string) { log.Println(msg) }
func LogFatal(msg string) { log.Fatal(msg) }
func HandleFatalError(msg string, err error) {
	if err != nil {
		LogFatal(msg + err.Error())
	}
}
func HandleError(msg string, err error) {
	if err != nil {
		LogError(msg + err.Error())
	}
}
