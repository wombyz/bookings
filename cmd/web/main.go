package main

import (
	"encoding/gob"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/wombyz/bookings/internal/config"
	"github.com/wombyz/bookings/internal/handlers"
	"github.com/wombyz/bookings/internal/helpers"
	"github.com/wombyz/bookings/internal/models"
	"github.com/wombyz/bookings/internal/render"
	"log"
	"net/http"
	"os"
	"time"
)

const portNumber = ":8080"

var session *scs.SessionManager
var app config.AppConfig
var infoLog *log.Logger
var errorLog *log.Logger

// main is the main application function
func main() {

	err := run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Starting application on port %s", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() error {

	//What am I going to put in the session
	gob.Register(models.Reservation{})

	// change to true when in production
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)
	helpers.NewHelpers(&app)

	return nil
}
