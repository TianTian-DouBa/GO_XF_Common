//package xf_utility
package main

import (
	"log"
	"os"
	"fmt"
)

var (
	LogThreadHold int = 30 //to add the log if 
	loggers map[int]*log.Logger
)

var LogCategory = map[int]string {
	10 : "[Error]    ",
	20 : "[Warning]  ",
	30 : "[Info]     ",
	40 : "[Trace]    ",
}

func Init() {
	handler := os.Stdout
	loggers = make(map[int]*log.Logger)
	for k, v := range LogCategory {
		loggers[k] = log.New(handler,v,log.Ldate|log.Ltime|log.Lshortfile)
	}
	loggers[99] = log.New(handler,"[undefine] ",log.Ldate|log.Ltime|log.Lshortfile) //undefined LogCategory 
}

func AddLog(logCategory int, logString string, args ...string) {
	if logCategory <= LogThreadHold {
		_, ok := loggers[logCategory]
		if ok {
			loggers[logCategory].Println(logString)
		} else {
			loggers[99].Println(logString)
		}	
	}
}

func main() {
	Init()
	/*
	AddLog(10, "test error")
	AddLog(20, "test warning")
	AddLog(30, "test info")
	AddLog(40, "test trace")
	AddLog(22, "test undefine")
	*/
	fmt.Println("----------------getMac---------------")
	nics, err := getMac()
	if err == nil {
		for _, nic := range nics {
			fmt.Println(nic.Index, nic.Name, nic.Mac)
		}
	}
	fmt.Println("----------------getMacOne---------------")
	nics1 := []Nic{
		{1,"aa","1c:1b:0d:e1:12:eb"},
		{2,"bb","00:50:56:c0:00:01"},
		{3,"cc","00:50:56:c0:00:08"},
		{4,"dd","1b:1b:0d:e1:12:eb"},
	}
	result, err := getMacOne(nics1)
	if err == nil {
		fmt.Println(result)
	}
	fmt.Println("----------------bios uuid---------------")
	biosUuid, err := getBiosUuid()
	if err == nil {
		fmt.Print(biosUuid)
	}
}

/*AddLog: generate log

func AddLog(logCategory int, logString string, logArgs ...string) {

}
*/