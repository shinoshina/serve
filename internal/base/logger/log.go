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

	debugLogger = log.New(os.Stderr, colorConvert(32,"[DEBUG]  \n"), 0)
	infoLogger = log.New(os.Stderr, colorConvert(32,"[INFO]  \n"), 0)
	warnLogger = log.New(os.Stderr, colorConvert(36,"[WARN]  \n"), 0)
	errLogger = log.New(os.Stderr, colorConvert(33,"[ERROR]  \n"), 0)
	fatalLogger = log.New(os.Stderr, colorConvert(31,"[FATAL]  \n"), 0)

}

func colorConvert(color int,raw string)(s string){

	return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", color,raw )
}

func handleRaw(raw string, color int) (s string) {
	pc, path, line, _ := runtime.Caller(1)
	location := fmt.Sprintf("[Location]: %s [Line]: %d\n", path, line)
	funcName := runtime.FuncForPC(pc).Name()
	info := location + "[Caller]: " + funcName + "\n"

	k := "[Time]: " + timer.CurrentTime() + "\n" + info + "[Message]: " + raw
	s = fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", color, k)
	return

}
func Debugf(format string, v ...interface{}) {
	debugLogger.Printf(handleRaw(format, 32), v...)
}
func Infof(format string, v ...interface{}) {
	infoLogger.Printf(handleRaw(format, 32), v...)
}
func Warnf(format string, v ...interface{}) {
	warnLogger.Printf(handleRaw(format, 36), v...)
}
func Errorf(format string, v ...interface{}) {
	errLogger.Printf(handleRaw(format, 33), v...)
}
func Fatalf(format string, v ...interface{}) {
	fatalLogger.Printf(handleRaw(format, 31), v...)
}
