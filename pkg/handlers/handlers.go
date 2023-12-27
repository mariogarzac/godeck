package handlers

import (
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mariogarzac/godeck/model"
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

func (m *Repository)HandleRegister(c echo.Context) error {
    username := c.FormValue("username")
    email := c.FormValue("email")
    password := c.FormValue("password")

    err := m.App.DataBase.RegisterUser(username, email, password)

    if err != nil {
        return err
    }

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

func (m *Repository)HandleHomePage(c echo.Context) error {
    return render.Render(c, home.LoadHome())
}

func (m *Repository)HandleFileUpload(c echo.Context) error {

    var fileType, fileName string
    var response models.Response
    var fileSize int64

    file, err := c.FormFile("audio-file")
    if err != nil {
        log.Println("Could not get file",err)
        return err
    }

    fileType, fileName, fileSize, err = utils.UploadFile(file)

    response = models.Response{
        Message: "Success",
        Data : models.File{
            FileName: fileName,
            FileType: fileType,
            FileSize: fileSize,
        },
    }

    if err != nil {
        response.Message = "Error"
        return c.JSON(http.StatusSeeOther, response)
    }

    fileExt := (strings.Split(fileType, "/"))[1]
    err = m.App.DataBase.SaveFileToDB(fileExt, fileName, fileSize)

    if err != nil {
        return err
    }

    return c.JSON(http.StatusOK, response)
}

func (m *Repository)HandleUploadMessage(c echo.Context) error {

    return c.JSON(http.StatusOK, "")
}

func (m *Repository)HandleShowUploadedFiles(c echo.Context) error {

    files, err := m.App.DataBase.GetFilesByUserId(1)

    if err != nil {
        return err
    }

    return c.JSON(http.StatusOK, files)
}
