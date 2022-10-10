package logger

import (
	"bytes"
	"log"

	"github.com/sirupsen/logrus"
)

var (
	buf      bytes.Buffer
	LogInfo  = log.New(&buf, "RESULT : ", log.Lmsgprefix)
	LogError = log.New(&buf, "ERROR : ", log.Lmsgprefix)
)

type MyLog struct {
}

func NewMyLog() *MyLog {
	return &MyLog{}
}

var mlog *MyLog

func (m *MyLog) Info(args ...interface{}) {
	LogInfo.Print(args...)
	logrus.Info(&buf)
	buf.Reset()
}

func (m *MyLog) InfoF(format string, args ...interface{}) {
	LogInfo.Printf(format, args...)
	logrus.Info(&buf)
	buf.Reset()
}

func (m *MyLog) Error(args ...interface{}) {
	LogError.Print(args...)
	logrus.Error(&buf)
	buf.Reset()
}

func (m *MyLog) ErrorF(format string, args ...interface{}) {
	LogError.Printf(format, args...)
	logrus.Error(&buf)
	buf.Reset()
}
