package main

import (
	"encoding/gob"
	"fmt"
	"net/http"
	"time"

	"github.com/sakurtek/goserver/bookingremyconcept/internal/config"
	"github.com/sakurtek/goserver/bookingremyconcept/internal/handlers"
	"github.com/sakurtek/goserver/bookingremyconcept/internal/model"
	"github.com/sakurtek/goserver/bookingremyconcept/internal/render"

	"github.com/alexedwards/scs/v2"
)

// untuk portnumber
const portNumber = ":2102"

// untuk mendefiniskan variabel confit.AppConfig
var app config.AppConfig
var session *scs.SessionManager

func main() {
	// registserasi session
	gob.Register(model.Reservation{})

	fileserver := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/*", http.StripPrefix("/static/", fileserver))

	// UBAH INI MENJADI -TRUE- APABILA PRODUKSI ATAU SUDAH MAU DI DEPLOY KE SERVER
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	/* SIMPAN sessionmanager KE config.AppConfig */
	app.Session = session

	/* pelajari ini dengan baik karena dengan dua fungsi ini Penggunaan SESSION berhasil */
	// Penting ini----
	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)
	// Penting ini----

	// pelajari kembali bagian ini
	render.NewTemplates(&app)

	//fmt.Println("Server running in localhost port", portNumber)
	config.GetInfoProgram(portNumber)

	myserver := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err := myserver.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
