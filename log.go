package ULog

import (
	"io"
)

type LogLevel string

const LEVEL_PANIC LogLevel = "PANI"
const LEVEL_FATAL LogLevel = "FATA"
const LEVEL_ERROR LogLevel = "ERRO"
const LEVEL_WARN LogLevel = "WARN"
const LEVEL_INFO LogLevel = "INFO"
const LEVEL_DEBUG LogLevel = "DEBU"
const LEVEL_TRACE LogLevel = "TRAC"

type (
	// Logger defines the logging interface.
	Logger interface {
		Output() io.Writer
		SetOutput(w io.Writer)
		Level() LogLevel
		SetLevel(v LogLevel)
		Trace(i ...interface{})
		//Traceln(i ...interface{})
		//Tracef(format string, args ...interface{})
		Debug(i ...interface{})
		//Debugln(i ...interface{})
		//Debugf(format string, args ...interface{})
		Info(i ...interface{})
		//Infoln(i ...interface{})
		//Infof(format string, args ...interface{})
		Warn(i ...interface{})
		//Warnln(i ...interface{})
		//Warnf(format string, args ...interface{})
		Error(i ...interface{})
		//Errorln(i ...interface{})
		//Errorf(format string, args ...interface{})
		Fatal(i ...interface{})
		//Fatalln(i ...interface{})
		//Fatalf(format string, args ...interface{})
		Panic(i ...interface{})
		//Panicln(i ...interface{})
		//Panicf(format string, args ...interface{})
	}
)
