package main

import (
	"fmt"
	"log"

	"github.com/Unknwon/macaron"
)

func main() {
	m := macaron.Classic()

	// 服务静态文件
	m.Use(macaron.Static("public"))

	// 在同一个请求中，多个处理器之间可相互传递数据
	m.Get("/", handler1, handler2, handler3)

	// 获取请求参数
	m.Get("/q", queryHandler)

	// 获取远程 IP 地址
	m.Get("/ip", ipHandler)

	// 处理器可以让出处理权限，让之后的处理器先执行
	m.Get("/next", next1, next2, next3)

	// 设置和获取 Cookie
	m.Get("/cookie/set", setCookie)
	m.Get("/cookie/get", getCookie)

	// 响应流
	m.Get("/resp", respHandler)

	// 请求对象
	m.Get("/req", reqHandler)

	// 请求级别容错恢复
	m.Get("/panic", panicHandler)

	// 全局日志器
	m.Get("/log", logger)

	m.Run()
}

func handler1(ctx *macaron.Context) {
	ctx.Data["Num"] = 1
}

func handler2(ctx *macaron.Context) {
	ctx.Data["Num"] = ctx.Data["Num"].(int) + 1
}

func handler3(ctx *macaron.Context) string {
	return fmt.Sprintf("Num: %d", ctx.Data["Num"])
}

func queryHandler(ctx *macaron.Context) {
	fmt.Println(ctx.Query("uid"))
	fmt.Println(ctx.QueryInt("uid"))
	fmt.Println(ctx.QueryInt64("uid"))
}

func ipHandler(ctx *macaron.Context) string {
	return ctx.RemoteAddr()
}

func next1(ctx *macaron.Context) {
	fmt.Println("位于处理器 1 中")

	ctx.Next()

	fmt.Println("退出处理器 1")
}

func next2(ctx *macaron.Context) {
	fmt.Println("位于处理器 2 中")

	ctx.Next()

	fmt.Println("退出处理器 2")
}

func next3(ctx *macaron.Context) {
	fmt.Println("位于处理器 3 中")

	ctx.Next()

	fmt.Println("退出处理器 3")
}

func setCookie(ctx *macaron.Context) {
	ctx.SetCookie("user", "wuwen")
}

func getCookie(ctx *macaron.Context) string {
	return ctx.GetCookie("user")
}

func respHandler(ctx *macaron.Context) {
	ctx.Resp.Write([]byte("你好，世界！"))
}

func reqHandler(ctx *macaron.Context) string {
	return ctx.Req.Method
}

func panicHandler() {
	panic("有钱，任性！")
}

func logger(l *log.Logger) {
	l.Println("打印一行日志")
}
