package logger

var mylog = NewMyLog()

func Info(args ...interface{}) {
	mlog.Info(args...)
}

func InfoF(format string, args ...interface{}) {
	mlog.InfoF(format, args...)
}

func Error(args ...interface{}) {
	mlog.Error(args...)
}

func ErrorF(format string, args ...interface{}) {
	mlog.ErrorF(format, args...)
}
