package admin

import (
	"github.com/labstack/echo/v4"
)

// RegisterRoutes 统一注册路由
func RegisterRoutes(g *echo.Group) {
	g.GET("", AdminIndex)
}
