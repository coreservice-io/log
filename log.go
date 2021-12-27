package ULog

import (
	"io"
)

type LogLevel uint32

const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel LogLevel = iota
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel
)

type (
	// Logger defines the logging interface.
	Logger interface {
		Output() io.Writer
		SetOutput(w io.Writer)
		GetLevel() LogLevel
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
