package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mariogarzac/godeck/pkg/config"
	"github.com/mariogarzac/godeck/pkg/models"
	"github.com/mariogarzac/godeck/pkg/render"
	"github.com/mariogarzac/godeck/pkg/utils"
	"github.com/mariogarzac/godeck/view/home"
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

func (m *Repository)HandleRegisterPage(c echo.Context) error {
    return render.Render(c, user.RegisterUser())
}

func (m *Repository)HandleHomePage(c echo.Context) error {
    return render.Render(c, home.LoadHome())
}

func (m *Repository)HandleFileList(c echo.Context) error {

    var files []models.File
    files, err := m.App.DataBase.GetFilesByUserId(1)

    if err != nil {
        return err
    }

    return render.Render(c, user.ShowFiles(files))
}

func (m *Repository)HandleRegister(c echo.Context) error {
    username := c.FormValue("username")
    email := c.FormValue("email")
    password := c.FormValue("password")

    err := m.App.DataBase.RegisterUser(username, email, password)

    if err != nil {
        return c.HTML(http.StatusSeeOther, "An error occured. Please try again.")
    }

    return c.Redirect(http.StatusSeeOther, "/")
}

func (m *Repository) HandleFileUpload(c echo.Context) error {
    const MAX_SIZE = 5000000

    var fileType, fileName string
    var fileSize int64

    file, err := c.FormFile("audio-file")
    if err != nil {
        return c.String(http.StatusBadRequest, "Error reading file.")
    }

    fileType, fileName, fileSize, err = utils.UploadFile(file)
    fmt.Println(fileName, fileSize)

    if err != nil {
        return c.String(http.StatusInternalServerError, "Error uploading file.")
    }

    if MAX_SIZE < fileSize {
        return c.String(http.StatusBadRequest, "File size exceeded.")
    }

    fileExt := (strings.Split(fileType, "/"))[1]
    err = m.App.DataBase.SaveFileToDB(fileExt, fileName, fileSize)

    if err != nil {
        return c.String(http.StatusInternalServerError, "Error saving to database. Try again.")
    }

    return c.String(http.StatusOK, fileName)
}
