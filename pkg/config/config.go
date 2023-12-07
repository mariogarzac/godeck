package config

import (
	"database/sql"
	"html/template"
	"log"

	"github.com/gorilla/sessions"
)

type AppConfig struct {
    UseCache bool
    TemplateCache map[string]*template.Template
    InfoLog *log.Logger
    Session *sessions.CookieStore
    DataBase *sql.DB
}
