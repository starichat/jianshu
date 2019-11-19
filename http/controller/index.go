package controller

import (
	"github.com/labstack/echo/v4"

	. "github.com/starichat/jianshu/http/"
)

type IndexController struct{}

// RegisterRoute 管理 IndexController 的路由
func (i IndexController) RegisterRoute(g *echo.Group) {
	g.GET("/", i.index)
}

func (i IndexController) index(ctx echo.Context) error {
	return Render(ctx, "index.html", nil)
}
