package config

import (
	"html/template"
	"log"

	"github.com/gorilla/sessions"
	"github.com/mariogarzac/moodeck/db"
)

type AppConfig struct {
    UseCache bool
    TemplateCache map[string]*template.Template
    InfoLog *log.Logger
    Session *sessions.CookieStore
    DataBase *db.Database
}
