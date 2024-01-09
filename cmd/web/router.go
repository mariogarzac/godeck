package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mariogarzac/godeck/pkg/handlers"
)

func routes(e *echo.Echo){

    e.GET("/", handlers.Repo.HandleHomePage)
    e.GET("/get-files", handlers.Repo.HandleFileList)

    e.GET("/register", handlers.Repo.HandleRegisterPage)
    e.POST("/register", handlers.Repo.HandleRegister)

    e.POST("/upload-file", handlers.Repo.HandleFileUpload)
}
