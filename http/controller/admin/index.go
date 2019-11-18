package admin

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

func AdminIndex(ctx echo.Context) error {
	return ctx.HTML(http.StatusOK, "Admin")
}
