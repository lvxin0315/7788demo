package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Logger 日志接口
type Logger interface {
	// WriteLog 写入内容
	WriteLog(logs ...string)
}

// TerminalLogger 终端方式
type TerminalLogger struct {
}

func (log TerminalLogger) WriteLog(logs ...string) {
	fmt.Println(logs)
}

// FileLogger 文件方式
type FileLogger struct {
}

func (log FileLogger) WriteLog(logs ...string) {
	_ = ioutil.WriteFile("info.log", []byte(strings.Join(logs, "")), 0755)
}

func main() {
	var logger Logger
	x := 1
	// 根据情况调用
	if x == 1 {
		logger = new(TerminalLogger)
	} else {
		logger = new(FileLogger)
	}
	writeLog(logger, "哈", "哈", "哈")
}

func writeLog(logger Logger, logs ...string) {
	logger.WriteLog(logs...)
}
