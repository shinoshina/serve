package logger

import (
	"fmt"
	"goserve/internal/base/timer"
	"io"
	"log"
	"os"
	"runtime"
)

type level int8

const (
	DEBUG level = iota
	INFO
	WARNNING
	ERROR
	FATAL
)
const (
	logFlag = log.Ldate | log.Ltime
)

var (
	logFile     io.Writer
	debugLogger *log.Logger
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errLogger   *log.Logger
	fatalLogger *log.Logger
)

func init() {

	debugLogger = log.New(os.Stderr, "[DEBUG]  \n", 0)
	infoLogger = log.New(os.Stderr, "[INFO]", logFlag)
	warnLogger = log.New(os.Stderr, "[WARN]", logFlag)
	errLogger = log.New(os.Stderr, "[ERROR]", logFlag)
	fatalLogger = log.New(os.Stderr, "[FATAL]", logFlag)

}

func handleRaw(raw string, color int) (s string) {
	pc, path, line, _ := runtime.Caller(1)
	location := fmt.Sprintf("[location]: %s [line]: %d\n", path, line)
	funcName := runtime.FuncForPC(pc).Name()
	info := location + "[caller]: " + funcName + "\n"

	k := "[time]: " + timer.CurrentTime() + "\n" + info + "[message]: " + raw
	s = fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", 32, k)
	return

}
func Debugf(format string, v ...interface{}) {

	debugLogger.Printf(handleRaw(format, 32), v...)
}
func Infof(format string, v ...interface{}) {
	infoLogger.Printf(format, v...)
}
func Warnf(format string, v ...interface{}) {
	warnLogger.Printf(format, v...)
}
func Errorf(format string, v ...interface{}) {
	errLogger.Printf(format, v...)
}
func Fatalf(format string, v ...interface{}) {
	fatalLogger.Printf(format, v...)
}
