package ULog

import (
	"fmt"
	"io"
	"os"
	"time"
)

type baseLogger struct {
	level LogLevel
	out   io.Writer
}

//Output() io.Writer
//SetOutput(w io.Writer)
//Level() LogLevel
//SetLevel(v LogLevel)
//Trace(i ...interface{})
//Debug(i ...interface{})
//Info(i ...interface{})
//Warn(i ...interface{})
//Error(i ...interface{})
//Fatal(i ...interface{})
//Panic(i ...interface{})

func New() Logger {
	l := &baseLogger{
		level: InfoLevel,
		out:   os.Stdout,
	}
	return l
}

func (l *baseLogger) Output() io.Writer {
	return l.out
}

func (l *baseLogger) SetOutput(w io.Writer) {
	l.out = w
}

func (l *baseLogger) SetLevel(v LogLevel) {
	l.level = v
}

func (l *baseLogger) GetLevel() LogLevel {
	return l.level
}

func (l *baseLogger) Traceln(i ...interface{}) {
	l.log(TraceLevel, i...)
}

func (l *baseLogger) Debugln(i ...interface{}) {
	l.log(DebugLevel, i...)
}
func (l *baseLogger) Infoln(i ...interface{}) {
	l.log(InfoLevel, i...)
}
func (l *baseLogger) Warnln(i ...interface{}) {
	l.log(WarnLevel, i...)
}
func (l *baseLogger) Errorln(i ...interface{}) {
	l.log(ErrorLevel, i...)
}
func (l *baseLogger) Fatalln(i ...interface{}) {
	l.log(FatalLevel, i...)
	os.Exit(1)
}
func (l *baseLogger) Panicln(i ...interface{}) {
	l.log(PanicLevel, i...)
	panic(i)
}

func (l *baseLogger) isLevelEnabled(level LogLevel) bool {
	return l.level >= level
}

func (l *baseLogger) log(level LogLevel, args ...interface{}) {
	if l.isLevelEnabled(level) {
		fmt.Fprintln(l.out, l.genTime(), l.genNoColorLogLevelStr(level), l.sprintlnn(args...))
	}
}

func (l *baseLogger) genNoColorLogLevelStr(level LogLevel) string {
	switch level {
	case TraceLevel:
		return string(TraceTagStr)
	case DebugLevel:
		return string(DebugTagStr)
	case InfoLevel:
		return string(InfoTagStr)
	case WarnLevel:
		return string(WarnTagStr)
	case ErrorLevel:
		return string(ErrorTagStr)
	case FatalLevel:
		return string(FatalTagStr)
	case PanicLevel:
		return string(PanicTagStr)
	default:
		return ""
	}
}

func (l *baseLogger) genTime() string {
	return time.Now().UTC().Format("2006-01-02 15:04:05")
}

func (l *baseLogger) sprintlnn(args ...interface{}) string {
	msg := fmt.Sprintln(args...)
	return msg[:len(msg)-1]
}
