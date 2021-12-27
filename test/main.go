package main

import (
	"os"

	"github.com/universe-30/ULog"
)

func main() {
	l := ULog.New()

	l.SetLevel(ULog.TraceLevel)

	f, _ := os.Create("logFile")
	l.SetOutput(f)

	l.Trace("trace", 1)
	l.Debug("debug", 2)
	l.Info("info", 3)
	l.Warn("warn", 4)
	l.Error("error", 5)
	//l.Fatal("fatal", 6)
	//l.Panic("panic")
}
