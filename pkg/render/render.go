package render

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type Template struct {
    templates *template.Template
}

func NewTemplate() *Template{
    return &Template{
        templates: template.Must(template.ParseGlob("static/templates/*.html")),
    }
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

