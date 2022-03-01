package logger

import (
	"github.com/onrik/logrus/filename"
	"github.com/sirupsen/logrus"
	"os"
)

type Logger struct {
	level       int // debug 0 info 3 err 5
	innerLogger *logrus.Logger
	jsonOut     bool
}

func New() *Logger {
	var log = logrus.New()
	formatter := new(logrus.TextFormatter)
	formatter.ForceColors = true
	filenameHook := filename.NewHook()
	filenameHook.Field = "SRC"
	filenameHook.Skip = 5
	filenameHook.SkipPrefixes = []string{"logrus/", "logrus@", "logger"}
	log.AddHook(filenameHook)
	log.SetFormatter(formatter)
	log.SetOutput(os.Stdout)

	newLogger := new(Logger)
	newLogger.jsonOut = false
	newLogger.innerLogger = log
	return newLogger
}

func (logobj *Logger) Info(format string, v ...interface{}) {
	if len(v) == 0 {
		logobj.innerLogger.Info(format)
	} else {
		logobj.innerLogger.Infof(format, v...)
	}
}

func (logobj *Logger) Infof(format string, v ...interface{}) {
	logobj.innerLogger.Infof(format, v...)
}

func (logobj *Logger) Sync() {
}
