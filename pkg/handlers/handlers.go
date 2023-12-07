package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mariogarzac/moodeck/config"
	"github.com/mariogarzac/moodeck/pkg/models"
)

var Repo *Repository

type Repository struct {
    App *config.AppConfig
}

func NewRepository(a *config.AppConfig) *Repository {
    return &Repository{
        App: a,
    }
}

func NewHandlers(r *Repository){
    Repo = r
}

func (m *Repository)Hello(c echo.Context) error {
    test := make(map[string]string)
    test["hello"] = "world"

    fmt.Println(m.App.DataBase)
    
    log.Println("done testing db")

    return c.Render(http.StatusOK, "index.html", models.TemplateData{
        StringMap: test,
    })
}

