package logger

import (
	"fmt"
	"os"
)

type consoleLog struct{

}

func NewConsoleLog() *consoleLog {
	return &consoleLog{}
}

func CreateConsoleLog (configurationName string) (interface{}, error) {
	return NewConsoleLog(), nil
}


func (this *consoleLog) Info(msg string) {
	fmt.Fprint(os.Stdout, "[INFO] " + msg + "\n")
}

func (this *consoleLog) Error(msg string) {
	fmt.Fprint(os.Stdout, "[ERROR] " + msg + "\n")
}
