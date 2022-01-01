package ULog

import (
	"io"
	"strings"
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

type logLevelTagStr string

const (
	panicTagStr logLevelTagStr = "[PANI]"
	fatalTagStr logLevelTagStr = "[FATA]"
	errorTagStr logLevelTagStr = "[ERRO]"
	warnTagStr  logLevelTagStr = "[WARN]"
	infoTagStr  logLevelTagStr = "[INFO]"
	debugTagStr logLevelTagStr = "[DEBU]"
	traceTagStr logLevelTagStr = "[TRAC]"
)

type (
	// Logger defines the logging interface.
	Logger interface {
		SetOutput(w io.Writer)
		SetLevel(v LogLevel)
		Traceln(i ...interface{})
		Debugln(i ...interface{})
		Infoln(i ...interface{})
		Warnln(i ...interface{})
		Errorln(i ...interface{})
		Fatalln(i ...interface{})
		Panicln(i ...interface{})
	}
)

func LogLevelStrToLevel(levelStr string) LogLevel {
	s := strings.ToLower(levelStr)

	switch s {
	case "panic", "pani":
		return PanicLevel
	case "fata", "fatal":
		return FatalLevel
	case "erro", "error":
		return ErrorLevel
	case "warn":
		return WarnLevel
	case "info":
		return InfoLevel
	case "debu", "debug":
		return DebugLevel
	case "trac", "trace":
		return TraceLevel
	default:
		return InfoLevel
	}
}
