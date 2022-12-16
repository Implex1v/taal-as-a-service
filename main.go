package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.GET("/words", GetWords)
	e.GET("/health", GetHealth)

	err := e.Start(":8080")
	if err != nil {
		return
	}
}
