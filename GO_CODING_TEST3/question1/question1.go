package main

import (
	"fmt"
	"os"
)

type Logger interface {

	Info(message string)

	Warning(message string)

	Error(message string)
}

type Log struct {
	file *os.File
}

func NewLog(myfile string) (*Log,error) {
	file,err:=os.OpenFile(myfile,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0600)
	if err!=nil {
		return nil,err
	}
	return &Log{file:file},nil
}

func (l *Log) Info(message string) {
	l.log("[Info]:",message)
}

func (l *Log) Warning(message string) {
	l.log("[Warning]:",message)
}

func (l *Log) Error(message string) {
	l.log("[Error]:",message)
}

func (l *Log) log(prefix,message string) {
	logMessage := fmt.Sprintf("%s %s\n",prefix,message)
	l.file.WriteString(logMessage)
}


func main() {
	log, err :=NewLog("myfile.log")
	if err != nil {
		// fmt.Println("Error creating log file:",err)
		panic(err)
	}

	defer log.file.Close()

	log.Info("This is an information message.")
	log.Warning("This is a warning message.")
	log.Error("This is an error message.")
}