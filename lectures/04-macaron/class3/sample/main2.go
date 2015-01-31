package main

import (
	"log"
	"os"

	"github.com/Unknwon/macaron"
)

var logger = log.New(os.Stdout, "[app] ", 0)

func main() {
	m := macaron.New()

	m.Get("/", func(l *log.Logger) {
		l.Println("这是默认日志器")
	})

	// 获取请求级别服务
	m.Get("/logger", myLogger, func(l *log.Logger) {
		l.Println("我正在使用全局日志器")
	})

	m.Run()
}

func myLogger(ctx *macaron.Context) {
	// 该服务的映射只会对当前请求的后续处理器有效
	ctx.Map(logger)
}
