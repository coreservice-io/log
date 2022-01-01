# ULog

Base logger of universe-30 project. Define different log level and logger interface

### How to use
```
go get github.com/universe-30/ULog
```

#### example
```go
package main

import (
	"os"

	"github.com/universe-30/ULog"
)

func main() {
	l := ULog.New()

	l.SetLevel(ULog.TraceLevel)

	//default output is StdOut
	f, _ := os.Create("logFile")
	l.SetOutput(f)

	l.Traceln("trace", 1)
	l.Debugln("debug", 2)
	l.Infoln("info", 3)
	l.Warnln("warn", 4)
	l.Errorln("error", 5)
	//l.Fatalln("fatal", 6)
	//l.Panicln("panic")
}
```