package logger

import (
	"os"
	"path"
	"runtime"
	"strconv"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

var (
	Console  *log.Logger
	File     *log.Logger
	filename string = "application.log"
)

func logFileConfig() {
	if os.Getenv("UNIT_TEST") == "1" {
		return
	}
	File = log.New()
	File.SetFormatter(&log.JSONFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			pc, file, line, _ := runtime.Caller(8)

			fileName := path.Dir(file) + "/" + path.Base(file) + ", lineNumber:" + strconv.Itoa(line)
			funcName := runtime.FuncForPC(pc).Name()

			return funcName, fileName
		},
	})
	File.SetReportCaller(true)
	File.SetLevel(logrus.TraceLevel)

	// set file
	logFile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		logrus.Fatal("failed open log file :", err.Error())
	}
	File.SetOutput(logFile)
}

func logConsoleConfig() {
	Console = log.New()
	Console.SetFormatter(&log.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			pc, file, line, _ := runtime.Caller(8)

			fileName := path.Dir(file) + "/" + path.Base(file) + ", lineNumber:" + strconv.Itoa(line)
			funcName := runtime.FuncForPC(pc).Name()

			return funcName, fileName
		},
	})
	Console.SetReportCaller(true)
	Console.SetLevel(logrus.InfoLevel)
}

func init() {
	logConsoleConfig()
}
