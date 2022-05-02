package main

import (
	"os"

	"github.com/coreservice-io/log"
)

func main() {

	l := log.NewBaseLogger()

	//default logLevel is Info
	l.SetLevel(log.TraceLevel)

	//default output is StdOut
	f, _ := os.Create("logFile")
	l.SetOutput(f)

	l.Traceln("trace", 1)
	l.Debugln("debug", 2)
	l.Infoln("info", 3)
	l.Warnln("warn", 4)
	l.Errorln("error", 5)
	l.PrintLastN(1000, []log.LogLevel{})
	//l.Fatalln("fatal", 6)
	//l.Panicln("panic")
}
