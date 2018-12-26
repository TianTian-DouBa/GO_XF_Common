//package xf_utility
package main

import (
	"log"
	"os"
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

	AddLog(10, "test error")
	AddLog(20, "test warning")
	AddLog(30, "test info")
	AddLog(40, "test trace")
	AddLog(22, "test undefine")
}

/*AddLog: generate log

func AddLog(logCategory int, logString string, logArgs ...string) {

}
*/