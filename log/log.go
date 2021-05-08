package log

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

/// 相比于标准库的 log 包，这里实现了所属文件、行号、日志等级等 feature
var (
	errorLog = log.New(os.Stdout, "\033[31m[error]\033[0m ", log.LstdFlags|log.Lshortfile)
	infoLog  = log.New(os.Stdout, "\033[34m[info ]\033[0m ", log.LstdFlags|log.Lshortfile)
	loggers  = []*log.Logger{errorLog, infoLog}
	mu       sync.Mutex
)

/// 日志方法
var (
	Error  = errorLog.Println
	Errorf = errorLog.Printf
	Info   = infoLog.Println
	infof  = infoLog.Printf
)

const (
	InfoLevel = iota
	ErrorLevel
	Disabled
)

func SetLevel(level int) {
	mu.Lock()
	defer mu.Unlock()

	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}

	if level > ErrorLevel {
		errorLog.SetOutput(ioutil.Discard)
	}
	if level > InfoLevel {
		infoLog.SetOutput(ioutil.Discard)
	}
}
