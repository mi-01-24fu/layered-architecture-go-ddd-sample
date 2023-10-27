package handler

import "github.com/labstack/echo/v4"

func InitRouting(e *echo.Echo, memoHandler MemoHandler) {
	e.POST("/memo", memoHandler.Post())
	e.GET("/memo", memoHandler.Get())
}
