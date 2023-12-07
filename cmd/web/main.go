package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mariogarzac/moodeck/config"
	"github.com/mariogarzac/moodeck/pkg/handlers"
	"github.com/mariogarzac/moodeck/pkg/render"
)


var app config.AppConfig

func main(){

    e := echo.New()
    e.Renderer = render.NewTemplate()

    e.Static("/static", "static")

    repo := handlers.NewRepository(&app)
    handlers.NewHandlers(repo)
}
