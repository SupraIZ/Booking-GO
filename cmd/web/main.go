package main

import (
	"fmt"
	"log"
	"github.com/SupraIZ/booking/packages/config"
	"github.com/SupraIZ/booking/packages/handlers"
	"github.com/SupraIZ/booking/packages/renders"
	"net/http"
	"time"
	"github.com/alexedwards/scs/v2"
)

const portNum = ":8080"
var app config.AppConfig
var session *scs.SessionManager

// main is the entry point of the program
func main() {
	//change this to true if in production
	app.Inproduction = false

	//sessions for the work to logged in
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true //if false then the cookie will be destroyed when browser closes.
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.Inproduction //if true then the site will be opened in https.

	app.Session = session

	tc, err := renders.CreateTemplateCache()
	if err!=nil{
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false // this one basically use cache memory to rebuilt the page 

	repo :=handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	renders.NewTemplate(&app)

	fmt.Println(fmt.Sprintf("Starting appication on port %s", portNum))
	
	srv := &http.Server{
		Addr: portNum,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

	//the '_' above basically tells the go compiler that it doesn't need to handle the error if any exists.
}

// in order to run all the files in the directory we have to use the command (go run .)
