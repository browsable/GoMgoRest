package routes

import (
	"GoServer/api/todo/route"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())

	todoroute.Init(e)
}
