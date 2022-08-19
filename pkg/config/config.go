package config

import (
	"log"
	"text/template"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool
	InfoLog       *log.Logger
	InProduction  bool
	TemplateCache map[string]*template.Template
	Session       *scs.SessionManager
}
