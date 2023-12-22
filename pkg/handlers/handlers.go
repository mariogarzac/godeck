package handlers

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/mariogarzac/godeck/model"
	"github.com/mariogarzac/godeck/pkg/config"
	"github.com/mariogarzac/godeck/pkg/render"
	"github.com/mariogarzac/godeck/view/user"
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

func (m *Repository)HandleRegister(c echo.Context) error {
    username := c.FormValue("username")
    email := c.FormValue("email")
    password := c.FormValue("password")

    // err := m.App.DataBase.RegisterUser(username, email, password)
    //
    // if err != nil {
    //     return err
    // }

    fmt.Println(username, email, password)

    u := model.User{
        Username: username,
        Email: email,
        Password: password,
    }

    return render.Render(c, user.Show(u))
}

func (m *Repository)HandleRegisterPage(c echo.Context) error {
    return render.Render(c, user.RegisterUser())
}

