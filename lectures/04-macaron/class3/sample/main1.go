package main

import (
	"log"
	"os"

	"github.com/Unknwon/macaron"
)

var logger = log.New(os.Stdout, "[app] ", 0)

func main() {
	m := macaron.New()
	m.Map(logger) // 映射全局服务

	// 获取全局服务
	m.Get("/logger", func(l *log.Logger) {
		l.Println("我正在使用全局日志器")
	})

	m.Run()
}
