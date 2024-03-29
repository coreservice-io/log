package log

import (
	"io"
	"strings"
)

type LogLevel int32

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

type logTag string

const (
	PanicTagStr logTag = "[PANI]"
	FatalTagStr logTag = "[FATA]"
	ErrorTagStr logTag = "[ERRO]"
	WarnTagStr  logTag = "[WARN]"
	InfoTagStr  logTag = "[INFO]"
	DebugTagStr logTag = "[DEBU]"
	TraceTagStr logTag = "[TRAC]"
)

type (
	// Logger defines the logging interface.
	Logger interface {
		SetOutput(w io.Writer)
		SetLevel(v LogLevel)
		GetLevel() LogLevel
		Traceln(i ...interface{})
		Debugln(i ...interface{})
		Infoln(i ...interface{})
		Warnln(i ...interface{})
		Errorln(i ...interface{})
		Fatalln(i ...interface{})
		Panicln(i ...interface{})
		PrintLastN(int64, []LogLevel)                 //print out last n lines of logs with specified levels sorted by time
		GetLastN(int64, []LogLevel) ([]string, error) //return last n lines of logs with specified levels sorted by time
	}
)

func ParseLogLevel(levelStr string) LogLevel {
	s := strings.TrimSpace(strings.ToLower(levelStr))

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

func LogLevelToTag(level LogLevel) logTag {
	switch level {
	case TraceLevel:
		return TraceTagStr
	case DebugLevel:
		return DebugTagStr
	case InfoLevel:
		return InfoTagStr
	case WarnLevel:
		return WarnTagStr
	case ErrorLevel:
		return ErrorTagStr
	case FatalLevel:
		return FatalTagStr
	case PanicLevel:
		return PanicTagStr
	default:
		return InfoTagStr
	}
}
