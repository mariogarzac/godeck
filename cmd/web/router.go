package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mariogarzac/moodeck/pkg/handlers"
)

func routes(e *echo.Echo){
    e.GET("/", handlers.Repo.HandleRegisterPage)
    e.POST("/", handlers.Repo.HandleRegister)
}
