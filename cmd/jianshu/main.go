package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"

	"github.com/starichat/jianshu/global"
	"github.com/starichat/jianshu/http/controller"
	"github.com/starichat/jianshu/http/controller/admin"
)

var (
	// 编译相关信息
	// 初始化为 unknown，如果编译时没有传入这些值，则为 unknown
	gitCommitLog = "unknown"
	buildTime    = "unknown"
	gitRelease   = "unknown"
)

func init() {
	global.Init()

	global.App.FillBuildInfo(gitCommitLog, buildTime, gitRelease)
}

func main() {
	// 获得 echo 的一个实例
	e := echo.New()

	// 全局中间件
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// 服务静态文件
	e.Static("/static", "static")

	// 区分前后端组，因为它们需要的中间件不一样

	// 前端路由组
	frontGroup := e.Group("")
	controller.RegisterRoutes(frontGroup)

	// 后端路由组
	adminGroup := e.Group("/admin")
	admin.RegisterRoutes(adminGroup)

	// 给默认端口
	viper.SetDefault("http.port", "9999")
	host := viper.GetString("http.host")
	port := viper.GetString("http.port")

	addr := host + ":" + port

	// 除了 Start 还有 StartTLS 和 StartAutoTLS
	e.Logger.Fatal(e.Start(addr))
}
