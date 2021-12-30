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

type LogLevelTagStr string

const (
	PanicTagStr LogLevelTagStr = "[PANI]"
	FatalTagStr LogLevelTagStr = "[FATA]"
	ErrorTagStr LogLevelTagStr = "[ERRO]"
	WarnTagStr  LogLevelTagStr = "[WARN]"
	InfoTagStr  LogLevelTagStr = "[INFO]"
	DebugTagStr LogLevelTagStr = "[DEBU]"
	TraceTagStr LogLevelTagStr = "[TRAC]"
)

type (
	// Logger defines the logging interface.
	Logger interface {
		SetOutput(w io.Writer)
		SetLevel(v LogLevel)
		Trace(i ...interface{})
		Debug(i ...interface{})
		Info(i ...interface{})
		Warn(i ...interface{})
		Error(i ...interface{})
		Fatal(i ...interface{})
		Panic(i ...interface{})
	}
)
