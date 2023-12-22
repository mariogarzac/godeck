package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mariogarzac/godeck/db"
	"github.com/mariogarzac/godeck/pkg/config"
	"github.com/mariogarzac/godeck/pkg/handlers"
)

var app config.AppConfig
var portNumber = ":3000"

func main(){

    e := echo.New()
    // e.Renderer = render.NewTemplate()

    e.Static("/assets", "assets")

    db.InitDB()
    app.DataBase = db.GetDB()
    defer app.DataBase.Close()

    repo := handlers.NewRepository(&app)
    handlers.NewHandlers(repo)


    routes(e)

    e.Logger.Fatal(e.Start(portNumber))
}
