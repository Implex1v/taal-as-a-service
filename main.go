package main

import (
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	enablePrometheus(e)

	e.GET("/words", GetWords)
	e.GET("/health", GetHealth)

	err := e.Start(":8080")
	if err != nil {
		return
	}
}

func enablePrometheus(e *echo.Echo) {
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)
}
